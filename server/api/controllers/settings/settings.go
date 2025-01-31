package settings

import (
	"net/http"
	"server/internal/helpers"
	"server/models"
	"server/services/settings"

	"github.com/gin-gonic/gin"
)

type SettingsController struct {
	settingsService *settings.SettingsService
	hs              *helpers.HelperService
}

func NewSettingsController(settingsService *settings.SettingsService, hs *helpers.HelperService) *SettingsController {
	return &SettingsController{settingsService: settingsService, hs: hs}
}

type GetSettingsResponse struct {
	Settings models.Settings `json:"settings"`
}

// @Summary Get settings
// @Description Get settings
// @Tags settings
// @Security cookieAuth
// @Accept json
// @Produce json
// @Success 200 {object} GetSettingsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/settings/get [get]
func (sc *SettingsController) GetSettings(c *gin.Context) {
	user, err := sc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	settings, err := sc.settingsService.GetSettings(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, GetSettingsResponse{Settings: settings})
}

type UpdateSettingsRequest struct {
	Settings models.Settings `json:"settings"`
}

type UpdateSettingsResponse struct {
	Message string `json:"message"`
}

// @Summary Update settings
// @Description Update settings
// @Tags settings
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UpdateSettingsRequest true "Settings"
// @Success 200 {object} UpdateSettingsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/settings/update [post]
func (sc *SettingsController) UpdateSettings(c *gin.Context) {
	user, err := sc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var request UpdateSettingsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = sc.settingsService.UpdateSettings(user.ID, request.Settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Settings updated"})
}
