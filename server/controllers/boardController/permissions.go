package boardController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type GetUsersWithAccessToBoardRequest struct {
	BoardID uint `uri:"board_id" binding:"required"`
}

type GetUsersWithAccessToBoardResponse struct {
	Users []models.User `json:"users"`
}

// @Summary Get users with access to a board
// @Description Get users with access to a board
// @Tags boards
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param board_id path uint true "Board ID"
// @Success 200 {object} GetUsersWithAccessToBoardResponse
// @Router /api/v1/boards/permissions/{board_id} [get]
func GetUsersWithAccessToBoard(c *gin.Context) {
	request := GetUsersWithAccessToBoardRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanAccessBoard, request.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	users, err := database.DB.BoardPermissionRepository.GetUsersWithAccessToBoard(request.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetUsersWithAccessToBoardResponse{Users: users})
}
