package service

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/vlonebara/todo-list/pkg/db"
)

// EditTaskStatus помечает задачу выполненной и переносит её или удаляет.
func EditTaskStatus(id string) error {
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil || taskID < 1 {
		return fmt.Errorf("Некорректный идентификатор задачи")
	}

	database, err := db.OpenDatabase()
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer database.Close()

	task, err := getTaskForStatus(database, taskID)
	if err != nil {
		return err
	}
	if task.Repeat == "" {
		return deleteCompletedTask(database, taskID)
	}

	date, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		return fmt.Errorf("calculate next task date: %w", err)
	}
	return updateCompletedTaskDate(database, taskID, date)
}

// getTaskForStatus загружает дату и правило повторения задачи.
func getTaskForStatus(database *sql.DB, id int64) (Task, error) {
	var task Task
	err := database.QueryRow(`SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`, id).
		Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err == sql.ErrNoRows {
		return Task{}, fmt.Errorf("Задача не найдена")
	}
	if err != nil {
		return Task{}, fmt.Errorf("get task: %w", err)
	}
	return task, nil
}

// deleteCompletedTask удаляет выполненную одноразовую задачу.
func deleteCompletedTask(database *sql.DB, id int64) error {
	result, err := database.Exec(`DELETE FROM scheduler WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete completed task: %w", err)
	}
	deleted, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted task: %w", err)
	}
	if deleted == 0 {
		return fmt.Errorf("Задача не найдена")
	}
	return nil
}

// updateCompletedTaskDate записывает дату следующего выполнения.
func updateCompletedTaskDate(database *sql.DB, id int64, date string) error {
	result, err := database.Exec(`UPDATE scheduler SET date = ? WHERE id = ?`, date, id)
	if err != nil {
		return fmt.Errorf("update task date: %w", err)
	}
	updated, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check updated task: %w", err)
	}
	if updated == 0 {
		return fmt.Errorf("Задача не найдена")
	}
	return nil
}
