package main

import (
	"log"

	"github.com/vlonebara/todo-list/server"
)

func main() {
	log.Fatal(server.Run())
}
