package eventBus

import "server/models"

func NewCommentEventBus() *EventBus[models.Comment] {
	return NewEventBus[models.Comment]()
}
