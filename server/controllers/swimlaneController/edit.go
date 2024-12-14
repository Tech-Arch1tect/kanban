package swimlaneController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

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
// @Router /api/v1/swimlanes/edit [post]
func EditSwimlane(c *gin.Context) {
	var request EditSwimlaneRequest
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

	swimlane.Name = request.Name
	swimlane.Order = request.Order

	err = database.DB.SwimlaneRepository.Update(&swimlane)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, EditSwimlaneResponse{Swimlane: swimlane})
}
