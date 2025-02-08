package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type NotificationConfigurationRepository interface {
	Repository[models.NotificationConfiguration]
	UpdateAndReplaceEventsAndBoards(notification *models.NotificationConfiguration) error
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

func (r *GormNotificationConfigurationRepository) UpdateAndReplaceEventsAndBoards(notification *models.NotificationConfiguration) error {
	if err := r.Update(notification); err != nil {
		return err
	}

	if err := r.db.Model(notification).Association("Events").Replace(notification.Events); err != nil {
		return err
	}

	if err := r.db.Model(notification).Association("Boards").Replace(notification.Boards); err != nil {
		return err
	}

	return nil
}
