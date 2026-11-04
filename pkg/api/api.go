package api

import "net/http"

// init инициализирует обработчики сервера
func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("POST /api/signin", authHandler)
	http.HandleFunc("GET /api/tasks", auth(getTasksListHandler))
	http.HandleFunc("GET /api/task", auth(getTaskHandler))
	http.HandleFunc("POST /api/task", auth(addTaskHandler))
	http.HandleFunc("PUT /api/task", auth(editTaskHandler))
	http.HandleFunc("DELETE /api/task", auth(deleteTaskHandler))
	http.HandleFunc("POST /api/task/done", auth(editTaskStatusHandler))
}
