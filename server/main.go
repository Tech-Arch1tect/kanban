package main

import (
	"log"
	"server/boot"
	"server/config"
	"server/database"
	"server/email"
	"server/middleware"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @title Server API
// @version 1.0.0

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

	// Rate limiting
	if config.CFG.RateLimit.Enabled {
		log.Printf("Rate limiting enabled with %d requests per %d minutes", config.CFG.RateLimit.Limit, config.CFG.RateLimit.Window)
		r.Use(middleware.RateLimit(config.CFG.RateLimit.Limit, time.Duration(config.CFG.RateLimit.Window)*time.Minute))
	}

	// Session middleware
	store := sessions.NewCookieStore([]byte(config.CFG.CookieSecret))
	r.Use(sessions.Sessions(config.CFG.SessionName, store))

	// Routes
	routes(r)

	r.Run(":8090")
}
