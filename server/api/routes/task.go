package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *router) RegisterTaskRoutes(router *gin.RouterGroup) {
	task := router.Group("/task")
	task.Use(r.mw.AuthRequired())
	{
		task.POST("/create", r.mw.CSRFTokenRequired(), r.cr.TaskController.CreateTask)
		task.GET("/get-query/:query", r.cr.TaskController.GetTaskQuery)
		task.GET("/get/:id", r.cr.TaskController.GetTask)
		task.POST("/edit", r.mw.CSRFTokenRequired(), r.cr.TaskController.EditTask)
		task.POST("/delete", r.mw.CSRFTokenRequired(), r.cr.TaskController.DeleteTask)
		task.POST("/move", r.mw.CSRFTokenRequired(), r.cr.TaskController.MoveTask)
	}
}
