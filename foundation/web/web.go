package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Logger func(ctx context.Context, msg string, args ...any)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

type App struct {
	log Logger
	*http.ServeMux
	mw []MidFunc
}

func NewApp(log Logger, mw ...MidFunc) *App {
	return &App{
		log:      log,
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc BILL'S METHOD
func (a *App) HandleFunc(pattern string, handler HandlerFunc, mw ...MidFunc) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTraceID(r.Context(), uuid.NewString())

		dataModel := handler(ctx, r)

		if err := Respond(ctx, w, dataModel); err != nil {
			a.log(ctx, "web-respond", "ERROR", err)
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)
}
