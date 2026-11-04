package api

import (
	"net/http"
	"os"

	"github.com/vlonebara/todo-list/pkg/service"
)

// auth проверяет JWT cookie, если для сервера настроен пароль.
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("TODO_PASSWORD") == "" {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil || !service.ValidateToken(cookie.Value) {
			writeAPIError(w, http.StatusUnauthorized, "Требуется аутентификация")
			return
		}
		next(w, r)
	}
}
