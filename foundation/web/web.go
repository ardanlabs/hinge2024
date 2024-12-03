package web

import (
	"context"
	"net/http"
)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

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

		ctx := r.Context()

		dataModel := handler(ctx, r)

		Respond(ctx, w, dataModel)

		// WE CAN DO WHAT WE WANT HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
