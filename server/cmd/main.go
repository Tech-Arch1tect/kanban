package main

import (
	"log"
	"server/cmd/initHelper"
)

// @title Server API
// @version 1.0.0

// @SecurityDefinitions.apikey csrf
// @in header
// @name X-CSRF-Token

func main() {
	router, _, cleanup := initHelper.SetupRouter()
	defer cleanup()

	log.Printf("Starting server on %s", "8090")
	if err := router.Run(":8090"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
