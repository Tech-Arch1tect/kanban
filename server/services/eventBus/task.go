package eventBus

import "server/models"

func NewTaskEventBus() *EventBus[models.Task] {
	return NewEventBus[models.Task]()
}
