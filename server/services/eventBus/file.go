package eventBus

import "server/models"

func NewFileEventBus() *EventBus[models.File] {
	return NewEventBus[models.File]()
}
