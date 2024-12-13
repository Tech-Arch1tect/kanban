package commentController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type EditCommentRequest struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

type EditCommentResponse struct {
	Comment models.Comment `json:"comment"`
}

// @Summary Edit a comment
// @Description Edit a comment
// @Tags comments
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body EditCommentRequest true "Edit comment request"
// @Success 200 {object} EditCommentResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/comments/edit [post]
func EditComment(c *gin.Context) {
	var request EditCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanEditComment, request.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	comment, err := database.DB.CommentRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	comment.Text = request.Text
	database.DB.CommentRepository.Update(&comment)

	c.JSON(http.StatusOK, EditCommentResponse{Comment: comment})
}
