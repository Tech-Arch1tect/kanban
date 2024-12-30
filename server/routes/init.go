package routes

import (
	"server/config"
	"server/controllers/adminController"
	"server/controllers/authController"
	"server/controllers/miscController"
	"server/middleware"
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
		misc.GET("/appname", miscController.GetAppName)
	}
}

func registerAuthRoutes(api *gin.RouterGroup, cfg *config.Config) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		if cfg.RateLimit.Enabled {
			auth.POST("/login", middleware.RateLimit(cfg.RateLimit.LoginLimit, time.Duration(cfg.RateLimit.LoginWindow)*time.Minute), authController.Login)
		} else {
			auth.POST("/login", authController.Login)
		}
		auth.POST("/logout", middleware.AuthRequired(), middleware.CSRFTokenRequired(), authController.Logout)
		auth.GET("/profile", middleware.AuthRequired(), authController.Profile)
		auth.GET("/csrf-token", middleware.AuthRequired(), authController.GetCSRFToken)
		auth.POST("/change-password", middleware.AuthRequired(), middleware.CSRFTokenRequired(), authController.ChangePassword)
		if cfg.RateLimit.Enabled {
			auth.POST("/password-reset", middleware.RateLimit(cfg.RateLimit.PasswordResetLimit, time.Duration(cfg.RateLimit.PasswordResetWindow)*time.Minute), authController.PasswordReset)
			auth.POST("/reset-password", middleware.RateLimit(cfg.RateLimit.PasswordResetLimit, time.Duration(cfg.RateLimit.PasswordResetWindow)*time.Minute), authController.ResetPassword)
		} else {
			auth.POST("/password-reset", authController.PasswordReset)
			auth.POST("/reset-password", authController.ResetPassword)
		}
	}

	authtotpLoginRequired := api.Group("/auth")
	authtotpLoginRequired.Use(middleware.TOTPTempAuthRequired())
	{
		authtotpLoginRequired.POST("/totp/confirm", authController.ConfirmTOTP)
	}

	authLoginRequired := api.Group("/auth")
	authLoginRequired.Use(middleware.AuthRequired())
	{
		authLoginRequired.POST("/totp/generate", middleware.CSRFTokenRequired(), authController.GenerateTOTP)
		authLoginRequired.POST("/totp/enable", middleware.CSRFTokenRequired(), authController.EnableTOTP)
		authLoginRequired.POST("/totp/disable", middleware.CSRFTokenRequired(), authController.DisableTOTP)
	}
}

func registerAdminRoutes(api *gin.RouterGroup) {
	admin := api.Group("admin")
	admin.Use(middleware.EnsureRole(models.RoleAdmin))
	{
		admin.DELETE("/users/:id", middleware.CSRFTokenRequired(), adminController.RemoveUser)
		admin.GET("/users", adminController.ListUsers)
		admin.PUT("/users/:id/role", middleware.CSRFTokenRequired(), adminController.UpdateUserRole)
	}
}
