package routes

import (
	"server/models"

	"github.com/gin-gonic/gin"
)

func (r *router) RegisterBoardRoutes(router *gin.RouterGroup) {
	board := router.Group("/boards")
	{
		board.POST("/create", r.mw.CSRFTokenRequired(), r.mw.AuthRequired(), r.mw.EnsureRole(models.RoleAdmin), r.cr.BoardController.CreateBoard)
	}
}
