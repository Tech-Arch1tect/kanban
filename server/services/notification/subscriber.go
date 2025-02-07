package notification

import (
	"server/config"
	"server/database/repository"
	"server/internal/email"
	"server/models"
	"server/services/eventBus"
)

type NotificationSubscriber struct {
	te    *eventBus.EventBus[models.Task]
	email *email.EmailService
	db    *repository.Database
	cfg   *config.Config
}

func NewNotificationSubscriber(te *eventBus.EventBus[models.Task], db *repository.Database, email *email.EmailService, cfg *config.Config) *NotificationSubscriber {
	return &NotificationSubscriber{
		te:    te,
		email: email,
		db:    db,
		cfg:   cfg,
	}
}

func (ns *NotificationSubscriber) Subscribe() {
	ns.te.SubscribeGlobal(func(event string, task models.Task, user models.User) {
		ns.HandleTaskEvent(event, task, user)
	})
}
