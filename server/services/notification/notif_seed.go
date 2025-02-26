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
	"task.updated.due-date",
	"task.deleted",
	"task.moved",
}

var commentEvents = []string{
	"comment.created",
	"comment.updated",
	"comment.deleted",
}

var reactionEvents = []string{
	"reaction.created",
	"reaction.deleted",
}

var fileEvents = []string{
	"file.created",
	"file.updated",
	"file.deleted",
}

var linkEvents = []string{
	"link.created",
	"link.deleted",
}

var externalLinkEvents = []string{
	"externallink.created",
	"externallink.updated",
	"externallink.deleted",
}

var AllEvents = append(append(append(append(append(taskEvents, commentEvents...), fileEvents...), linkEvents...), externalLinkEvents...), reactionEvents...)

func (ns *NotificationService) SeedNotificationEvents() error {
	for _, event := range AllEvents {
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
