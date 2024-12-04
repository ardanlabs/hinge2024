package apitest

import (
	"testing"

	"github.com/ardanlabs/service/app/sdk/mux"
	"github.com/ardanlabs/service/business/sdk/dbtest"
)

// New initialized the system to run a test.
func New(t *testing.T, testName string) *Test {
	db := dbtest.New(t, testName)

	return &Test{
		DB:  db,
		mux: mux.WebAPI(db.Log, db.DB),
	}
}
