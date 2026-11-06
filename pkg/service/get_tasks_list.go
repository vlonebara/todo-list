package service

// GetTasksList возвращает до 15 задач по дате или тексту поиска.
func (s TaskService) GetTasksList(search string) ([]Task, error) {
	return s.store.List(search)
}
