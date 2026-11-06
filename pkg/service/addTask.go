package service

import (
	"fmt"
	"strings"
	"time"
)

// AddTask проверяет и сохраняет задачу, возвращая её ID.
func (s TaskService) AddTask(date, title, comment, repeat string) (int64, error) {
	date, err := prepareTaskDate(date, title, repeat, time.Now())
	if err != nil {
		return 0, err
	}

	return s.store.Add(Task{Date: date, Title: title, Comment: comment, Repeat: repeat})
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
