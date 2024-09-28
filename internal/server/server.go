package server

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/purisaurabh/database-connection/config"
	"github.com/purisaurabh/database-connection/internal/api"
	"github.com/purisaurabh/database-connection/internal/app"
	"github.com/purisaurabh/database-connection/internal/repository"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func StartApiServer() {
	ctx := context.Background()
	port := config.AppPort()

	server := negroni.Classic()
	repo, err := repository.Init(ctx)
	if err != nil {
		log.Fatal("Error initializing repository:", err)
	}

	service := app.NewService(&repo)

	router := api.Routes(ctx, service)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodPatch},
		AllowedHeaders:   []string{"*"},
	})

	server.Use(corsHandler)
	server.UseHandler(router)

	log.Println("Server is running on port", port)
	server.Run(":" + strconv.Itoa(port))
}
