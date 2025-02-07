package notification

import (
	"server/config"
	"server/database/repository"
	"server/internal/email"
	"server/models"
	"server/services/comment"
	"server/services/eventBus"
)

type NotificationSubscriber struct {
	te             *eventBus.EventBus[models.Task]
	ce             *eventBus.EventBus[models.Comment]
	email          *email.EmailService
	db             *repository.Database
	cfg            *config.Config
	CommentService *comment.CommentService
}

func NewNotificationSubscriber(te *eventBus.EventBus[models.Task], ce *eventBus.EventBus[models.Comment], db *repository.Database, email *email.EmailService, cfg *config.Config, commentService *comment.CommentService) *NotificationSubscriber {
	return &NotificationSubscriber{
		te:             te,
		ce:             ce,
		email:          email,
		db:             db,
		cfg:            cfg,
		CommentService: commentService,
	}
}

func (ns *NotificationSubscriber) Subscribe() {
	ns.ce.SubscribeGlobal(func(event string, comment models.Comment, user models.User) {
		ns.HandleCommentEvent(event, comment, user)
	})
	ns.te.SubscribeGlobal(func(event string, task models.Task, user models.User) {
		ns.HandleTaskEvent(event, task, user)
	})
}
