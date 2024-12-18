package swimlaneController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type MoveSwimlaneRequest struct {
	ID         uint   `json:"id" binding:"required"`
	RelativeID uint   `json:"relative_id" binding:"required"`
	Direction  string `json:"direction" binding:"required,oneof=before after"`
}

type MoveSwimlaneResponse struct {
	Swimlane models.Swimlane `json:"swimlane"`
}

// @Summary Move a swimlane
// @Description Move a swimlane relative to another swimlane
// @Tags swimlanes
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body MoveSwimlaneRequest true "Swimlane ID and relative swimlane ID and direction"
// @Success 200 {object} MoveSwimlaneResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/swimlanes/move [post]
func MoveSwimlane(c *gin.Context) {
	var request MoveSwimlaneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	swimlane, err := database.DB.SwimlaneRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch swimlane"})
		return
	}
	relativeSwimlane, err := database.DB.SwimlaneRepository.GetByID(request.RelativeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch relative swimlane"})
		return
	}

	if swimlane.BoardID != relativeSwimlane.BoardID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Swimlanes are not on the same board"})
		return
	}

	if can, _ := permissions.Can(c, permissions.CanEditBoard, swimlane.BoardID); !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	allBoardSwimlanes, err := database.DB.SwimlaneRepository.GetByBoardID(swimlane.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch swimlanes"})
		return
	}

	swimlaneMap := make(map[uint]int)
	for i, s := range allBoardSwimlanes {
		swimlaneMap[s.ID] = i
	}

	currentIdx, relativeIdx := swimlaneMap[swimlane.ID], swimlaneMap[relativeSwimlane.ID]

	targetIdx := relativeIdx
	if request.Direction == "before" {
		targetIdx++
	}

	allBoardSwimlanes = append(allBoardSwimlanes[:currentIdx], allBoardSwimlanes[currentIdx+1:]...)

	if currentIdx < targetIdx {
		targetIdx--
	}

	allBoardSwimlanes = append(allBoardSwimlanes[:targetIdx], append([]models.Swimlane{swimlane}, allBoardSwimlanes[targetIdx:]...)...)

	for i, s := range allBoardSwimlanes {
		s.Order = i + 1
		if err := database.DB.SwimlaneRepository.Update(&s); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update swimlane order"})
			return
		}
	}

	c.JSON(http.StatusOK, MoveSwimlaneResponse{Swimlane: swimlane})
}

