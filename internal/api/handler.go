package api

import (
	"net/http"

	"github.com/purisaurabh/database-connection/internal/app"
)

func PostDataApi(service app.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
