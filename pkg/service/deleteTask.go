package service

import (
	"fmt"
	"strconv"
)

// DeleteTask удаляет задачу с указанным ID.
func DeleteTask(id string) error {
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil || taskID < 1 {
		return fmt.Errorf("Некорректный идентификатор задачи")
	}

	database, err := openDatabase()
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer database.Close()

	result, err := database.Exec(`DELETE FROM scheduler WHERE id = ?`, taskID)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
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
