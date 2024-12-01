package miscController

import (
	"net/http"
	"server/config"

	"github.com/gin-gonic/gin"
)

type AppNameResponse struct {
	AppName string `json:"app_name"`
}

// @Summary Get the app name
// @Description Get the app name
// @Tags misc
// @Produce json
// @Success 200 {object} AppNameResponse
// @Router /api/v1/misc/appname [get]
func GetAppName(c *gin.Context) {
	response := AppNameResponse{
		AppName: config.CFG.AppName,
	}
	c.JSON(http.StatusOK, response)
}
