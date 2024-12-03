package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

type App struct {
	*http.ServeMux
	mw []MidFunc
}

func NewApp(mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc BILL'S METHOD
func (a *App) HandleFunc(pattern string, handler HandlerFunc, mw ...MidFunc) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {

		// WE CAN DO WHAT WE WANT HERE

		ctx := setTraceID(r.Context(), uuid.NewString())

		dataModel := handler(ctx, r)

		Respond(ctx, w, dataModel)

		// WE CAN DO WHAT WE WANT HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
