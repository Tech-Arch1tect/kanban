package notification

import (
	"server/database/repository"
	"server/models"
)

func (ns *NotificationService) GetNotificationConfigurations(user *models.User) ([]models.NotificationConfiguration, error) {
	return ns.db.NotificationConfigurationRepository.GetAll(repository.WithWhere("user_id = ?", user.ID), repository.WithPreload("User"), repository.WithPreload("Events"), repository.WithPreload("Boards"))
}
