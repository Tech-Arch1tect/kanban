package controllers

import (
	"net/http"
	"server/config"

	"github.com/gin-gonic/gin"
)

type MiscController struct{}

type MiscAppNameResponse struct {
	AppName string `json:"app_name"`
}

// @Summary Get the app name
// @Description Get the app name
// @Tags misc
// @Produce json
// @Success 200 {object} MiscAppNameResponse
// @Router /api/v1/misc/appname [get]
func (m *MiscController) GetAppName(c *gin.Context) {
	response := MiscAppNameResponse{
		AppName: config.CFG.AppName,
	}
	c.JSON(http.StatusOK, response)
}
