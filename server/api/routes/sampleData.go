package routes

import (
	"server/models"

	"github.com/gin-gonic/gin"
)

func (r *router) RegisterSampleDataRoutes(router *gin.RouterGroup) {
	sampleData := router.Group("/sample-data")
	sampleData.Use(r.mw.AuthRequired())
	sampleData.Use(r.mw.CSRFTokenRequired())
	{
		sampleData.POST("/insert", r.mw.EnsureRole(models.RoleAdmin), r.cr.SampleDataController.InsertSampleData)
	}
}
