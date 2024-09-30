package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/purisaurabh/database-connection/internal/app"
)

func Routes(ctx context.Context, svc app.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/post", PostDataApi(ctx, svc)).Methods(http.MethodPost)
	router.HandleFunc("/list", ListProfileData(ctx, svc)).Methods(http.MethodGet)
	router.HandleFunc("/profile_update/{id}", UpdateProfileData(ctx, svc)).Methods(http.MethodPut)
	router.HandleFunc("/profile_delete/{id}", DeleteProfileData(ctx, svc)).Methods(http.MethodDelete)
	return router
}
