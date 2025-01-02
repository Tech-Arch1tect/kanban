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
	RegisterRoutes(engine *gin.Engine)
}

type router struct {
	cr  *controllers.Controllers
	cfg *config.Config
}

func NewRouter(cr *controllers.Controllers, cfg *config.Config) Router {
	return &router{
		cr:  cr,
		cfg: cfg,
	}
}

func (r *router) RegisterRoutes(engine *gin.Engine) {
	root := engine.Group("/")
	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api/v1")
	{
		r.registerMiscRoutes(api, r.cr)
		r.registerAuthRoutes(api, r.cr)
		r.registerAdminRoutes(api, r.cr)
	}
}

func (r *router) registerMiscRoutes(api *gin.RouterGroup, cr *controllers.Controllers) {
	misc := api.Group("/misc")
	{
		misc.GET("/appname", cr.MiscController.GetAppName)
	}
}

func (r *router) registerAuthRoutes(api *gin.RouterGroup, cr *controllers.Controllers) {
	auth := api.Group("/auth")
	{
		if r.cfg.RateLimit.Enabled {
			auth.POST("/login", middleware.RateLimit(r.cfg.RateLimit.LoginLimit, time.Duration(r.cfg.RateLimit.LoginWindow)*time.Minute), cr.AuthController.Login)
		} else {
			auth.POST("/login", cr.AuthController.Login)
		}
		auth.POST("/register", cr.AuthController.Register)
		auth.POST("/logout", middleware.CSRFTokenRequired(), middleware.AuthRequired(), cr.AuthController.Logout)
		auth.GET("/profile", middleware.AuthRequired(), cr.AuthController.Profile)
		auth.GET("/csrf-token", middleware.AuthRequired(), cr.AuthController.GetCSRFToken)
		auth.POST("/change-password", middleware.AuthRequired(), middleware.CSRFTokenRequired(), cr.AuthController.ChangePassword)
		if r.cfg.RateLimit.Enabled {
			auth.POST("/password-reset", middleware.RateLimit(r.cfg.RateLimit.PasswordResetLimit, time.Duration(r.cfg.RateLimit.PasswordResetWindow)*time.Minute), cr.AuthController.PasswordReset)
			auth.POST("/reset-password", middleware.RateLimit(r.cfg.RateLimit.PasswordResetLimit, time.Duration(r.cfg.RateLimit.PasswordResetWindow)*time.Minute), cr.AuthController.ResetPassword)
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

func (r *router) registerAdminRoutes(api *gin.RouterGroup, cr *controllers.Controllers) {
	admin := api.Group("admin")
	admin.Use(middleware.EnsureRole(models.RoleAdmin))
	admin.Use(middleware.AuthRequired())
	{
		admin.DELETE("/users/:id", middleware.CSRFTokenRequired(), cr.AdminController.RemoveUser)
		admin.GET("/users", cr.AdminController.ListUsers)
		admin.PUT("/users/:id/role", middleware.CSRFTokenRequired(), cr.AdminController.UpdateUserRole)
	}
}
