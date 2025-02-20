package routes

import "github.com/gin-gonic/gin"

func (r *router) RegisterCommentRoutes(router *gin.RouterGroup) {
	comment := router.Group("/comments")
	comment.Use(r.mw.AuthRequired())
	comment.Use(r.mw.CSRFTokenRequired())
	{
		comment.POST("/create", r.cr.CommentController.CreateComment)
		comment.POST("/delete", r.cr.CommentController.DeleteComment)
		comment.POST("/edit", r.cr.CommentController.EditComment)
		comment.POST("/create-reaction", r.cr.CommentController.CreateCommentReaction)
		comment.POST("/delete-reaction", r.cr.CommentController.DeleteCommentReaction)
	}
}
