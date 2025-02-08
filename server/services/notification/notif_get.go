package notification

import (
	"server/database/repository"
	"server/models"
)

func (ns *NotificationService) GetNotificationConfigurations(user *models.User) ([]models.NotificationConfiguration, error) {
	return ns.db.NotificationConfigurationRepository.GetAll(repository.WithWhere("user_id = ?", user.ID), repository.WithPreload("User"), repository.WithPreload("Events"), repository.WithPreload("Boards"))
}

func (ns *NotificationService) GetNotificationConfiguration(id uint) (*models.NotificationConfiguration, error) {
	notif, err := ns.db.NotificationConfigurationRepository.GetFirst(repository.WithWhere("id = ?", id), repository.WithPreload("User"), repository.WithPreload("Events"), repository.WithPreload("Boards"))
	if err != nil {
		return nil, err
	}
	return &notif, nil
}
