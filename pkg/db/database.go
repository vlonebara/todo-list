package db

import (
	"database/sql"
	"fmt"
)

// OpenDatabase создаёт общий пул соединений и подготавливает БД.
func OpenDatabase(file string) (*sql.DB, error) {
	database, err := sql.Open("sqlite", file)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// Одно соединение упорядочивает записи SQLite и переиспользуется запросами.
	database.SetMaxOpenConns(1)
	database.SetMaxIdleConns(1)

	if err := database.Ping(); err != nil {
		_ = database.Close()
		return nil, fmt.Errorf("connect database: %w", err)
	}
	if err := Create(database); err != nil {
		_ = database.Close()
		return nil, err
	}
	return database, nil
}
