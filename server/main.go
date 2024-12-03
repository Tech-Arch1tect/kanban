package main

import (
	"log"
	"server/boot"
	"server/config"
	"server/database"
	"server/email"
	"server/middleware"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @SecurityDefinitions.apikey csrf
// @in header
// @name X-CSRF-Token

func main() {
	config.LoadConfig()
	// Initialize database
	database.Init()
	boot.Boot()
	err := email.Init()
	if err != nil {
		log.Fatalf("Failed to initialize email: %v", err)
	}

	r := gin.Default()
	// add cors headers
	r.Use(middleware.Cors())

	// Session middleware
	store := sessions.NewCookieStore([]byte(config.CFG.CookieSecret))
	r.Use(sessions.Sessions(config.CFG.SessionName, store))

	// Routes
	routes(r)

	r.Run(":8090")
}
