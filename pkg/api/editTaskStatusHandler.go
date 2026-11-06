package api

import "net/http"

// editTaskStatusHandler помечает задачу выполненной.
func (h taskHandlers) editTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeAPIError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	if err := h.tasks.EditTaskStatus(r.URL.Query().Get("id")); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
