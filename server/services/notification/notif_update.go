package notification

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (ns *NotificationService) UpdateNotification(user *models.User, id uint, name, method, webhookURL, email string, events []string, boardIDs []uint, onlyAssignee bool) (*models.NotificationConfiguration, error) {
	notif, err := ns.GetNotificationConfiguration(id)
	if err != nil {
		return nil, err
	}

	if notif.UserID != user.ID && user.Role != models.RoleAdmin {
		return nil, errors.New("forbidden")
	}

	notif.Name = name
	notif.Method = method
	notif.WebhookURL = webhookURL
	notif.Email = email
	notif.OnlyAssignee = onlyAssignee

	var boards []models.Board
	for _, id := range boardIDs {
		can, _ := ns.rs.CheckRole(user.ID, id, role.ReaderRole)
		if !can {
			return nil, errors.New("forbidden")
		}
		boards = append(boards, models.Board{Model: models.Model{ID: id}})
	}

	notif.Boards = boards

	var configEvents []models.NotificationEvent
	for _, eventName := range events {
		event, err := ns.db.NotificationEventRepository.GetFirst(repository.WithWhere("name = ?", eventName))
		if err != nil {
			return nil, err
		}
		configEvents = append(configEvents, event)
	}
	notif.Events = configEvents

	if err := ns.db.NotificationConfigurationRepository.UpdateAndReplaceEventsAndBoards(notif); err != nil {
		return nil, err
	}
	return notif, nil
}
