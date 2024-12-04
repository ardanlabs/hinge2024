package userapp

import (
	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/business/domain/userbus/stores/userdb"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
	"github.com/jmoiron/sqlx"
)

func Routes(log *logger.Logger, db *sqlx.DB, web *web.App) {
	bus := userbus.NewBusiness(log, userdb.NewStore(log, db))
	app := newApp(bus)

	web.HandleFunc("POST /users", app.create)
}
