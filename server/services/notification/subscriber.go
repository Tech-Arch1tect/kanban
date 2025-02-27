package notification

import (
	"server/config"
	"server/database/repository"
	"server/internal/email"
	"server/models"
	"server/services/comment"
	"server/services/eventBus"
	"server/services/task"
)

type NotificationSubscriber struct {
	te             *eventBus.EventBus[models.Task]
	ce             *eventBus.EventBus[models.Comment]
	fe             *eventBus.EventBus[models.File]
	le             *eventBus.EventBus[models.TaskLinks]
	lee            *eventBus.EventBus[models.TaskExternalLink]
	cre            *eventBus.EventBus[models.Reaction]
	email          *email.EmailService
	db             *repository.Database
	cfg            *config.Config
	CommentService *comment.CommentService
	TaskService    *task.TaskService
}

func NewNotificationSubscriber(te *eventBus.EventBus[models.Task], ce *eventBus.EventBus[models.Comment], fe *eventBus.EventBus[models.File], le *eventBus.EventBus[models.TaskLinks], lee *eventBus.EventBus[models.TaskExternalLink], cre *eventBus.EventBus[models.Reaction], db *repository.Database, email *email.EmailService, cfg *config.Config, commentService *comment.CommentService, taskService *task.TaskService) *NotificationSubscriber {
	return &NotificationSubscriber{
		te:             te,
		ce:             ce,
		fe:             fe,
		le:             le,
		lee:            lee,
		cre:            cre,
		email:          email,
		db:             db,
		cfg:            cfg,
		CommentService: commentService,
		TaskService:    taskService,
	}
}

func (ns *NotificationSubscriber) Subscribe() {
	ns.ce.SubscribeGlobal(func(event string, change eventBus.Change[models.Comment], user models.User) {
		ns.HandleCommentEvent(event, change, user)
	})
	ns.te.SubscribeGlobal(func(event string, change eventBus.Change[models.Task], user models.User) {
		ns.HandleTaskEvent(event, change, user)
	})
	ns.fe.SubscribeGlobal(func(event string, change eventBus.Change[models.File], user models.User) {
		ns.HandleFileEvent(event, change, user)
	})
	ns.le.SubscribeGlobal(func(event string, change eventBus.Change[models.TaskLinks], user models.User) {
		ns.HandleLinkEvent(event, change, user)
	})
	ns.lee.SubscribeGlobal(func(event string, change eventBus.Change[models.TaskExternalLink], user models.User) {
		ns.HandleExternalLinkEvent(event, change, user)
	})
	ns.cre.SubscribeGlobal(func(event string, change eventBus.Change[models.Reaction], user models.User) {
		ns.HandleCommentReactionEvent(event, change, user)
	})
}
