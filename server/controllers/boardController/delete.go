package boardController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

type DeleteBoardRequest struct {
	ID uint `json:"id" binding:"required"`
}

type DeleteBoardResponse struct {
	Message string `json:"message"`
}

// DeleteBoard godoc
// @Summary Delete a board
// @Description Delete a board by ID
// @Tags boards
// @Accept json
// @Produce json
// @Param request body DeleteBoardRequest true "Board ID"
// @Success 200 {object} DeleteBoardResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/{id} [post]
func DeleteBoard(c *gin.Context) {
	var req DeleteBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	board, err := database.DB.BoardRepository.GetByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Board not found"})
		return
	}

	if board.OwnerID != user.ID && user.Role != models.RoleAdmin {
		permission, err := database.DB.BoardRepository.GetPermission(user.ID, req.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No permission to delete this board"})
			return
		}

		if !permission.Delete {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this board"})
			return
		}
	}

	err = database.DB.BoardRepository.Delete(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete board"})
		return
	}

	c.JSON(http.StatusOK, DeleteBoardResponse{Message: "Board deleted"})
}
