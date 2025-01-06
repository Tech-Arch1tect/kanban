package routes

import (
	"server/models"

	"github.com/gin-gonic/gin"
)

func (r *router) RegisterBoardRoutes(router *gin.RouterGroup) {
	board := router.Group("/boards")
	board.Use(r.mw.AuthRequired())
	{
		board.POST("/create", r.mw.CSRFTokenRequired(), r.mw.EnsureRole(models.RoleAdmin), r.cr.BoardController.CreateBoard)
		board.GET("/get/:id", r.cr.BoardController.GetBoard)
		board.GET("/get-by-slug/:slug", r.cr.BoardController.GetBoardBySlug)
		board.POST("/delete", r.mw.CSRFTokenRequired(), r.mw.EnsureRole(models.RoleAdmin), r.cr.BoardController.DeleteBoard)
		board.GET("/list", r.cr.BoardController.ListBoards)
		board.GET("/permissions/:board_id", r.cr.BoardController.GetUsersWithAccessToBoard)
	}
}
