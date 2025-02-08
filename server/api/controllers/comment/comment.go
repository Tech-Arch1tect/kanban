package comment

import (
	"net/http"
	"server/database/repository"
	"server/internal/helpers"
	"server/models"
	"server/services/comment"
	"server/services/eventBus"
	"server/services/role"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	cs *comment.CommentService
	hs *helpers.HelperService
	rs *role.RoleService
	db *repository.Database
	ce *eventBus.EventBus[models.Comment]
}

func NewCommentController(cs *comment.CommentService, hs *helpers.HelperService, rs *role.RoleService, db *repository.Database, ce *eventBus.EventBus[models.Comment]) *CommentController {
	return &CommentController{cs: cs, hs: hs, rs: rs, db: db, ce: ce}
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
func (cc *CommentController) CreateComment(c *gin.Context) {
	var request CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	comment, err := cc.cs.CreateComment(user.ID, request.TaskID, request.Text)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	cc.ce.Publish("comment.created", comment, user)

	c.JSON(http.StatusOK, CreateCommentResponse{Comment: comment})
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
func (cc *CommentController) DeleteComment(c *gin.Context) {
	var request DeleteCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	comment, err := cc.cs.DeleteComment(user.ID, request.ID)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	cc.ce.Publish("comment.deleted", comment, user)

	c.JSON(http.StatusOK, DeleteCommentResponse{Comment: comment})
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
func (cc *CommentController) EditComment(c *gin.Context) {
	var request EditCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	comment, err := cc.cs.EditComment(&user, request.ID, request.Text)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	cc.ce.Publish("comment.updated", comment, user)

	c.JSON(http.StatusOK, EditCommentResponse{Comment: comment})
}
