package boardController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

type ListBoardsResponse struct {
	Boards []models.Board `json:"boards"`
}

// ListBoards godoc
// @Summary List all boards
// @Description List all boards for the current user
// @Tags boards
// @Security cookieAuth
// @Success 200 {object} ListBoardsResponse "Boards"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/boards/list [get]
func ListBoards(c *gin.Context) {
	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	boards, err := database.DB.BoardRepository.GetAllByAccess(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorised"})
		return
	}

	if user.Role == models.RoleAdmin {
		boards, err = database.DB.BoardRepository.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorised"})
			return
		}
	}

	c.JSON(http.StatusOK, ListBoardsResponse{Boards: boards})
}
