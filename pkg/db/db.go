package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var schema = []string{
	`CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8),
		title TEXT,
		comment TEXT,
		repeat VARCHAR(128) CHECK (repeat IS NULL OR length(repeat) <= 128)
	)`,
	`CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler (date)`,
}

// Create создаёт таблицу и индекс в открытой БД, если их ещё нет.
func Create(database *sql.DB) error {
	for _, statement := range schema {
		if _, err := database.Exec(statement); err != nil {
			return fmt.Errorf("execute schema: %w", err)
		}
	}

	return nil
}
