package notification

import (
	"server/database/repository"
)

type NotificationService struct {
	db *repository.Database
}

func NewNotificationService(db *repository.Database) *NotificationService {
	return &NotificationService{db: db}
}
