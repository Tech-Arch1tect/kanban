package notification

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"server/database/repository"
	"server/models"
)

func (ns *NotificationSubscriber) HandleCommentReactionEvent(event string, reaction models.Reaction, user models.User) {
	sendErr := ns.SendCommentReactionNotifications(reaction, event, user)
	if sendErr != nil {
		log.Println("Error dispatching comment reaction notifications:", sendErr)
	}
}

func (ns *NotificationSubscriber) GetCommentReactionGenericTemplate(event string, reaction models.Reaction, user models.User) (string, template.HTML) {
	var body string

	if reaction.Comment.Task.Board.Name != "" {
		body += "<strong>Board:</strong> " + reaction.Comment.Task.Board.Name + "<br>"
	}
	body += "User <strong>@" + user.Username + "</strong> "

	switch event {
	case "reaction.created":
		subject := "New Reaction on Comment"
		body += "has added a new reaction to:<br>"
		body += fmt.Sprintf(
			"<p><strong>Comment:</strong></p><blockquote>%s</blockquote>"+
				"<p><strong>Comment Author:</strong> %s</p>"+
				"<p><strong>Task Title:</strong> %s</p>"+
				"<p><strong>Status:</strong> %s</p>"+
				"<p><strong>Reaction:</strong> %s</p>",
			reaction.Comment.Text, reaction.Comment.User.Username, reaction.Comment.Task.Title, reaction.Comment.Task.Status, reaction.Reaction,
		)
		return subject, template.HTML(body)

	case "reaction.deleted":
		subject := "Reaction Deleted on Comment"
		body += "has removed a reaction from:<br>"
		body += fmt.Sprintf(
			"<p><strong>Comment:</strong></p><blockquote>%s</blockquote>"+
				"<p><strong>Comment Author:</strong> %s</p>"+
				"<p><strong>Task Title:</strong> %s</p>"+
				"<p><strong>Status:</strong> %s</p>"+
				"<p><strong>Reaction:</strong> %s</p>",
			reaction.Comment.Text, reaction.Comment.User.Username, reaction.Comment.Task.Title, reaction.Comment.Task.Status, reaction.Reaction,
		)
		return subject, template.HTML(body)

	default:
		subject := "Reaction Notification"
		body += fmt.Sprintf("has triggered an unrecognised event (%s) for reaction (ID: %d).", event, reaction.ID)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherCommentReactionNotificationConfigurations(reaction models.Reaction, event string) ([]models.NotificationConfiguration, error) {
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", reaction.Comment.Task.BoardID, event),
	)
}

func (ns *NotificationSubscriber) SendCommentReactionNotifications(reaction models.Reaction, event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherCommentReactionNotificationConfigurations(reaction, event)
	if err != nil {
		log.Println("Error gathering notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for reaction:", reaction.ID, "in board:", reaction.Comment.Task.BoardID)
		return nil
	}
	for _, config := range configs {
		if config.OnlyAssignee && config.UserID != reaction.Comment.Task.AssigneeID {
			continue
		}

		subject, body := ns.GetCommentReactionGenericTemplate(event, reaction, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"commentId": reaction.Comment.ID,
			"taskId":    reaction.Comment.Task.ID,
			"taskTitle": reaction.Comment.Task.Title,
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
