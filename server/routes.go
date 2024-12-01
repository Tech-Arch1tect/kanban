package main

import (
	"server/controllers/adminController"
	"server/controllers/authController"
	"server/controllers/miscController"
	"server/middleware"
	"server/models"

	_ "server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func routes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api/v1")
	{
		misc := api.Group("/misc")
		{
			misc.GET("/appname", miscController.GetAppName)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.POST("/logout", middleware.AuthRequired(), authController.Logout)
			auth.GET("/profile", middleware.AuthRequired(), authController.Profile)
			auth.POST("/change-password", middleware.AuthRequired(), authController.ChangePassword)
			auth.POST("/password-reset", authController.PasswordReset)
			auth.POST("/reset-password", authController.ResetPassword)
		}

		authtotpLoginRequired := api.Group("/auth")
		authtotpLoginRequired.Use(middleware.TOTPTempAuthRequired())
		{
			authtotpLoginRequired.POST("/totp/confirm", authController.ConfirmTOTP)
		}

		authLoginRequired := api.Group("/auth")
		authLoginRequired.Use(middleware.AuthRequired())
		{
			authLoginRequired.POST("/totp/generate", authController.GenerateTOTP)
			authLoginRequired.POST("/totp/enable", authController.EnableTOTP)
			authLoginRequired.POST("/totp/disable", authController.DisableTOTP)
		}

		admin := api.Group("admin")
		admin.Use(middleware.EnsureRole(models.RoleAdmin))
		{
			admin.DELETE("/users/:id", adminController.RemoveUser)
			admin.GET("/users", adminController.ListUsers)
			admin.PUT("/users/:id/role", adminController.UpdateUserRole)
		}
	}
}
