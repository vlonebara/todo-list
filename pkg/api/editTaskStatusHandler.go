package api

import (
	"net/http"

	"github.com/vlonebara/todo-list/pkg/service"
)

// editTaskStatusHandler помечает задачу выполненной.
func editTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeAPIError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	if err := service.EditTaskStatus(r.URL.Query().Get("id")); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
