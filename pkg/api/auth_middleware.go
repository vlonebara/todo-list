package api

import (
	"net/http"

	"github.com/vlonebara/todo-list/pkg/service"
)

// auth проверяет JWT cookie, если для сервера настроен пароль.
func auth(authentication service.AuthService, next http.HandlerFunc) http.HandlerFunc {
	if !authentication.Enabled() {
		return next
	}
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || !authentication.ValidateToken(cookie.Value) {
			writeAPIError(w, http.StatusUnauthorized, "authentication required")
			return
		}
		next(w, r)
	}
}
