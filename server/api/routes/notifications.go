package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *router) RegisterNotificationRoutes(router *gin.RouterGroup) {
	notification := router.Group("/notifications")
	notification.Use(r.mw.AuthRequired())
	{
		notification.POST("/create", r.mw.CSRFTokenRequired(), r.cr.NotificationController.CreateNotification)
		notification.GET("/events", r.cr.NotificationController.GetEvents)
		notification.GET("/list", r.cr.NotificationController.GetNotifications)
	}
}
