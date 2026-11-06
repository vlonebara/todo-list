package service

// DeleteTask удаляет задачу с указанным ID.
func (s TaskService) DeleteTask(id string) error {
	taskID, err := parseTaskID(id)
	if err != nil {
		return err
	}
	return s.store.Delete(taskID)
}
