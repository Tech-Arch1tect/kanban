package controllers

import (
	"server/api/controllers/admin"
	"server/api/controllers/auth"
	"server/api/controllers/board"
	"server/api/controllers/misc"
	"server/config"
	"server/database/repository"
	"server/internal/helpers"
	"server/services"
)

type Controllers struct {
	AuthController  *auth.AuthController
	AdminController *admin.AdminController
	MiscController  *misc.MiscController
	BoardController *board.BoardController
}

func NewControllers(cfg *config.Config, authS *services.AuthService, adminS *services.AdminService, db *repository.Database, hs *helpers.HelperService, bs *services.BoardService, ps *services.PermissionService) *Controllers {
	return &Controllers{
		AuthController:  auth.NewAuthController(authS, db, hs),
		AdminController: admin.NewAdminController(adminS),
		MiscController:  misc.NewMiscController(cfg),
		BoardController: board.NewBoardController(bs, ps, db, hs),
	}
}
