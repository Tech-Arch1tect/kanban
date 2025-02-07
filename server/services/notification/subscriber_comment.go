package notification

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"server/database/repository"
	"server/models"
)

func (ns *NotificationSubscriber) HandleCommentEvent(event string, comment models.Comment, user models.User) {
	c, err := ns.CommentService.GetComment(comment.ID)
	if err != nil {
		log.Println("Error retrieving comment:", err)
		return
	}
	sendErr := ns.SendCommentNotifications(c, event, user)
	if sendErr != nil {
		log.Println("Error dispatching comment notifications:", sendErr)
	}
}

func (ns *NotificationSubscriber) GetCommentGenericTemplate(event string, comment models.Comment, user models.User) (string, template.HTML) {
	var body string

	if comment.Task.Board.Name != "" {
		body += "<strong>Board:</strong> " + comment.Task.Board.Name + "<br>"
	}
	body += "User <strong>@" + user.Username + "</strong> "

	switch event {
	case "comment.created":
		subject := fmt.Sprintf("New Comment on Task: %s", comment.Task.Title)
		body += "has added a new comment to the task.<br>"
		body += fmt.Sprintf(
			"<p><strong>Comment:</strong></p><blockquote>%s</blockquote>"+
				"<p><strong>Task Title:</strong> %s</p>"+
				"<p><strong>Description:</strong> %s</p>"+
				"<p><strong>Status:</strong> %s</p>",
			comment.Text, comment.Task.Title, comment.Task.Description, comment.Task.Status,
		)
		return subject, template.HTML(body)

	case "comment.updated":
		subject := fmt.Sprintf("Comment Updated on Task: %s", comment.Task.Title)
		body += fmt.Sprintf(
			"has updated the comment (ID: %d) on the task.<br>"+
				"<p><strong>Revised Comment:</strong></p><blockquote>%s</blockquote>",
			comment.ID, comment.Text,
		)
		return subject, template.HTML(body)

	case "comment.deleted":
		subject := fmt.Sprintf("Comment Deleted from Task: %s", comment.Task.Title)
		body += fmt.Sprintf(
			"has removed the comment (ID: %d) from the task.<br>"+
				"<p><strong>Removed Comment:</strong></p><blockquote>%s</blockquote>",
			comment.ID, comment.Text,
		)
		return subject, template.HTML(body)

	default:
		subject := "Comment Notification"
		body += fmt.Sprintf("has triggered an unrecognised event (%s) for comment (ID: %d).", event, comment.ID)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherCommentNotificationConfigurations(comment models.Comment, event string) ([]models.NotificationConfiguration, error) {
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", comment.Task.BoardID, event),
	)
}

func (ns *NotificationSubscriber) SendCommentNotifications(comment models.Comment, event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherCommentNotificationConfigurations(comment, event)
	if err != nil {
		log.Println("Error gathering notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for comment:", comment.ID, "in board:", comment.Task.BoardID)
		return nil
	}
	for _, config := range configs {
		if config.OnlyAssignee && config.UserID != comment.Task.AssigneeID {
			continue
		}

		subject, body := ns.GetCommentGenericTemplate(event, comment, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"commentId": comment.ID,
			"taskId":    comment.Task.ID,
			"taskTitle": comment.Task.Title,
			"appUrl":    ns.cfg.AppUrl,
		}
		err := ns.SendNotification(event, subject, tmplData, config)
		if err != nil {
			log.Println("Error dispatching notification:", err)
			errRes = append(errRes, err)
		}
	}
	return errors.Join(errRes...)
}
