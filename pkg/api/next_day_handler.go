package api

import (
	"net/http"
	"time"

	"github.com/vlonebara/todo-list/pkg/service"
)

const dateLayout = "20060102"

// nextDayHandler возвращает ближайшую дату задачи по переданному правилу.
func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	if value := r.FormValue("now"); value != "" {
		parsed, err := time.Parse(dateLayout, value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		now = parsed
	}

	date, err := service.NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, _ = w.Write([]byte(date))
}
