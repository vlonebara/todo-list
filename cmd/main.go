package main

import (
	"log"
	"net/http"
)

func main() {
	webDir := http.FileServer(http.Dir("./web"))
	http.Handle("/", webDir)

	log.Println("Server is listening on :7540")
	log.Fatal(http.ListenAndServe(":7540", nil))
}
