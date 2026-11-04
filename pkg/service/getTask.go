package service

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/vlonebara/todo-list/pkg/db"
)

// Task содержит сохранённые данные задачи.
type Task struct {
	ID      int64
	Date    string
	Title   string
	Comment string
	Repeat  string
}

// GetTask возвращает задачу с указанным ID.
func GetTask(id string) (Task, error) {
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil || taskID < 1 {
		return Task{}, fmt.Errorf("Некорректный идентификатор задачи")
	}

	database, err := db.OpenDatabase()
	if err != nil {
		return Task{}, fmt.Errorf("open database: %w", err)
	}
	defer database.Close()

	var task Task
	err = database.QueryRow(
		`SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`,
		taskID,
	).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err == sql.ErrNoRows {
		return Task{}, fmt.Errorf("Задача не найдена")
	}
	if err != nil {
		return Task{}, fmt.Errorf("get task: %w", err)
	}
	return task, nil
}
