package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"server/api/controllers"
	"server/api/middleware"
	"server/api/routes"
	"server/config"
	"server/database"
	"server/database/repository"
	"server/internal/email"
	"server/internal/helpers"
	"server/models"
	"server/services/admin"
	"server/services/auth"
	"server/services/board"
	"server/services/column"
	"server/services/comment"
	"server/services/eventBus"
	"server/services/notification"
	"server/services/role"
	"server/services/settings"
	"server/services/swimlane"
	"server/services/task"
)

// @title Server API
// @version 1.0.0

// @SecurityDefinitions.apikey csrf
// @in header
// @name X-CSRF-Token

type Params struct {
	fx.In
	Config               *config.Config
	DB                   *repository.Database
	AuthS                *auth.AuthService
	AdminS               *admin.AdminService
	RoleS                *role.RoleService
	BoardS               *board.BoardService
	ColumnS              *column.ColumnService
	SwimlaneS            *swimlane.SwimlaneService
	TaskS                *task.TaskService
	CommentS             *comment.CommentService
	SettingsS            *settings.SettingsService
	EmailS               *email.EmailService
	Helpers              *helpers.HelperService
	MW                   *middleware.Middleware
	NotifS               *notification.NotificationService
	TaskEventBus         *eventBus.EventBus[models.Task]
	CommentEventBus      *eventBus.EventBus[models.Comment]
	FileEventBus         *eventBus.EventBus[models.File]
	LinkEventBus         *eventBus.EventBus[models.TaskLinks]
	ExternalLinkEventBus *eventBus.EventBus[models.TaskExternalLink]
	NotifSubscriber      *notification.NotificationSubscriber
}

func NewRouter(p Params) (*gin.Engine, error) {
	router := gin.Default()

	corsMiddleware := middleware.Cors(p.Config)
	var rateLimitMiddleware gin.HandlerFunc
	if p.Config.RateLimit.Enabled {
		log.Printf("Rate limit enabled: limit %d, window %d", p.Config.RateLimit.Limit, p.Config.RateLimit.Window)
		rateLimitMiddleware = middleware.RateLimit(
			p.Config.RateLimit.Limit,
			time.Duration(p.Config.RateLimit.Window)*time.Minute,
		)
	}

	sessionStore := sessions.NewCookieStore([]byte(p.Config.CookieSecret))
	sessionMiddleware := sessions.Sessions(p.Config.SessionName, sessionStore)

	router.Use(corsMiddleware)
	if p.Config.RateLimit.Enabled {
		router.Use(rateLimitMiddleware)
	}
	router.Use(sessionMiddleware)
	router.Use(p.MW.EnsureCSRFTokenExistsInSession())

	ctrls := controllers.NewControllers(
		p.Config,
		p.AuthS,
		p.AdminS,
		p.DB,
		p.Helpers,
		p.BoardS,
		p.RoleS,
		p.ColumnS,
		p.SwimlaneS,
		p.TaskS,
		p.CommentS,
		p.SettingsS,
		p.NotifS,
		p.TaskEventBus,
		p.CommentEventBus,
		p.FileEventBus,
		p.LinkEventBus,
		p.ExternalLinkEventBus,
	)
	appRouter := routes.NewRouter(ctrls, p.Config, p.DB, p.MW)
	appRouter.RegisterRoutes(router)

	if err := p.RoleS.SeedRoles(); err != nil {
		return nil, fmt.Errorf("failed to seed roles: %w", err)
	}

	if err := p.NotifS.SeedNotificationEvents(); err != nil {
		return nil, fmt.Errorf("failed to seed notification events: %w", err)
	}

	return router, nil
}

func setupRouter() (*gin.Engine, *config.Config, func()) {
	var (
		router *gin.Engine
		cfg    *config.Config
	)

	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			database.Init,
			email.NewEmailService,
			helpers.NewHelperService,
			middleware.NewMiddleware,
			auth.NewAuthService,
			admin.NewAdminService,
			role.NewRoleService,
			board.NewBoardService,
			column.NewColumnService,
			swimlane.NewSwimlaneService,
			task.NewTaskService,
			comment.NewCommentService,
			settings.NewSettingsService,
			notification.NewNotificationService,
			eventBus.NewTaskEventBus,
			eventBus.NewCommentEventBus,
			eventBus.NewFileEventBus,
			eventBus.NewTaskLinkEventBus,
			eventBus.NewTaskExternalLinkEventBus,
			notification.NewNotificationSubscriber,
			NewRouter,
		),
		fx.Populate(&router, &cfg),
		fx.Invoke(func(ns *notification.NotificationSubscriber) {
			go ns.Subscribe()
		}),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start fx app: %v", err)
	}

	cleanup := func() {
		if err := app.Stop(context.Background()); err != nil {
			log.Fatalf("Failed to stop fx app: %v", err)
		}
	}

	return router, cfg, cleanup
}

func main() {
	router, _, cleanup := setupRouter()
	defer cleanup()

	log.Printf("Starting server on %s", "8090")
	if err := router.Run(":8090"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
