package misc

import (
	"net/http"
	"server/config"

	"github.com/gin-gonic/gin"
)

type MiscController struct {
	cfg *config.Config
}

func NewMiscController(cfg *config.Config) *MiscController {
	return &MiscController{cfg: cfg}
}

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
		AppName: m.cfg.AppName,
	}
	c.JSON(http.StatusOK, response)
}
