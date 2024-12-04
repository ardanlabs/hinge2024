package user_test

import (
	"testing"

	"github.com/ardanlabs/service/app/sdk/apitest"
)

func Test_User(t *testing.T) {
	t.Parallel()

	test := apitest.New(t, "Test_User")

	// -------------------------------------------------------------------------

	sd, err := insertSeedData(test.DB)
	if err != nil {
		t.Fatalf("Seeding error: %s", err)
	}

	// -------------------------------------------------------------------------

	test.Run(t, query200(sd), "query-200")
}
