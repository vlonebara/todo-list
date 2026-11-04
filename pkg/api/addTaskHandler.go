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

	request, err := decodeAddTaskRequest(r)
	if err != nil {
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

// decodeAddTaskRequest читает один JSON-объект с данными задачи.
func decodeAddTaskRequest(r *http.Request) (addTaskRequest, error) {
	var request addTaskRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		return request, errors.New("Некорректный JSON")
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return request, errors.New("Некорректный JSON")
	}
	return request, nil
}
