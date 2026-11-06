package db

import (
	"fmt"
	"time"
)

const searchDateLayout = "02.01.2006"

// List возвращает до 15 задач по дате или тексту поиска.
func (s TaskStore) List(search string) ([]Task, error) {
	query, args := tasksListQuery(search)
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("get tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]Task, 0, 15)
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, fmt.Errorf("read task: %w", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read tasks: %w", err)
	}
	return tasks, nil
}

// tasksListQuery строит запрос для списка, даты или текстового поиска.
func tasksListQuery(search string) (string, []any) {
	const fields = `SELECT id, date, title, comment, repeat FROM scheduler`
	const orderAndLimit = ` ORDER BY date, id LIMIT 15`
	if search == "" {
		return fields + orderAndLimit, nil
	}

	date, err := time.Parse(searchDateLayout, search)
	if err == nil && date.Format(searchDateLayout) == search {
		return fields + ` WHERE date = ?` + orderAndLimit, []any{date.Format("20060102")}
	}

	pattern := "%" + search + "%"
	return fields + ` WHERE title LIKE ? OR comment LIKE ?` + orderAndLimit, []any{pattern, pattern}
}
