package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type NotificationEventRepository interface {
	Repository[models.NotificationEvent]
}

type GormNotificationEventRepository struct {
	*GormRepository[models.NotificationEvent]
}

func NewNotificationEventRepository(db *gorm.DB) NotificationEventRepository {
	return &GormNotificationEventRepository{
		GormRepository: NewGormRepository[models.NotificationEvent](db),
	}
}

func (r *GormNotificationEventRepository) Migrate() error {
	return r.db.AutoMigrate(&models.NotificationEvent{})
}
