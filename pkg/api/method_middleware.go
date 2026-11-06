package api

import (
	"net/http"
	"sort"
	"strings"
)

// methods выбирает обработчик по методу или возвращает 405 со списком Allow.
func methods(handlers map[string]http.HandlerFunc) http.HandlerFunc {
	allowed := make([]string, 0, len(handlers))
	for method := range handlers {
		allowed = append(allowed, method)
	}
	sort.Strings(allowed)
	allow := strings.Join(allowed, ", ")

	return func(w http.ResponseWriter, r *http.Request) {
		if handler, ok := handlers[r.Method]; ok {
			handler(w, r)
			return
		}
		w.Header().Set("Allow", allow)
		writeAPIError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}
