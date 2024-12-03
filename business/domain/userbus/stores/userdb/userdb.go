package userdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/foundation/logger"
)

type Store struct {
	log *logger.Logger
	db  sqlx.ExtContext
}

func NewStore(log *logger.Logger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

func (s *Store) Create(ctx context.Context, usr userbus.User) error {
	const q = `
	INSERT INTO users
		(user_id, name, email, password_hash, roles, department, enabled, date_created, date_updated)
	VALUES
		(:user_id, :name, :email, :password_hash, :roles, :department, :enabled, :date_created, :date_updated)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBUser(usr)); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", userbus.ErrUniqueEmail)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}
