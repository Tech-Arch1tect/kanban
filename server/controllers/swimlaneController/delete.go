package swimlaneController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

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

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	swimlane, err := database.DB.SwimlaneRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	board, err := database.DB.BoardRepository.GetByID(swimlane.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if board.OwnerID != user.ID && user.Role != models.RoleAdmin {
		permission, err := database.DB.BoardRepository.GetPermission(user.ID, board.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !permission.Edit {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
	}

	database.DB.SwimlaneRepository.Delete(swimlane.ID)

	c.JSON(http.StatusOK, DeleteSwimlaneResponse{Swimlane: swimlane})
}
