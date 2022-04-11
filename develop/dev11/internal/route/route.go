package route

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"wb_l2/develop/dev11/internal/converter"
	"wb_l2/develop/dev11/internal/event"
	"wb_l2/develop/dev11/internal/handlers"
)

type Mux struct {
	mux     *http.ServeMux
	handler *handlers.Handler
}

func NewMux(cache *event.Cache) *Mux {
	return &Mux{
		handler: &handlers.Handler{Cache: cache},
		mux:     http.NewServeMux(),
	}
}

func (m *Mux) SetRoutes() *Logger {
	m.mux.HandleFunc("/create_event", CheckMethod(http.MethodPost, http.HandlerFunc(m.handler.CreateEvent)))
	m.mux.HandleFunc("/update_event", CheckMethod(http.MethodPost, http.HandlerFunc(m.handler.UpdateEvent)))
	m.mux.HandleFunc("/delete_event", CheckMethod(http.MethodPost, http.HandlerFunc(m.handler.DeleteEvent)))
	m.mux.HandleFunc("/events_for_day", CheckMethod(http.MethodGet, http.HandlerFunc(m.handler.GetEventsForDay)))
	m.mux.HandleFunc("/events_for_week", CheckMethod(http.MethodGet, http.HandlerFunc(m.handler.GetEventsForWeek)))
	m.mux.HandleFunc("/events_for_month", CheckMethod(http.MethodGet, http.HandlerFunc(m.handler.GetEventsForMonth)))

	return NewLogger(m.mux)
}

func CheckMethod(method string, handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			handler.ServeHTTP(w, r)
		} else {
			converter.JsonResponse(w, http.StatusMethodNotAllowed,
				fmt.Sprintf("Only %v is allowed", method), true)
		}
	}
}

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
