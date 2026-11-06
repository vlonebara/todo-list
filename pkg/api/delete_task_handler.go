package api

import "net/http"

// deleteTaskHandler удаляет задачу с указанным ID.
func (h apiHandlers) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.tasks.DeleteTask(r.URL.Query().Get("id")); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
