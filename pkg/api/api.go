package api

import "net/http"

// init инициализирует обработчики сервера
func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("GET /api/tasks", getTasksListHandler)
	http.HandleFunc("GET /api/task", getTaskHandler)
	http.HandleFunc("POST /api/task", addTaskHandler)
	http.HandleFunc("DELETE /api/task", deleteTaskHandler)
}
