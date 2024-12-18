package swimlaneController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type CreateSwimlaneRequest struct {
	BoardID uint   `json:"board_id"`
	Name    string `json:"name"`
}

type CreateSwimlaneResponse struct {
	Swimlane models.Swimlane `json:"swimlane"`
}

// @Summary Create a swimlane
// @Description Create a swimlane for a board
// @Tags swimlanes
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateSwimlaneRequest true "Swimlane details"
// @Success 200 {object} CreateSwimlaneResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/swimlanes/create [post]
func CreateSwimlane(c *gin.Context) {
	var request CreateSwimlaneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanEditBoard, request.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	order, err := database.DB.SwimlaneRepository.GetNextOrder(request.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	swimlane := models.Swimlane{
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   order,
	}

	err = database.DB.SwimlaneRepository.Create(&swimlane)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateSwimlaneResponse{Swimlane: swimlane})
}
