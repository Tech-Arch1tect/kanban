package notification

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (ns *NotificationService) CreateNotification(user *models.User, name, method, webhookURL, email string, events []string, boardIDs []uint, onlyAssignee bool) (*models.NotificationConfiguration, error) {
	if method == "webhook" && webhookURL == "" {
		return nil, errors.New("webhook URL is required for webhook method")
	}
	if method == "email" && email == "" {
		return nil, errors.New("email is required for email method")
	}

	var configEvents []models.NotificationEvent
	for _, eventName := range events {
		event, err := ns.db.NotificationEventRepository.GetFirst(repository.WithWhere("name = ?", eventName))
		if err != nil {
			return nil, err
		}
		configEvents = append(configEvents, event)
	}

	var boards []models.Board
	for _, id := range boardIDs {
		can, _ := ns.rs.CheckRole(user.ID, id, role.ReaderRole)
		if !can {
			return nil, errors.New("forbidden")
		}
		boards = append(boards, models.Board{Model: models.Model{ID: id}})
	}

	config := &models.NotificationConfiguration{
		UserID:       user.ID,
		Name:         name,
		Method:       method,
		WebhookURL:   webhookURL,
		Email:        email,
		Events:       configEvents,
		Boards:       boards,
		OnlyAssignee: onlyAssignee,
	}

	if err := ns.db.NotificationConfigurationRepository.Create(config); err != nil {
		return nil, err
	}

	return config, nil
}
