package notification

import (
	"errors"
	"fmt"
	"html/template"
	"server/database/repository"
	"server/models"
	"server/services/eventBus"

	"go.uber.org/zap"
)

func (ns *NotificationSubscriber) HandleCommentEvent(event string, change eventBus.Change[models.Comment], user models.User) {
	sendErr := ns.SendCommentNotifications(change, event, user)
	if sendErr != nil {
		ns.logger.Error("Error dispatching comment notifications", zap.Error(sendErr))
	}
}

func (ns *NotificationSubscriber) GetCommentGenericTemplate(event string, change eventBus.Change[models.Comment], user models.User) (string, template.HTML) {
	var body string

	if change.New.Task.Board.Name != "" {
		body += "<strong>Board:</strong> " + change.New.Task.Board.Name + "<br>"
	}
	if change.New.Task.Board.Name == "" && change.Old.Task.Board.Name != "" {
		body += "<strong>Board:</strong> " + change.Old.Task.Board.Name + "<br>"
	}

	body += "User <strong>@" + user.Username + "</strong> "

	switch event {
	case "comment.created":
		subject := fmt.Sprintf("New Comment on Task: %s", change.New.Task.Title)
		body += "has added a new comment to the task.<br>"
		body += fmt.Sprintf(
			"<p><strong>Comment:</strong></p><blockquote>%s</blockquote>"+
				"<p><strong>Task Title:</strong> %s</p>"+
				"<p><strong>Description:</strong> %s</p>"+
				"<p><strong>Status:</strong> %s</p>",
			change.New.Text, change.New.Task.Title, change.New.Task.Description, change.New.Task.Status,
		)
		return subject, template.HTML(body)

	case "comment.updated":
		subject := fmt.Sprintf("Comment Updated on Task: %s", change.New.Task.Title)
		body += fmt.Sprintf(
			"has updated the comment (ID: %d) on the task.<br>"+
				"<p><strong>Revised Comment:</strong></p><blockquote>%s</blockquote>"+
				"<p><strong>Old Comment:</strong></p><blockquote>%s</blockquote>",
			change.New.ID, change.New.Text, change.Old.Text,
		)
		return subject, template.HTML(body)

	case "comment.deleted":
		subject := fmt.Sprintf("Comment Deleted from Task: %s", change.Old.Task.Title)
		body += fmt.Sprintf(
			"has removed the comment (ID: %d) from the task.<br>"+
				"<p><strong>Removed Comment:</strong></p><blockquote>%s</blockquote>",
			change.Old.ID, change.Old.Text,
		)
		return subject, template.HTML(body)

	default:
		subject := "Comment Notification"
		body += fmt.Sprintf("has triggered an unrecognised event (%s) for comment (ID: %d).", event, change.Old.ID)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherCommentNotificationConfigurations(change eventBus.Change[models.Comment], event string) ([]models.NotificationConfiguration, error) {
	boardID := change.New.Task.BoardID
	if boardID == 0 {
		boardID = change.Old.Task.BoardID
	}
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", boardID, event),
	)
}

func (ns *NotificationSubscriber) SendCommentNotifications(change eventBus.Change[models.Comment], event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherCommentNotificationConfigurations(change, event)
	if err != nil {
		ns.logger.Error("Error gathering notification configurations", zap.Error(err))
		return err
	}
	if len(configs) == 0 {
		ns.logger.Info("No notification configurations found for comment", zap.Uint("id", change.New.ID), zap.Uint("board_id", change.New.Task.BoardID))
		return nil
	}
	for _, config := range configs {
		assigneeID := change.New.Task.AssigneeID
		if assigneeID == 0 {
			assigneeID = change.Old.Task.AssigneeID
		}
		if config.OnlyAssignee && config.UserID != assigneeID {
			continue
		}

		id := change.New.ID
		if id == 0 {
			id = change.Old.ID
		}
		taskID := change.New.Task.ID
		if taskID == 0 {
			taskID = change.Old.Task.ID
		}
		taskTitle := change.New.Task.Title
		if taskTitle == "" {
			taskTitle = change.Old.Task.Title
		}

		subject, body := ns.GetCommentGenericTemplate(event, change, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"commentId": id,
			"taskId":    taskID,
			"taskTitle": taskTitle,
			"appUrl":    ns.cfg.AppUrl,
		}
		err := ns.SendNotification(event, subject, tmplData, config)
		if err != nil {
			ns.logger.Error("Error dispatching notification", zap.Error(err))
			errRes = append(errRes, err)
		}
	}
	return errors.Join(errRes...)
}
