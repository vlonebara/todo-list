package api

import (
	"net/http"

	"github.com/vlonebara/todo-list/pkg/service"
)

// apiHandlers передаёт готовые сервисы HTTP-обработчикам.
type apiHandlers struct {
	tasks          service.TaskService
	authentication service.AuthService
}

// Init регистрирует обработчики с готовым сервисом задач.
func Init(mux *http.ServeMux, tasks service.TaskService, authentication service.AuthService) {
	handlers := apiHandlers{tasks: tasks, authentication: authentication}
	mux.HandleFunc("/api/nextdate", methods(map[string]http.HandlerFunc{
		http.MethodGet: nextDayHandler,
	}))
	mux.HandleFunc("/api/signin", methods(map[string]http.HandlerFunc{
		http.MethodPost: handlers.authHandler,
	}))
	mux.HandleFunc("/api/tasks", methods(map[string]http.HandlerFunc{
		http.MethodGet: auth(authentication, handlers.getTasksListHandler),
	}))
	mux.HandleFunc("/api/task", methods(map[string]http.HandlerFunc{
		http.MethodGet:    auth(authentication, handlers.getTaskHandler),
		http.MethodPost:   auth(authentication, handlers.addTaskHandler),
		http.MethodPut:    auth(authentication, handlers.editTaskHandler),
		http.MethodDelete: auth(authentication, handlers.deleteTaskHandler),
	}))
	mux.HandleFunc("/api/task/done", methods(map[string]http.HandlerFunc{
		http.MethodPost: auth(authentication, handlers.editTaskStatusHandler),
	}))
}
