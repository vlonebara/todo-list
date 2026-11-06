package api

import "net/http"

type editTaskRequest struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// editTaskHandler обновляет задачу и возвращает ошибку, если обновление не удалось.
func (h taskHandlers) editTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeAPIError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	var request editTaskRequest
	if err := decodeJSONRequest(r, &request); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.tasks.EditTask(request.ID, request.Date, request.Title, request.Comment, request.Repeat); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
