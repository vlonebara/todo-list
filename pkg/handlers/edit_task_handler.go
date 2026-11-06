package handlers

import (
	"errors"
	"net/http"

	"github.com/vlonebara/todo-list/pkg/model"
)

// editTaskHandler обновляет задачу и возвращает ошибку, если обновление не удалось.
func (h Handler) editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var request taskData
	if err := decodeJSONRequest(r, &request); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.tasks.EditTask(request.ID, request.Date, request.Title, request.Comment, request.Repeat); err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, model.ErrInvalidTask):
			status = http.StatusBadRequest
		case errors.Is(err, model.ErrTaskNotFound):
			status = http.StatusNotFound
		}
		writeAPIError(w, status, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
