package handlers

import (
	"net/http"
)

type authRequest struct {
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

// authHandler проверяет пароль и возвращает JWT-токен.
func (h Handler) authHandler(w http.ResponseWriter, r *http.Request) {
	var request authRequest
	if err := decodeJSONRequest(r, &request); err != nil {
		writeAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.authentication.Authenticate(request.Password)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, authResponse{Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, authResponse{Token: token})
}
