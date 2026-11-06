package api

import "net/http"

// getTaskHandler возвращает данные задачи с указанным ID.
func (h apiHandlers) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	task, err := h.tasks.GetTask(r.URL.Query().Get("id"))
	if err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, taskToResponse(task))
}
