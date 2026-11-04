package api

import (
	"net/http"
	"strconv"

	"github.com/vlonebara/todo-list/pkg/service"
)

type taskResponse struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// taskToResponse преобразует задачу сервиса в JSON-формат API.
func taskToResponse(task service.Task) taskResponse {
	return taskResponse{
		ID:      strconv.FormatInt(task.ID, 10),
		Date:    task.Date,
		Title:   task.Title,
		Comment: task.Comment,
		Repeat:  task.Repeat,
	}
}

// getTaskHandler возвращает данные задачи с указанным ID.
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeAPIError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	task, err := service.GetTask(r.URL.Query().Get("id"))
	if err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, taskToResponse(task))
}
