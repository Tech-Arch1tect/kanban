package eventBus

import "server/models"

func NewTaskLinkEventBus() *EventBus[models.TaskLinks] {
	return NewEventBus[models.TaskLinks]()
}

func NewTaskExternalLinkEventBus() *EventBus[models.TaskExternalLink] {
	return NewEventBus[models.TaskExternalLink]()
}
