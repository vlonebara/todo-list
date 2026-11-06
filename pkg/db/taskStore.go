package db

import (
	"database/sql"
	"fmt"
)

// Task содержит сохранённые данные задачи.
type Task struct {
	ID      int64
	Date    string
	Title   string
	Comment string
	Repeat  string
}

// TaskStore выполняет запросы к задачам через общий пул соединений.
type TaskStore struct {
	db *sql.DB
}

// NewTaskStore создаёт хранилище с уже открытым пулом соединений.
func NewTaskStore(database *sql.DB) TaskStore {
	return TaskStore{db: database}
}

// Add сохраняет задачу и возвращает её ID.
func (s TaskStore) Add(task Task) (int64, error) {
	result, err := s.db.Exec(
		`INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`,
		task.Date, task.Title, task.Comment, task.Repeat,
	)
	if err != nil {
		return 0, fmt.Errorf("insert task: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get task ID: %w", err)
	}
	return id, nil
}

// Get возвращает задачу с указанным ID.
func (s TaskStore) Get(id int64) (Task, error) {
	var task Task
	err := s.db.QueryRow(
		`SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`, id,
	).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err == sql.ErrNoRows {
		return Task{}, fmt.Errorf("Задача не найдена")
	}
	if err != nil {
		return Task{}, fmt.Errorf("get task: %w", err)
	}
	return task, nil
}

// Update сохраняет новые значения полей задачи.
func (s TaskStore) Update(task Task) error {
	result, err := s.db.Exec(
		`UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`,
		task.Date, task.Title, task.Comment, task.Repeat, task.ID,
	)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}
	return checkTaskChanged(result)
}

// Delete удаляет задачу с указанным ID.
func (s TaskStore) Delete(id int64) error {
	result, err := s.db.Exec(`DELETE FROM scheduler WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}
	return checkTaskChanged(result)
}

// SetDate записывает дату следующего выполнения задачи.
func (s TaskStore) SetDate(id int64, date string) error {
	result, err := s.db.Exec(`UPDATE scheduler SET date = ? WHERE id = ?`, date, id)
	if err != nil {
		return fmt.Errorf("update task date: %w", err)
	}
	return checkTaskChanged(result)
}

// checkTaskChanged проверяет, что запрос затронул существующую задачу.
func checkTaskChanged(result sql.Result) error {
	changed, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check changed task: %w", err)
	}
	if changed == 0 {
		return fmt.Errorf("Задача не найдена")
	}
	return nil
}
