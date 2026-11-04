package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vlonebara/todo-list/pkg/db"
	"github.com/vlonebara/todo-list/pkg/server"
)

func main() {
	if err := loadEnvFile(".env"); err != nil {
		log.Fatal("load .env: ", err)
	}

	//проверка наличия бд
	dbFile := "scheduler.db"
	if envDBFile := os.Getenv("TODO_DBFILE"); envDBFile != "" {
		dbFile = envDBFile
	}

	if _, err := os.Stat(dbFile); errors.Is(err, os.ErrNotExist) {
		if err := db.Create(dbFile); err != nil {
			log.Fatal("create database: ", err)
		}
	} else if err != nil {
		log.Fatal("check database file: ", err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// loadEnvFile загружает переменные из файла, не заменяя значения окружения.
func loadEnvFile(file string) error {
	input, err := os.Open(file)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		key = strings.TrimSpace(key)
		if !ok || key == "" {
			return fmt.Errorf("invalid .env entry on line %d", lineNumber)
		}
		if _, exists := os.LookupEnv(key); exists {
			continue
		}
		if err := os.Setenv(key, strings.TrimSpace(value)); err != nil {
			return fmt.Errorf("set %s: %w", key, err)
		}
	}
	return scanner.Err()
}
