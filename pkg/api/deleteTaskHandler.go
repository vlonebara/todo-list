package api

import (
	"net/http"

	"github.com/vlonebara/todo-list/pkg/service"
)

// deleteTaskHandler удаляет задачу с указанным ID.
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeAPIError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	if err := service.DeleteTask(r.URL.Query().Get("id")); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
