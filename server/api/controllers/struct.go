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

func NewControllers(cfg *config.Config, authService *services.AuthService, adminService *services.AdminService, db *repository.Database, helperService *helpers.HelperService, bs *services.BoardService, ps *services.PermissionService) *Controllers {
	return &Controllers{
		AuthController:  auth.NewAuthController(authService, db, helperService),
		AdminController: admin.NewAdminController(adminService),
		MiscController:  misc.NewMiscController(cfg),
		BoardController: board.NewBoardController(bs, ps, db),
	}
}
