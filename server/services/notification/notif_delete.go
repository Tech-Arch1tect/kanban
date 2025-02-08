package notification

import (
	"errors"
	"server/database/repository"
	"server/models"
)

func (s *NotificationService) DeleteNotification(user *models.User, id uint) error {
	notification, err := s.db.NotificationConfigurationRepository.GetByID(id, repository.WithPreload("User"))
	if err != nil {
		return err
	}

	if notification.UserID != user.ID && user.Role != models.RoleAdmin {
		return errors.New("forbidden")
	}

	return s.db.NotificationConfigurationRepository.Delete(id)
}
