package swimlane

import (
	"server/database/repository"
	"server/services/role"
)

type SwimlaneService struct {
	db *repository.Database
	rs *role.RoleService
}

func NewSwimlaneService(db *repository.Database, rs *role.RoleService) *SwimlaneService {
	return &SwimlaneService{db: db, rs: rs}
}
