package server

import (
	"log"
	"net/http"
	"os"
)

func Run() error {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	addr := ":" + port

	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/", fileServer)

	log.Printf("Server is listening on %s", addr)
	return http.ListenAndServe(addr, nil)
}
