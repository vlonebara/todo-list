package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/vlonebara/todo-list/pkg/service"
)

type addTaskRequest struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type addTaskResponse struct {
	ID string `json:"id"`
}

// addTaskHandler обрабатывает POST-запрос и возвращает ID задачи или ошибку.
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeAPIError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	var request addTaskRequest
	if err := decodeJSONRequest(r, &request); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := service.AddTask(request.Date, request.Title, request.Comment, request.Repeat)
	if err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, addTaskResponse{ID: strconv.FormatInt(id, 10)})
}

// decodeJSONRequest читает один JSON-объект из тела запроса.
func decodeJSONRequest(r *http.Request, value any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(value); err != nil {
		return errors.New("Некорректный JSON")
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return errors.New("Некорректный JSON")
	}
	return nil
}
