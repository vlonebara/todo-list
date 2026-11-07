package service

import "github.com/vlonebara/todo-list/pkg/model"

// GetTask возвращает задачу с указанным ID.
func (s TaskService) GetTask(id string) (model.Task, error) {
	taskID, err := parseTaskID(id)
	if err != nil {
		return model.Task{}, err
	}
	return s.store.Get(taskID)
}
