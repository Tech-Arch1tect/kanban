package board

import (
	"server/config"
	"server/database/repository"
	"server/internal/email"
	"server/services/role"
)

type BoardService struct {
	db  *repository.Database
	rs  *role.RoleService
	es  *email.EmailService
	cfg *config.Config
}

func NewBoardService(db *repository.Database, rs *role.RoleService, es *email.EmailService, cfg *config.Config) *BoardService {
	return &BoardService{
		db:  db,
		rs:  rs,
		es:  es,
		cfg: cfg,
	}
}
