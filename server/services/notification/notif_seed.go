package notification

import (
	"errors"
	"server/database/repository"
	"server/models"

	"gorm.io/gorm"
)

var taskEvents = []string{
	"task.created",
	"task.updated.title",
	"task.updated.description",
	"task.updated.status",
	"task.updated.assignee",
	"task.deleted",
	"task.moved",
}

var AllEvents = taskEvents // TODO: add other events

func (ns *NotificationService) SeedNotificationEvents() error {
	for _, event := range taskEvents {
		_, err := ns.db.NotificationEventRepository.GetFirst(repository.WithWhere("name = ?", event))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := ns.db.NotificationEventRepository.Create(&models.NotificationEvent{Name: event}); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}
