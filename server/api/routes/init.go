package routes

import (
	"server/api/controllers"
	"server/api/middleware"
	"server/config"
	"server/database/repository"
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
	db  *repository.Database
	mw  *middleware.Middleware
}

func NewRouter(cr *controllers.Controllers, cfg *config.Config, db *repository.Database, mw *middleware.Middleware) Router {
	return &router{
		cr:  cr,
		cfg: cfg,
		db:  db,
		mw:  mw,
	}
}

func (r *router) RegisterRoutes(engine *gin.Engine) {
	root := engine.Group("/")
	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api/v1")
	{
		r.registerMiscRoutes(api)
		r.registerAuthRoutes(api)
		r.registerAdminRoutes(api)
	}
}

func (r *router) registerMiscRoutes(api *gin.RouterGroup) {
	misc := api.Group("/misc")
	{
		misc.GET("/appname", r.cr.MiscController.GetAppName)
	}
}

func (r *router) registerAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		if r.cfg.RateLimit.Enabled {
			auth.POST("/login", middleware.RateLimit(r.cfg.RateLimit.LoginLimit, time.Duration(r.cfg.RateLimit.LoginWindow)*time.Minute), r.cr.AuthController.Login)
		} else {
			auth.POST("/login", r.cr.AuthController.Login)
		}
		auth.POST("/register", r.cr.AuthController.Register)
		auth.POST("/logout", r.mw.CSRFTokenRequired(), r.mw.AuthRequired(), r.cr.AuthController.Logout)
		auth.GET("/profile", r.mw.AuthRequired(), r.cr.AuthController.Profile)
		auth.GET("/csrf-token", r.mw.AuthRequired(), r.cr.AuthController.GetCSRFToken)
		auth.POST("/change-password", r.mw.AuthRequired(), r.mw.CSRFTokenRequired(), r.cr.AuthController.ChangePassword)
		if r.cfg.RateLimit.Enabled {
			auth.POST("/password-reset", middleware.RateLimit(r.cfg.RateLimit.PasswordResetLimit, time.Duration(r.cfg.RateLimit.PasswordResetWindow)*time.Minute), r.cr.AuthController.PasswordReset)
			auth.POST("/reset-password", middleware.RateLimit(r.cfg.RateLimit.PasswordResetLimit, time.Duration(r.cfg.RateLimit.PasswordResetWindow)*time.Minute), r.cr.AuthController.ResetPassword)
		} else {
			auth.POST("/password-reset", r.cr.AuthController.PasswordReset)
			auth.POST("/reset-password", r.cr.AuthController.ResetPassword)
		}
	}

	authtotpLoginRequired := api.Group("/auth")
	authtotpLoginRequired.Use(r.mw.TOTPTempAuthRequired())
	{
		authtotpLoginRequired.POST("/totp/confirm", r.cr.AuthController.ConfirmTOTP)
	}

	authLoginRequired := api.Group("/auth")
	authLoginRequired.Use(r.mw.AuthRequired())
	{
		authLoginRequired.POST("/totp/generate", r.mw.CSRFTokenRequired(), r.cr.AuthController.GenerateTOTP)
		authLoginRequired.POST("/totp/enable", r.mw.CSRFTokenRequired(), r.cr.AuthController.EnableTOTP)
		authLoginRequired.POST("/totp/disable", r.mw.CSRFTokenRequired(), r.cr.AuthController.DisableTOTP)
	}
}

func (r *router) registerAdminRoutes(api *gin.RouterGroup) {
	admin := api.Group("admin")
	admin.Use(r.mw.EnsureRole(models.RoleAdmin))
	admin.Use(r.mw.AuthRequired())
	{
		admin.DELETE("/users/:id", r.mw.CSRFTokenRequired(), r.cr.AdminController.RemoveUser)
		admin.GET("/users", r.cr.AdminController.ListUsers)
		admin.PUT("/users/:id/role", r.mw.CSRFTokenRequired(), r.cr.AdminController.UpdateUserRole)
	}
}
