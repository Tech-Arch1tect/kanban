package notification

import (
	"server/database/repository"
	"server/internal/email"
	"server/models"
	"server/services/eventBus"
)

type NotificationSubscriber struct {
	te    *eventBus.EventBus[models.Task]
	email *email.EmailService
	db    *repository.Database
}

func NewNotificationSubscriber(te *eventBus.EventBus[models.Task], db *repository.Database, email *email.EmailService) *NotificationSubscriber {
	return &NotificationSubscriber{
		te:    te,
		email: email,
		db:    db,
	}
}

func (ns *NotificationSubscriber) Subscribe() {
	ns.te.SubscribeGlobal(func(event string, task models.Task, user models.User) {
		ns.HandleTaskEvent(event, task, user)
	})
}
