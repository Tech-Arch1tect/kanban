package notification

import (
	"server/models"
)

var taskEvents = []string{
	"task.created",
	"task.updated",
	"task.deleted",
	"task.moved",
}

func (ns *NotificationService) SeedNotificationEvents() error {
	for _, event := range taskEvents {
		ns.db.NotificationEventRepository.Create(&models.NotificationEvent{Name: event})
	}
	return nil
}
