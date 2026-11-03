package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	webDir := http.FileServer(http.Dir("./web"))
	http.Handle("/", webDir)

	port := os.Getenv("TODO_PORT")
	//Проверка на пустой порт
	if port == "" {
		port = "7540"
	}
	addr := ":" + port

	log.Printf("Server is listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
