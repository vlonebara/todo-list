package service

// GetTask возвращает задачу с указанным ID.
func (s TaskService) GetTask(id string) (Task, error) {
	taskID, err := parseTaskID(id)
	if err != nil {
		return Task{}, err
	}
	return s.store.Get(taskID)
}
