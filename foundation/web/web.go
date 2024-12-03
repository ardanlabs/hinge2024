package web

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	return &App{
		ServeMux: http.NewServeMux(),
	}
}

// HandleFunc BILL'S METHOD
func (a *App) HandleFunc(pattern string, handler HandlerFunc) {

	h := func(w http.ResponseWriter, r *http.Request) {

		// WE CAN DO WHAT WE WANT HERE

		handler(r.Context(), w, r)

		// WE CAN DO WHAT WE WANT HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
