package initialisation

import (
	"fmt"
	"server/api/controllers"
	"server/api/middleware"
	"server/api/routes"
	"server/config"
	"server/database"
	"server/services"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Initialiser struct {
	Config       *config.Config
	AuthService  *services.AuthService
	AdminService *services.AdminService
}

func NewInitialiser(cfg *config.Config) *Initialiser {

	return &Initialiser{
		Config: cfg,
	}
}

func (i *Initialiser) Initialise() (*gin.Engine, error) {
	if err := database.Init(i.Config); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	router := gin.Default()

	corsMiddleware := middleware.Cors(i.Config)
	var rateLimitMiddleware gin.HandlerFunc
	if i.Config.RateLimit.Enabled {
		fmt.Println("Rate limit enabled: limit", i.Config.RateLimit.Limit, "window", i.Config.RateLimit.Window)
		rateLimitMiddleware = middleware.RateLimit(
			i.Config.RateLimit.Limit,
			time.Duration(i.Config.RateLimit.Window)*time.Minute,
		)
	}
	sessionStore := sessions.NewCookieStore([]byte(i.Config.CookieSecret))
	sessionMiddleware := sessions.Sessions(i.Config.SessionName, sessionStore)

	router.Use(corsMiddleware)
	if i.Config.RateLimit.Enabled {
		router.Use(rateLimitMiddleware)
	}
	router.Use(sessionMiddleware)

	authService := services.NewAuthService(i.Config)
	adminService := services.NewAdminService()

	controllers := controllers.NewControllers(i.Config, authService, adminService)

	appRouter := routes.NewRouter(controllers, i.Config)

	appRouter.RegisterRoutes(router)

	return router, nil
}
