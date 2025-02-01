package role

import (
	"server/database/repository"
)

type AppRole struct {
	Name        string
	AccessLevel int
}

var (
	AdminRole  = AppRole{Name: "admin", AccessLevel: 999}
	MemberRole = AppRole{Name: "member", AccessLevel: 200}
	ReaderRole = AppRole{Name: "reader", AccessLevel: 100}
)

var roleMap = map[string]AppRole{
	AdminRole.Name:  AdminRole,
	MemberRole.Name: MemberRole,
	ReaderRole.Name: ReaderRole,
}

type RoleService struct {
	db *repository.Database
}

func NewRoleService(db *repository.Database) *RoleService {
	return &RoleService{db: db}
}
