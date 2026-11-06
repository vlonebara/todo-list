package api

import (
	"net/http"

	"github.com/vlonebara/todo-list/pkg/service"
)

// taskHandlers передаёт общий сервис задач HTTP-обработчикам.
type taskHandlers struct {
	tasks service.TaskService
}

// Init регистрирует обработчики с готовым сервисом задач.
func Init(mux *http.ServeMux, tasks service.TaskService) {
	handlers := taskHandlers{tasks: tasks}
	mux.HandleFunc("/api/nextdate", nextDayHandler)
	mux.HandleFunc("POST /api/signin", authHandler)
	mux.HandleFunc("GET /api/tasks", auth(handlers.getTasksListHandler))
	mux.HandleFunc("GET /api/task", auth(handlers.getTaskHandler))
	mux.HandleFunc("POST /api/task", auth(handlers.addTaskHandler))
	mux.HandleFunc("PUT /api/task", auth(handlers.editTaskHandler))
	mux.HandleFunc("DELETE /api/task", auth(handlers.deleteTaskHandler))
	mux.HandleFunc("POST /api/task/done", auth(handlers.editTaskStatusHandler))
}
