package main

import (
	"server/config"
	"server/controllers/adminController"
	"server/controllers/authController"
	"server/controllers/boardController"
	"server/controllers/columnController"
	"server/controllers/commentController"
	"server/controllers/miscController"
	"server/controllers/sampleDataController"
	"server/controllers/swimlaneController"
	"server/controllers/taskController"
	"server/middleware"
	"server/models"
	"time"

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
			if config.CFG.RateLimit.Enabled {
				auth.POST("/login", middleware.RateLimit(config.CFG.RateLimit.LoginLimit, time.Duration(config.CFG.RateLimit.LoginWindow)*time.Minute), authController.Login)
			} else {
				auth.POST("/login", authController.Login)
			}
			auth.POST("/logout", middleware.AuthRequired(), middleware.CSRFTokenRequired(), authController.Logout)
			auth.GET("/profile", middleware.AuthRequired(), authController.Profile)
			auth.GET("/csrf-token", middleware.AuthRequired(), authController.GetCSRFToken)
			auth.POST("/change-password", middleware.AuthRequired(), middleware.CSRFTokenRequired(), authController.ChangePassword)
			if config.CFG.RateLimit.Enabled {
				auth.POST("/password-reset", middleware.RateLimit(config.CFG.RateLimit.PasswordResetLimit, time.Duration(config.CFG.RateLimit.PasswordResetWindow)*time.Minute), authController.PasswordReset)
				auth.POST("/reset-password", middleware.RateLimit(config.CFG.RateLimit.PasswordResetLimit, time.Duration(config.CFG.RateLimit.PasswordResetWindow)*time.Minute), authController.ResetPassword)
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

		admin := api.Group("admin")
		admin.Use(middleware.EnsureRole(models.RoleAdmin))
		{
			admin.DELETE("/users/:id", middleware.CSRFTokenRequired(), adminController.RemoveUser)
			admin.GET("/users", adminController.ListUsers)
			admin.PUT("/users/:id/role", middleware.CSRFTokenRequired(), adminController.UpdateUserRole)
		}

		sampleData := api.Group("/sample-data")
		sampleData.Use(middleware.EnsureRole(models.RoleAdmin))
		{
			sampleData.POST("/insert", middleware.CSRFTokenRequired(), sampleDataController.InsertSampleData)
		}

		board := api.Group("/boards")
		board.Use(middleware.AuthRequired())
		{
			board.POST("/create", middleware.CSRFTokenRequired(), middleware.EnsureRole(models.RoleAdmin), boardController.CreateBoard)
			board.GET("/get/:id", boardController.GetBoard)
			board.GET("/get-by-slug/:slug", boardController.GetBoardBySlug)
			board.GET("/list", boardController.ListBoards)
			board.POST("/delete", middleware.CSRFTokenRequired(), boardController.DeleteBoard)
		}

		swimlane := api.Group("/swimlanes")
		swimlane.Use(middleware.AuthRequired())
		{
			swimlane.POST("/create", middleware.CSRFTokenRequired(), swimlaneController.CreateSwimlane)
			swimlane.POST("/delete", middleware.CSRFTokenRequired(), swimlaneController.DeleteSwimlane)
			swimlane.POST("/edit", middleware.CSRFTokenRequired(), swimlaneController.EditSwimlane)
			swimlane.POST("/move", middleware.CSRFTokenRequired(), swimlaneController.MoveSwimlane)
		}

		task := api.Group("/tasks")
		task.Use(middleware.AuthRequired())
		{
			task.POST("/create", middleware.CSRFTokenRequired(), taskController.CreateTask)
			task.GET("/get-query/:query", taskController.GetTaskQuery)
			task.GET("/get/:id", taskController.GetTask)
			task.POST("/edit", middleware.CSRFTokenRequired(), taskController.EditTask)
			task.POST("/delete", middleware.CSRFTokenRequired(), taskController.DeleteTask)
			task.POST("/move", middleware.CSRFTokenRequired(), taskController.MoveTask)
		}

		comment := api.Group("/comments")
		comment.Use(middleware.AuthRequired())
		{
			comment.POST("/create", middleware.CSRFTokenRequired(), commentController.CreateComment)
			comment.POST("/edit", middleware.CSRFTokenRequired(), commentController.EditComment)
			comment.POST("/delete", middleware.CSRFTokenRequired(), commentController.DeleteComment)
		}

		column := api.Group("/columns")
		column.Use(middleware.AuthRequired())
		{
			column.POST("/create", middleware.CSRFTokenRequired(), columnController.CreateColumn)
			column.POST("/edit", middleware.CSRFTokenRequired(), columnController.EditColumn)
			column.POST("/delete", middleware.CSRFTokenRequired(), columnController.DeleteColumn)
			column.POST("/move", middleware.CSRFTokenRequired(), columnController.MoveColumn)
		}
	}
}
