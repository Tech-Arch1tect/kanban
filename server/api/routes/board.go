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
		board.POST("/add-or-invite", r.mw.CSRFTokenRequired(), r.cr.BoardController.AddOrInviteUserToBoard)
		board.GET("/pending-invites/:board_id", r.cr.BoardController.GetPendingInvites)
		board.POST("/remove-pending-invite/:invite_id", r.mw.CSRFTokenRequired(), r.cr.BoardController.RemovePendingInvite)
		board.POST("/remove-user", r.mw.CSRFTokenRequired(), r.cr.BoardController.RemoveUserFromBoard)
		board.POST("/change-role", r.mw.CSRFTokenRequired(), r.cr.BoardController.ChangeBoardRole)
		board.POST("/rename", r.mw.CSRFTokenRequired(), r.cr.BoardController.RenameBoard)
		board.POST("/update-slug", r.mw.CSRFTokenRequired(), r.cr.BoardController.UpdateBoardSlug)
		board.GET("/can-administrate/:board_id", r.cr.BoardController.CanAdministrateBoard)
	}
}
