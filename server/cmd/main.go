package main

import (
	"log"
	"server/config"
	"server/initialisation"
)

// @title Server API
// @version 1.0.0

// @SecurityDefinitions.apikey csrf
// @in header
// @name X-CSRF-Token

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	initialiser := initialisation.NewInitialiser(cfg)
	router, err := initialiser.Initialise()
	if err != nil {
		log.Fatalf("Initialisation error: %v", err)
	}

	if err := router.Run(":8090"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
