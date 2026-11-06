package service

import "time"

// EditTask проверяет и сохраняет новые данные задачи.
func (s TaskService) EditTask(id, date, title, comment, repeat string) error {
	taskID, err := parseTaskID(id)
	if err != nil {
		return err
	}

	date, err = prepareTaskDate(date, title, repeat, time.Now())
	if err != nil {
		return err
	}

	return s.store.Update(Task{ID: taskID, Date: date, Title: title, Comment: comment, Repeat: repeat})
}
