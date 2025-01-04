package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *router) RegisterColumnRoutes(router *gin.RouterGroup) {
	column := router.Group("/columns")
	column.Use(r.mw.AuthRequired())
	column.Use(r.mw.CSRFTokenRequired())
	{
		column.POST("/create", r.cr.ColumnController.CreateColumn)
		column.POST("/move", r.cr.ColumnController.MoveColumn)
		column.POST("/edit", r.cr.ColumnController.EditColumn)
		column.POST("/delete", r.cr.ColumnController.DeleteColumn)
	}
}
