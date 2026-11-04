package service

import (
	"fmt"
	"strings"
	"time"
)

// AddTask validates and saves a task, returning its database ID.
func AddTask(date, title, comment, repeat string) (int64, error) {
	if strings.TrimSpace(title) == "" {
		return 0, fmt.Errorf("Не указан заголовок задачи")
	}

	now := time.Now()
	today := now.Format(dateLayout)
	if date == "" {
		date = today
	}
	if _, err := time.Parse(dateLayout, date); err != nil {
		return 0, fmt.Errorf("Некорректная дата: %w", err)
	}

	if repeat != "" {
		nextDate, err := NextDate(now, date, repeat)
		if err != nil {
			return 0, fmt.Errorf("Некорректное правило повторения: %w", err)
		}
		if date < today {
			date = nextDate
		}
	} else if date < today {
		date = today
	}

	database, err := openDatabase()
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
