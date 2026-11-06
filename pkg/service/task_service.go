package service

import (
	"fmt"
	"strconv"

	"github.com/vlonebara/todo-list/pkg/model"
)

// TaskStore описывает операции хранилища, необходимые сервису задач.
type TaskStore interface {
	Add(task model.Task) (int64, error)
	Get(id int64) (model.Task, error)
	List(filter model.TaskFilter) ([]model.Task, error)
	Update(task model.Task) error
	Delete(id int64) error
	SetDate(id int64, date string) error
}

// TaskService проверяет данные задач и обращается к общему хранилищу.
type TaskService struct {
	store TaskStore
}

// NewTaskService создаёт сервис с готовым хранилищем задач.
func NewTaskService(store TaskStore) TaskService {
	return TaskService{store: store}
}

// parseTaskID проверяет и преобразует ID задачи в число.
func parseTaskID(id string) (int64, error) {
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil || taskID < 1 {
		return 0, fmt.Errorf("%w: invalid task id", model.ErrInvalidTask)
	}
	return taskID, nil
}
