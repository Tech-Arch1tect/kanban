package eventBus

import "server/models"

func NewTaskLinkEventBus() *EventBus[models.TaskLinks] {
	return NewEventBus[models.TaskLinks]()
}
