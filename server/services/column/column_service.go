package column

import (
	"server/database/repository"
	"server/services/role"
)

type ColumnService struct {
	db *repository.Database
	rs *role.RoleService
}

func NewColumnService(db *repository.Database, rs *role.RoleService) *ColumnService {
	return &ColumnService{db: db, rs: rs}
}
