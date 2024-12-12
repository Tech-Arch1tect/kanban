package swimlaneController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

type EditSwimlaneRequest struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

type EditSwimlaneResponse struct {
	Swimlane models.Swimlane `json:"swimlane"`
}

// @Summary Edit a swimlane
// @Description Edit a swimlane by ID
// @Tags swimlanes
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body EditSwimlaneRequest true "Swimlane details"
// @Success 200 {object} EditSwimlaneResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/swimlanes/{id} [post]
func EditSwimlane(c *gin.Context) {
	var request EditSwimlaneRequest
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

	swimlane.Name = request.Name
	swimlane.Order = request.Order

	err = database.DB.SwimlaneRepository.Update(&swimlane)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, EditSwimlaneResponse{Swimlane: swimlane})
}
