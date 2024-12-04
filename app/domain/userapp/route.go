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

	web.HandleFunc("GET /users", app.query)
	web.HandleFunc("POST /users", app.create)
	web.HandleFunc("PUT /users/{user_id}", app.update)
	web.HandleFunc("DELETE /users/{user_id}", app.delete)
	web.HandleFunc("GET /users/{user_id}", app.queryByID)
	web.HandleFunc("PUT /users/role/{user_id}", app.updateRole)
}
