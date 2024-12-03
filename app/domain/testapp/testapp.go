package testapp

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func test(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
	status := status{
		Status: "OK",
	}

	return status
}
