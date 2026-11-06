package service

import (
	"fmt"
	"strconv"

	"github.com/vlonebara/todo-list/pkg/db"
)

// Task использует тот же тип данных задачи, что и хранилище.
type Task = db.Task

// TaskService проверяет данные задач и обращается к общему хранилищу.
type TaskService struct {
	store db.TaskStore
}

// NewTaskService создаёт сервис с готовым хранилищем задач.
func NewTaskService(store db.TaskStore) TaskService {
	return TaskService{store: store}
}

// parseTaskID проверяет и преобразует ID задачи в число.
func parseTaskID(id string) (int64, error) {
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil || taskID < 1 {
		return 0, fmt.Errorf("Некорректный идентификатор задачи")
	}
	return taskID, nil
}
