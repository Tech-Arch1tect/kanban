package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *router) RegisterTaskRoutes(router *gin.RouterGroup) {
	task := router.Group("/tasks")
	task.Use(r.mw.AuthRequired())
	{
		task.POST("/create", r.mw.CSRFTokenRequired(), r.cr.TaskController.CreateTask)
		task.GET("/get-query/:board_id/:query", r.cr.TaskController.GetTaskQuery)
		task.GET("/get/:id", r.cr.TaskController.GetTask)
		task.POST("/delete", r.mw.CSRFTokenRequired(), r.cr.TaskController.DeleteTask)
		task.POST("/move", r.mw.CSRFTokenRequired(), r.cr.TaskController.MoveTask)
		task.POST("/update/title", r.mw.CSRFTokenRequired(), r.cr.TaskController.UpdateTaskTitle)
		task.POST("/update/description", r.mw.CSRFTokenRequired(), r.cr.TaskController.UpdateTaskDescription)
		task.POST("/update/status", r.mw.CSRFTokenRequired(), r.cr.TaskController.UpdateTaskStatus)
		task.POST("/update/assignee", r.mw.CSRFTokenRequired(), r.cr.TaskController.UpdateTaskAssignee)
		task.GET("/download/:file_id", r.cr.TaskController.DownloadFile)
		task.GET("/get-image/:file_id", r.cr.TaskController.GetImage)
		task.POST("/upload", r.mw.CSRFTokenRequired(), r.cr.TaskController.UploadFile)
		task.POST("/delete-file", r.mw.CSRFTokenRequired(), r.cr.TaskController.DeleteFile)
		task.POST("/create-link", r.mw.CSRFTokenRequired(), r.cr.TaskController.CreateTaskLink)
		task.POST("/delete-link", r.mw.CSRFTokenRequired(), r.cr.TaskController.DeleteTaskLink)
		task.POST("/query-all-boards", r.mw.CSRFTokenRequired(), r.cr.TaskController.QueryAllBoards)
		task.POST("/update-external-link", r.mw.CSRFTokenRequired(), r.cr.TaskController.UpdateTaskExternalLink)
		task.POST("/delete-external-link", r.mw.CSRFTokenRequired(), r.cr.TaskController.DeleteTaskExternalLink)
		task.POST("/create-external-link", r.mw.CSRFTokenRequired(), r.cr.TaskController.CreateTaskExternalLink)
	}
}
