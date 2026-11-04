package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/vlonebara/todo-list/pkg/db"
)

// AddTask validates and saves a task, returning its database ID.
func AddTask(date, title, comment, repeat string) (int64, error) {
	date, err := prepareTaskDate(date, title, repeat, time.Now())
	if err != nil {
		return 0, err
	}

	database, err := db.OpenDatabase()
	if err != nil {
		return 0, fmt.Errorf("open database: %w", err)
	}
	defer database.Close()

	result, err := database.Exec(
		`INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`,
		date, title, comment, repeat,
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

// prepareTaskDate проверяет поля задачи и исправляет просроченную дату.
func prepareTaskDate(date, title, repeat string, now time.Time) (string, error) {
	if strings.TrimSpace(title) == "" {
		return "", fmt.Errorf("Не указан заголовок задачи")
	}

	today := now.Format(dateLayout)
	if date == "" {
		date = today
	}
	if _, err := time.Parse(dateLayout, date); err != nil {
		return "", fmt.Errorf("Некорректная дата: %w", err)
	}

	if repeat != "" {
		nextDate, err := NextDate(now, date, repeat)
		if err != nil {
			return "", fmt.Errorf("Некорректное правило повторения: %w", err)
		}
		if date < today {
			date = nextDate
		}
	} else if date < today {
		date = today
	}
	return date, nil
}
