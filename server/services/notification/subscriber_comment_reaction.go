package notification

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"server/database/repository"
	"server/models"
	"server/services/eventBus"
)

func (ns *NotificationSubscriber) HandleCommentReactionEvent(event string, change eventBus.Change[models.Reaction], user models.User) {
	sendErr := ns.SendCommentReactionNotifications(change, event, user)
	if sendErr != nil {
		log.Println("Error dispatching comment reaction notifications:", sendErr)
	}
}

func (ns *NotificationSubscriber) GetCommentReactionGenericTemplate(event string, change eventBus.Change[models.Reaction], user models.User) (string, template.HTML) {
	var body string

	if change.New.Comment.Task.Board.Name != "" {
		body += "<strong>Board:</strong> " + change.New.Comment.Task.Board.Name + "<br>"
	}
	if change.New.Comment.Task.Board.Name == "" {
		body += "<strong>Board:</strong> " + change.Old.Comment.Task.Board.Name + "<br>"
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
			change.New.Comment.Text, change.New.Comment.User.Username, change.New.Comment.Task.Title, change.New.Comment.Task.Status, change.New.Reaction,
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
			change.Old.Comment.Text, change.Old.Comment.User.Username, change.Old.Comment.Task.Title, change.Old.Comment.Task.Status, change.Old.Reaction,
		)
		return subject, template.HTML(body)

	default:
		subject := "Reaction Notification"
		body += fmt.Sprintf("has triggered an unrecognised event (%s) for reaction (ID: %d).", event, change.New.ID)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherCommentReactionNotificationConfigurations(change eventBus.Change[models.Reaction], event string) ([]models.NotificationConfiguration, error) {
	boardID := change.New.Comment.Task.BoardID
	if boardID == 0 {
		boardID = change.Old.Comment.Task.BoardID
	}
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", boardID, event),
	)
}

func (ns *NotificationSubscriber) SendCommentReactionNotifications(change eventBus.Change[models.Reaction], event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherCommentReactionNotificationConfigurations(change, event)
	if err != nil {
		log.Println("Error gathering notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for reaction:", change.New.ID, "in board:", change.New.Comment.Task.BoardID)
		return nil
	}
	for _, config := range configs {
		assigneeID := change.New.Comment.Task.AssigneeID
		if assigneeID == 0 {
			assigneeID = change.Old.Comment.Task.AssigneeID
		}
		if config.OnlyAssignee && config.UserID != assigneeID {
			continue
		}

		commentID := change.New.Comment.ID
		if commentID == 0 {
			commentID = change.Old.Comment.ID
		}

		taskID := change.New.Comment.Task.ID
		if taskID == 0 {
			taskID = change.Old.Comment.Task.ID
		}

		taskTitle := change.New.Comment.Task.Title
		if taskTitle == "" {
			taskTitle = change.Old.Comment.Task.Title
		}

		subject, body := ns.GetCommentReactionGenericTemplate(event, change, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"commentId": commentID,
			"taskId":    taskID,
			"taskTitle": taskTitle,
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
