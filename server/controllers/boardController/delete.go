package boardController

import (
	"net/http"
	"server/database"
	"server/permissions"

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
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteBoardRequest true "Board ID"
// @Success 200 {object} DeleteBoardResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/delete [post]
func DeleteBoard(c *gin.Context) {
	var req DeleteBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanDeleteBoard, req.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	err := database.DB.BoardRepository.Delete(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete board"})
		return
	}

	c.JSON(http.StatusOK, DeleteBoardResponse{Message: "Board deleted"})
}
