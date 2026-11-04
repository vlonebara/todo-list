package service

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

// openDatabase открывает файл БД, указанный в TODO_DBFILE или по умолчанию.
func openDatabase() (*sql.DB, error) {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}
	return sql.Open("sqlite", dbFile)
}
