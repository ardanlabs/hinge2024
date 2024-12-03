package mux

import (
	"net/http"

	"github.com/ardanlabs/service/app/domain/testapp"
)

func WebAPI() *http.ServeMux {
	mux := http.NewServeMux()

	testapp.Routes(mux)

	return mux
}
