package commentController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type DeleteCommentRequest struct {
	ID uint `json:"id"`
}

type DeleteCommentResponse struct {
	Comment models.Comment `json:"comment"`
}

// @Summary Delete a comment
// @Description Delete a comment
// @Tags comments
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteCommentRequest true "Delete comment request"
// @Success 200 {object} DeleteCommentResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/comments/delete [post]
func DeleteComment(c *gin.Context) {
	var request DeleteCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanDeleteComment, request.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	comment, err := database.DB.CommentRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.CommentRepository.Delete(comment.ID)

	c.JSON(http.StatusOK, DeleteCommentResponse{Comment: comment})
}
