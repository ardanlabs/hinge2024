package apitest

import (
	"github.com/ardanlabs/service/business/domain/userbus"
)

// User extends the dbtest user for api test support.
type User struct {
	userbus.User
}

// SeedData represents users for api tests.
type SeedData struct {
	Users  []User
	Admins []User
}

// Table represent fields needed for running an api test.
type Table struct {
	Name       string
	URL        string
	Method     string
	StatusCode int
	Input      any
	GotResp    any
	ExpResp    any
	CmpFunc    func(got any, exp any) string
}
