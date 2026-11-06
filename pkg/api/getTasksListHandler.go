package api

import "net/http"

type getTasksListResponse struct {
	Tasks []taskResponse `json:"tasks"`
}

// getTasksListHandler возвращает ближайшие задачи, отсортированные по дате.
func (h taskHandlers) getTasksListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeAPIError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	tasks, err := h.tasks.GetTasksList(r.URL.Query().Get("search"))
	if err != nil {
		writeAPIError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := getTasksListResponse{Tasks: make([]taskResponse, 0, len(tasks))}
	for _, task := range tasks {
		response.Tasks = append(response.Tasks, taskToResponse(task))
	}
	writeJSON(w, http.StatusOK, response)
}
