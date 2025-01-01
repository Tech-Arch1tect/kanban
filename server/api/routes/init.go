package routes

import (
	"server/api/controllers"
	"server/api/middleware"
	"server/config"
	"server/models"
	"time"

	_ "server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	RegisterRoutes(r *gin.Engine)
}

type router struct {
	config *config.Config
}

func NewRouter(cfg *config.Config) Router {
	return &router{config: cfg}
}

func registerSwaggerRoutes(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *router) RegisterRoutes(engine *gin.Engine) {
	root := engine.Group("/")
	registerSwaggerRoutes(root)

	// API versioned routes
	api := engine.Group("/api/v1")
	{
		registerMiscRoutes(api)
		registerAuthRoutes(api, r.config)
		registerAdminRoutes(api)
	}
}

func registerMiscRoutes(api *gin.RouterGroup) {
	misc := api.Group("/misc")
	{
		misc.GET("/appname", controllers.MiscGetAppName)
	}
}

func registerAuthRoutes(api *gin.RouterGroup, cfg *config.Config) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.AuthRegister)
		if cfg.RateLimit.Enabled {
			auth.POST("/login", middleware.RateLimit(cfg.RateLimit.LoginLimit, time.Duration(cfg.RateLimit.LoginWindow)*time.Minute), controllers.AuthLogin)
		} else {
			auth.POST("/login", controllers.AuthLogin)
		}
		auth.POST("/logout", middleware.AuthRequired(), middleware.CSRFTokenRequired(), controllers.AuthLogout)
		auth.GET("/profile", middleware.AuthRequired(), controllers.AuthProfile)
		auth.GET("/csrf-token", middleware.AuthRequired(), controllers.AuthGetCSRFToken)
		auth.POST("/change-password", middleware.AuthRequired(), middleware.CSRFTokenRequired(), controllers.AuthChangePassword)
		if cfg.RateLimit.Enabled {
			auth.POST("/password-reset", middleware.RateLimit(cfg.RateLimit.PasswordResetLimit, time.Duration(cfg.RateLimit.PasswordResetWindow)*time.Minute), controllers.AuthPasswordReset)
			auth.POST("/reset-password", middleware.RateLimit(cfg.RateLimit.PasswordResetLimit, time.Duration(cfg.RateLimit.PasswordResetWindow)*time.Minute), controllers.AuthResetPassword)
		} else {
			auth.POST("/password-reset", controllers.AuthPasswordReset)
			auth.POST("/reset-password", controllers.AuthResetPassword)
		}
	}

	authtotpLoginRequired := api.Group("/auth")
	authtotpLoginRequired.Use(middleware.TOTPTempAuthRequired())
	{
		authtotpLoginRequired.POST("/totp/confirm", controllers.AuthConfirmTOTP)
	}

	authLoginRequired := api.Group("/auth")
	authLoginRequired.Use(middleware.AuthRequired())
	{
		authLoginRequired.POST("/totp/generate", middleware.CSRFTokenRequired(), controllers.AuthGenerateTOTP)
		authLoginRequired.POST("/totp/enable", middleware.CSRFTokenRequired(), controllers.AuthEnableTOTP)
		authLoginRequired.POST("/totp/disable", middleware.CSRFTokenRequired(), controllers.AuthDisableTOTP)
	}
}

func registerAdminRoutes(api *gin.RouterGroup) {
	admin := api.Group("admin")
	admin.Use(middleware.EnsureRole(models.RoleAdmin))
	{
		admin.DELETE("/users/:id", middleware.CSRFTokenRequired(), controllers.AdminRemoveUser)
		admin.GET("/users", controllers.AdminListUsers)
		admin.PUT("/users/:id/role", middleware.CSRFTokenRequired(), controllers.AdminUpdateUserRole)
	}
}
