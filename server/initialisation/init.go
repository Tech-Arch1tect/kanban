package initialisation

import (
	"fmt"
	"server/api/controllers"
	"server/api/middleware"
	"server/config"
	"server/database"
	"server/internal/email"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ServerInitialiser interface {
	Initialise() (*gin.Engine, error)
}

type serverInitialiser struct {
	config *config.Config
}

func (si *serverInitialiser) GetDBType() string {
	return si.config.Database.Type
}

func (si *serverInitialiser) GetSQLiteConfig() database.SQLiteConfig {
	return si.config.Database.SQLite
}

func (si *serverInitialiser) GetMySQLConfig() database.MySQLConfig {
	return si.config.Database.MySQL
}

func NewServerInitialiser(cfg *config.Config) ServerInitialiser {
	return &serverInitialiser{config: cfg}
}

func (si *serverInitialiser) Initialise() (*gin.Engine, error) {
	// Initialise database
	if err := database.Init(si); err != nil {
		return nil, fmt.Errorf("failed to initialise database: %v", err)
	}

	// Initialise email
	err := email.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialise email: %v", err)
	}

	r := gin.Default()

	// Add CORS headers
	r.Use(middleware.Cors())

	// Configure rate limiting
	if si.config.RateLimit.Enabled {
		fmt.Printf("Rate limiting enabled with %d requests per %d minutes",
			si.config.RateLimit.Limit, si.config.RateLimit.Window)
		r.Use(middleware.RateLimit(si.config.RateLimit.Limit,
			time.Duration(si.config.RateLimit.Window)*time.Minute))
	}

	// Configure session middleware
	store := sessions.NewCookieStore([]byte(si.config.CookieSecret))
	r.Use(sessions.Sessions(si.config.SessionName, store))

	if err := controllers.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialise controllers: %v", err)
	}

	return r, nil
}
