package commentController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	TaskID uint   `json:"task_id"`
	Text   string `json:"text"`
}

type CreateCommentResponse struct {
	Comment models.Comment `json:"comment"`
}

// @Summary Create a comment
// @Description Create a comment for a task
// @Tags comments
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateCommentRequest true "Comment details"
// @Success 200 {object} CreateCommentResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/comments/create [post]
func CreateComment(c *gin.Context) {
	var request CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := database.DB.TaskRepository.GetByID(request.TaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanCreateComment, task.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	comment := models.Comment{
		TaskID: request.TaskID,
		Text:   request.Text,
		UserID: user.ID,
	}

	database.DB.CommentRepository.Create(&comment)

	c.JSON(http.StatusOK, CreateCommentResponse{Comment: comment})
}
