package boardController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

type GetBoardRequest struct {
	ID uint `json:"id" binding:"required"`
}

type GetBoardResponse struct {
	Board models.Board `json:"board"`
}

// GetBoard godoc
// @Summary Get a board
// @Description Get a board by ID
// @Tags boards
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param id path int true "Board ID"
// @Success 200 {object} GetBoardResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/{id} [get]
func GetBoard(c *gin.Context) {
	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req GetBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	board, err := database.DB.BoardRepository.GetWithPreload(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorised"})
		return
	}

	if board.OwnerID != user.ID && user.Role != models.RoleAdmin {
		permission, err := database.DB.BoardRepository.GetPermission(user.ID, req.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorised"})
			return
		}

		if permission.UserID != user.ID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorised"})
			return
		}
	}

	c.JSON(http.StatusOK, GetBoardResponse{Board: board})
}
