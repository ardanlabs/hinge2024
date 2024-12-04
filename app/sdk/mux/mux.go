package mux

import (
	"context"

	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/app/domain/userapp"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
	"github.com/jmoiron/sqlx"
)

func WebAPI(log *logger.Logger, db *sqlx.DB) *web.App {
	l := func(ctx context.Context, msg string, args ...any) {
		log.Info(ctx, msg, args...)
	}

	app := web.NewApp(l, mid.Logger(log), mid.Error(log), mid.Panics())

	testapp.Routes(app)
	userapp.Routes(log, db, app)

	return app
}
