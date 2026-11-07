package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type apiErrorResponse struct {
	Error string `json:"error"`
}

// writeAPIError возвращает клиенту JSON с описанием ошибки.
func writeAPIError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, apiErrorResponse{Error: message})
}

// writeJSON записывает значение в ответ в формате JSON.
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("write response: %v", err)
	}
}
