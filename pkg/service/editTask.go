package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vlonebara/todo-list/pkg/db"
)

// EditTask проверяет и сохраняет новые данные задачи.
func EditTask(id, date, title, comment, repeat string) error {
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil || taskID < 1 {
		return fmt.Errorf("Некорректный идентификатор задачи")
	}

	date, err = prepareTaskDate(date, title, repeat, time.Now())
	if err != nil {
		return err
	}

	database, err := db.OpenDatabase()
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer database.Close()

	result, err := database.Exec(
		`UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`,
		date, title, comment, repeat, taskID,
	)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
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
