package swimlaneController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

type CreateSwimlaneRequest struct {
	BoardID uint   `json:"board_id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
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
// @Router /api/v1/swimlanes [post]
func CreateSwimlane(c *gin.Context) {
	var request CreateSwimlaneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	board, err := database.DB.BoardRepository.GetByID(request.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if board.OwnerID != user.ID && user.Role != models.RoleAdmin {
		permission, err := database.DB.BoardRepository.GetPermission(user.ID, request.BoardID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if !permission.Edit {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
	}

	swimlane := models.Swimlane{
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   request.Order,
	}

	err = database.DB.SwimlaneRepository.Create(&swimlane)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateSwimlaneResponse{Swimlane: swimlane})
}
