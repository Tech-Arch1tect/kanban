package eventBus

import "server/models"

func NewTaskEventBus() *EventBus[models.Task] {
	return NewEventBus[models.Task]()
}

// used for mentions
type TaskOrComment struct {
	Task          *models.Task    // task description
	Comment       *models.Comment // comment text
	MentionedUser *models.User    // user mentioned
}

func NewTaskOrCommentEventBus() *EventBus[TaskOrComment] {
	return NewEventBus[TaskOrComment]()
}
