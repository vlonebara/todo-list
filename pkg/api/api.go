package api

import "net/http"

// init инициализирует обработчики сервера
func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
}
