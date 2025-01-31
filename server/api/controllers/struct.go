package controllers

import (
	"server/api/controllers/admin"
	"server/api/controllers/auth"
	"server/api/controllers/board"
	"server/api/controllers/column"
	"server/api/controllers/comment"
	"server/api/controllers/misc"
	"server/api/controllers/sampleData"
	"server/api/controllers/settings"
	"server/api/controllers/swimlane"
	"server/api/controllers/task"
	"server/config"
	"server/database/repository"
	"server/internal/helpers"

	admins "server/services/admin"
	auths "server/services/auth"
	boards "server/services/board"
	columns "server/services/column"
	comments "server/services/comment"
	roles "server/services/role"
	settingss "server/services/settings"
	swimlanes "server/services/swimlane"
	tasks "server/services/task"
)

type Controllers struct {
	AuthController       *auth.AuthController
	AdminController      *admin.AdminController
	MiscController       *misc.MiscController
	BoardController      *board.BoardController
	ColumnController     *column.ColumnController
	SwimlaneController   *swimlane.SwimlaneController
	TaskController       *task.TaskController
	CommentController    *comment.CommentController
	SampleDataController *sampleData.SampleDataController
	SettingsController   *settings.SettingsController
}

func NewControllers(cfg *config.Config, auths *auths.AuthService, adminS *admins.AdminService, db *repository.Database, hs *helpers.HelperService, bs *boards.BoardService, rs *roles.RoleService, cols *columns.ColumnService, ss *swimlanes.SwimlaneService, ts *tasks.TaskService, coms *comments.CommentService, settingsS *settingss.SettingsService) *Controllers {
	return &Controllers{
		AuthController:       auth.NewAuthController(auths, db, hs),
		AdminController:      admin.NewAdminController(adminS),
		MiscController:       misc.NewMiscController(cfg),
		BoardController:      board.NewBoardController(bs, rs, db, hs),
		ColumnController:     column.NewColumnController(db, cols, rs, hs),
		SwimlaneController:   swimlane.NewSwimlaneController(db, ss, rs, hs),
		TaskController:       task.NewTaskController(db, ts, rs, bs, hs),
		CommentController:    comment.NewCommentController(coms, hs, rs, db),
		SampleDataController: sampleData.NewSampleDataController(db, ts, hs),
		SettingsController:   settings.NewSettingsController(settingsS, hs),
	}
}
