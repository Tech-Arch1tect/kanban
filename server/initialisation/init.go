package initialisation

import (
	"log"
	"server/config"
	"server/database"
	"server/email"
	"server/middleware"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ServerInitialiser interface {
	Initialise() *gin.Engine
}

type serverInitialiser struct {
	config *config.Config
}

func (si *serverInitialiser) GetDBType() string {
	return si.config.DBType
}

func (si *serverInitialiser) GetSQLiteConfig() database.SQLiteConfig {
	return si.config.SQLite
}

func (si *serverInitialiser) GetMySQLConfig() database.MySQLConfig {
	return si.config.MySQL
}

func NewServerInitialiser(cfg *config.Config) ServerInitialiser {
	return &serverInitialiser{config: cfg}
}

func (si *serverInitialiser) Initialise() *gin.Engine {
	// Initialise database
	database.Init(si)

	// Initialise email
	err := email.Init()
	if err != nil {
		log.Fatalf("Failed to initialise email: %v", err)
	}

	r := gin.Default()

	// Add CORS headers
	r.Use(middleware.Cors())

	// Configure rate limiting
	if si.config.RateLimit.Enabled {
		log.Printf("Rate limiting enabled with %d requests per %d minutes",
			si.config.RateLimit.Limit, si.config.RateLimit.Window)
		r.Use(middleware.RateLimit(si.config.RateLimit.Limit,
			time.Duration(si.config.RateLimit.Window)*time.Minute))
	}

	// Configure session middleware
	store := sessions.NewCookieStore([]byte(si.config.CookieSecret))
	r.Use(sessions.Sessions(si.config.SessionName, store))

	return r
}
