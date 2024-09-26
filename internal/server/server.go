package server

import (
	"log"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/purisaurabh/database-connection/config"
	"github.com/purisaurabh/database-connection/internal/api"
	"github.com/purisaurabh/database-connection/internal/app"
	"github.com/purisaurabh/database-connection/internal/repository"
	"github.com/urfave/negroni"
)

func StartApiServer() {
	port := config.AppPort()

	server := negroni.Classic()
	repo, err := repository.Init()
	if err != nil {
		log.Fatal("Error initializing repository:", err)
	}

	service := app.NewService(repo)
	router := mux.NewRouter()
	api.Routes(router, service)

	server.UseHandler(router)

	log.Println("Server is running on port", port)
	server.Run(":" + strconv.Itoa(port))
}
