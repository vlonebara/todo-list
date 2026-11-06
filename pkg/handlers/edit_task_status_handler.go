package handlers

import "net/http"

// editTaskStatusHandler помечает задачу выполненной.
func (h Handler) editTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.tasks.EditTaskStatus(r.URL.Query().Get("id")); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
