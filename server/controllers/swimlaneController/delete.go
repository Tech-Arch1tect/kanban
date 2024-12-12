package swimlaneController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type DeleteSwimlaneRequest struct {
	ID uint `json:"id"`
}

type DeleteSwimlaneResponse struct {
	Swimlane models.Swimlane `json:"swimlane"`
}

// @Summary Delete a swimlane
// @Description Delete a swimlane by ID
// @Tags swimlanes
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteSwimlaneRequest true "Swimlane ID"
// @Success 200 {object} DeleteSwimlaneResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/swimlanes/{id} [post]
func DeleteSwimlane(c *gin.Context) {
	var request DeleteSwimlaneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	swimlane, err := database.DB.SwimlaneRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanEditBoard, swimlane.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	err = database.DB.SwimlaneRepository.Delete(swimlane.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeleteSwimlaneResponse{Swimlane: swimlane})
}
