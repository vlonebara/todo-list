package db

import (
	"database/sql"
	"fmt"
	"os"
)

// OpenDatabase создаёт общий пул соединений и подготавливает БД.
func OpenDatabase() (*sql.DB, error) {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}
	database, err := sql.Open("sqlite", dbFile)
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
