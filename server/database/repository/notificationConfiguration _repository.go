package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type NotificationConfigurationRepository interface {
	Repository[models.NotificationConfiguration]
}

type GormNotificationConfigurationRepository struct {
	*GormRepository[models.NotificationConfiguration]
}

func NewNotificationConfigurationRepository(db *gorm.DB) NotificationConfigurationRepository {
	return &GormNotificationConfigurationRepository{
		GormRepository: NewGormRepository[models.NotificationConfiguration](db),
	}
}

func (r *GormNotificationConfigurationRepository) Migrate() error {
	return r.db.AutoMigrate(&models.NotificationConfiguration{})
}
