package main

import (
	"server/config"
	"server/initialisation"
	"server/routes"
)

// @title Server API
// @version 1.0.0

// @SecurityDefinitions.apikey csrf
// @in header
// @name X-CSRF-Token

func main() {
	config.LoadConfig()

	// Initialise server
	initialiser := initialisation.NewServerInitialiser(config.CFG)
	r := initialiser.Initialise()

	// Initialise routes
	router := routes.NewRouter(config.CFG)
	router.RegisterRoutes(r)

	r.Run(":8090")
}
