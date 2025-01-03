package initialisation

import (
	"fmt"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"server/api/controllers"
	"server/api/middleware"
	"server/api/routes"
	"server/config"
	"server/database"
	"server/database/repository"
	"server/internal/email"
	"server/internal/helpers"
	"server/services"
)

type Initialiser struct {
	c      *config.Config
	authS  *services.AuthService
	adminS *services.AdminService
	db     *repository.Database
	es     *email.EmailService
	mw     *middleware.Middleware
	hs     *helpers.HelperService
}

func NewInitialiser(cfg *config.Config) *Initialiser {
	return &Initialiser{
		c: cfg,
	}
}

func (i *Initialiser) Initialise() (*gin.Engine, error) {
	db, err := database.Init(i.c)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	i.db = &db

	router := gin.Default()

	corsMiddleware := middleware.Cors(i.c)
	var rateLimitMiddleware gin.HandlerFunc
	if i.c.RateLimit.Enabled {
		fmt.Println("Rate limit enabled: limit", i.c.RateLimit.Limit, "window", i.c.RateLimit.Window)
		rateLimitMiddleware = middleware.RateLimit(
			i.c.RateLimit.Limit,
			time.Duration(i.c.RateLimit.Window)*time.Minute,
		)
	}

	sessionStore := sessions.NewCookieStore([]byte(i.c.CookieSecret))
	sessionMiddleware := sessions.Sessions(i.c.SessionName, sessionStore)

	router.Use(corsMiddleware)
	if i.c.RateLimit.Enabled {
		router.Use(rateLimitMiddleware)
	}
	router.Use(sessionMiddleware)

	i.es = email.NewEmailService(i.c)
	i.hs = helpers.NewHelperService(i.db)
	i.mw = middleware.NewMiddleware(i.db, i.hs)
	i.authS = services.NewAuthService(i.c, i.es, i.db, i.hs)
	i.adminS = services.NewAdminService(i.db)

	controllers := controllers.NewControllers(i.c, i.authS, i.adminS, i.db, i.hs)
	appRouter := routes.NewRouter(controllers, i.c, i.db, i.mw)

	appRouter.RegisterRoutes(router)

	return router, nil
}
