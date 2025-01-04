package controllers

import (
	"server/api/controllers/admin"
	"server/api/controllers/auth"
	"server/api/controllers/board"
	"server/api/controllers/column"
	"server/api/controllers/misc"
	"server/api/controllers/swimlane"
	"server/config"
	"server/database/repository"
	"server/internal/helpers"
	"server/services"
)

type Controllers struct {
	AuthController     *auth.AuthController
	AdminController    *admin.AdminController
	MiscController     *misc.MiscController
	BoardController    *board.BoardController
	ColumnController   *column.ColumnController
	SwimlaneController *swimlane.SwimlaneController
}

func NewControllers(cfg *config.Config, authS *services.AuthService, adminS *services.AdminService, db *repository.Database, hs *helpers.HelperService, bs *services.BoardService, rs *services.RoleService, cs *services.ColumnService, ss *services.SwimlaneService) *Controllers {
	return &Controllers{
		AuthController:     auth.NewAuthController(authS, db, hs),
		AdminController:    admin.NewAdminController(adminS),
		MiscController:     misc.NewMiscController(cfg),
		BoardController:    board.NewBoardController(bs, rs, db, hs),
		ColumnController:   column.NewColumnController(db, cs, rs, hs),
		SwimlaneController: swimlane.NewSwimlaneController(db, ss, rs, hs),
	}
}
