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
	RegisterRoutes(r *gin.Engine, cr *controllers.Controllers)
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

func (r *router) RegisterRoutes(engine *gin.Engine, cr *controllers.Controllers) {
	root := engine.Group("/")
	registerSwaggerRoutes(root)

	// API versioned routes
	api := engine.Group("/api/v1")
	{
		registerMiscRoutes(api, cr)
		registerAuthRoutes(api, cr, r.config)
		registerAdminRoutes(api, cr)
	}
}

func registerMiscRoutes(api *gin.RouterGroup, cr *controllers.Controllers) {
	misc := api.Group("/misc")
	{
		misc.GET("/appname", cr.MiscController.GetAppName)
	}
}

func registerAuthRoutes(api *gin.RouterGroup, cr *controllers.Controllers, cfg *config.Config) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", cr.AuthController.Register)
		if cfg.RateLimit.Enabled {
			auth.POST("/login", middleware.RateLimit(cfg.RateLimit.LoginLimit, time.Duration(cfg.RateLimit.LoginWindow)*time.Minute), cr.AuthController.Login)
		} else {
			auth.POST("/login", cr.AuthController.Login)
		}
		auth.POST("/logout", middleware.AuthRequired(), middleware.CSRFTokenRequired(), cr.AuthController.Logout)
		auth.GET("/profile", middleware.AuthRequired(), cr.AuthController.Profile)
		auth.GET("/csrf-token", middleware.AuthRequired(), cr.AuthController.GetCSRFToken)
		auth.POST("/change-password", middleware.AuthRequired(), middleware.CSRFTokenRequired(), cr.AuthController.ChangePassword)
		if cfg.RateLimit.Enabled {
			auth.POST("/password-reset", middleware.RateLimit(cfg.RateLimit.PasswordResetLimit, time.Duration(cfg.RateLimit.PasswordResetWindow)*time.Minute), cr.AuthController.PasswordReset)
			auth.POST("/reset-password", middleware.RateLimit(cfg.RateLimit.PasswordResetLimit, time.Duration(cfg.RateLimit.PasswordResetWindow)*time.Minute), cr.AuthController.ResetPassword)
		} else {
			auth.POST("/password-reset", cr.AuthController.PasswordReset)
			auth.POST("/reset-password", cr.AuthController.ResetPassword)
		}
	}

	authtotpLoginRequired := api.Group("/auth")
	authtotpLoginRequired.Use(middleware.TOTPTempAuthRequired())
	{
		authtotpLoginRequired.POST("/totp/confirm", cr.AuthController.ConfirmTOTP)
	}

	authLoginRequired := api.Group("/auth")
	authLoginRequired.Use(middleware.AuthRequired())
	{
		authLoginRequired.POST("/totp/generate", middleware.CSRFTokenRequired(), cr.AuthController.GenerateTOTP)
		authLoginRequired.POST("/totp/enable", middleware.CSRFTokenRequired(), cr.AuthController.EnableTOTP)
		authLoginRequired.POST("/totp/disable", middleware.CSRFTokenRequired(), cr.AuthController.DisableTOTP)
	}
}

func registerAdminRoutes(api *gin.RouterGroup, cr *controllers.Controllers) {
	admin := api.Group("admin")
	admin.Use(middleware.EnsureRole(models.RoleAdmin))
	{
		admin.DELETE("/users/:id", middleware.CSRFTokenRequired(), cr.AdminController.RemoveUser)
		admin.GET("/users", cr.AdminController.ListUsers)
		admin.PUT("/users/:id/role", middleware.CSRFTokenRequired(), cr.AdminController.UpdateUserRole)
	}
}
