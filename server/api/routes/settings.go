package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *router) RegisterSettingsRoutes(router *gin.RouterGroup) {
	settings := router.Group("/settings")
	settings.Use(r.mw.AuthRequired())
	{
		settings.POST("/update", r.mw.CSRFTokenRequired(), r.cr.SettingsController.UpdateSettings)
		settings.GET("/get", r.cr.SettingsController.GetSettings)
	}
}
