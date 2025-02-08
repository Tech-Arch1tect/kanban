package notification

import (
	"server/database/repository"
	"server/services/role"
)

type NotificationService struct {
	db *repository.Database
	rs *role.RoleService
}

func NewNotificationService(db *repository.Database, rs *role.RoleService) *NotificationService {
	return &NotificationService{db: db, rs: rs}
}
