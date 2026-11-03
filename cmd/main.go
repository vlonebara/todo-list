package main

import (
	"errors"
	"log"
	"os"

	"github.com/vlonebara/todo-list/pkg/db"
	"github.com/vlonebara/todo-list/pkg/server"
)

func main() {
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

	log.Fatal(server.Run())
}
