package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type NotificationLogRepository interface {
	Repository[models.NotificationLog]
}

type GormNotificationLogRepository struct {
	*GormRepository[models.NotificationLog]
}

func NewNotificationLogRepository(db *gorm.DB) NotificationLogRepository {
	return &GormNotificationLogRepository{
		GormRepository: NewGormRepository[models.NotificationLog](db),
	}
}

func (r *GormNotificationLogRepository) Migrate() error {
	return r.db.AutoMigrate(&models.NotificationLog{})
}
