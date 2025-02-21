package board

import (
	"server/config"
	"server/database/repository"
	"server/internal/email"
	"server/services/column"
	"server/services/role"
	"server/services/swimlane"
	"server/services/task"
)

type BoardService struct {
	db  *repository.Database
	rs  *role.RoleService
	es  *email.EmailService
	cfg *config.Config
	ss  *swimlane.SwimlaneService
	cs  *column.ColumnService
	ts  *task.TaskService
}

func NewBoardService(db *repository.Database, rs *role.RoleService, es *email.EmailService, cfg *config.Config, ss *swimlane.SwimlaneService, cs *column.ColumnService, ts *task.TaskService) *BoardService {
	return &BoardService{
		db:  db,
		rs:  rs,
		es:  es,
		cfg: cfg,
		ss:  ss,
		cs:  cs,
		ts:  ts,
	}
}
