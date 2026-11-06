package api

import (
	"errors"
	"net/http"

	"github.com/vlonebara/todo-list/pkg/service"
)

// editTaskHandler обновляет задачу и возвращает ошибку, если обновление не удалось.
func (h apiHandlers) editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var request taskData
	if err := decodeJSONRequest(r, &request); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.tasks.EditTask(request.ID, request.Date, request.Title, request.Comment, request.Repeat); err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, service.ErrInvalidTask):
			status = http.StatusBadRequest
		case errors.Is(err, service.ErrTaskNotFound):
			status = http.StatusNotFound
		}
		writeAPIError(w, status, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
