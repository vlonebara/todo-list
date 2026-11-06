package main

import (
	"log"

	"github.com/vlonebara/todo-list/pkg/config"
	"github.com/vlonebara/todo-list/pkg/server"
)

// main загружает конфигурацию и запускает сервер.
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
