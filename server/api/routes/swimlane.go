package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *router) RegisterSwimlaneRoutes(router *gin.RouterGroup) {
	swimlane := router.Group("/swimlanes")
	swimlane.Use(r.mw.AuthRequired())
	swimlane.Use(r.mw.CSRFTokenRequired())
	{
		swimlane.POST("/create", r.cr.SwimlaneController.CreateSwimlane)
		swimlane.POST("/move", r.cr.SwimlaneController.MoveSwimlane)
		swimlane.POST("/rename", r.cr.SwimlaneController.RenameSwimlane)
		swimlane.POST("/delete", r.cr.SwimlaneController.DeleteSwimlane)
	}
}
