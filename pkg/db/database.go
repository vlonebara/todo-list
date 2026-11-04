package db

import (
	"database/sql"
	"os"
)

// OpenDatabase открывает файл БД из TODO_DBFILE или scheduler.db по умолчанию.
func OpenDatabase() (*sql.DB, error) {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}
	return sql.Open("sqlite", dbFile)
}
