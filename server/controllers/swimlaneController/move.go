package swimlaneController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"
	"sort"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	swimlane, err := database.DB.SwimlaneRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	relativeSwimlane, err := database.DB.SwimlaneRepository.GetByID(request.RelativeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if relativeSwimlane.BoardID != swimlane.BoardID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Swimlanes are not on the same board"})
		return
	}

	can, _ := permissions.Can(c, permissions.CanEditBoard, swimlane.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	allBoardSwimlanes, err := database.DB.SwimlaneRepository.GetByBoardID(swimlane.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sort.Slice(allBoardSwimlanes, func(i, j int) bool {
		return allBoardSwimlanes[i].Order < allBoardSwimlanes[j].Order
	})

	currentIdx := -1
	relativeIdx := -1
	for i, s := range allBoardSwimlanes {
		if s.ID == swimlane.ID {
			currentIdx = i
		}
		if s.ID == relativeSwimlane.ID {
			relativeIdx = i
		}
	}

	targetIdx := relativeIdx
	if request.Direction == "after" {
		targetIdx++
	}

	if currentIdx < targetIdx {
		for i := currentIdx; i < targetIdx-1; i++ {
			allBoardSwimlanes[i].Order = allBoardSwimlanes[i+1].Order
		}
	} else {
		for i := currentIdx; i > targetIdx; i-- {
			allBoardSwimlanes[i].Order = allBoardSwimlanes[i-1].Order
		}
	}

	for _, s := range allBoardSwimlanes {
		if err := database.DB.SwimlaneRepository.Update(&s); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, MoveSwimlaneResponse{Swimlane: swimlane})
}
