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

func (ns *NotificationSubscriber) HandleExternalLinkEvent(event string, change eventBus.Change[models.TaskExternalLink], user models.User) {
	err := ns.SendExternalLinkNotifications(change, event, user)
	if err != nil {
		log.Println("Error sending link notifications:", err)
	}
}

func (ns *NotificationSubscriber) GetExternalLinkGenericTemplate(event string, change eventBus.Change[models.TaskExternalLink], user models.User) (string, template.HTML) {
	var body string
	appUrl := ns.cfg.AppUrl

	if change.New.Task.BoardID != 0 && change.New.Task.Board.Name != "" {
		boardDisplay := change.New.Task.Board.Name
		if change.New.Task.Board.ID != 0 {
			boardDisplay = fmt.Sprintf("<a href=\"%s/boards/%s\">%s</a>", appUrl, change.New.Task.Board.Slug, change.New.Task.Board.Name)
		}
		body += fmt.Sprintf("<p><strong>Board test:</strong> %s</p>", boardDisplay)
	}

	if change.New.Task.BoardID == 0 && change.Old.Task.BoardID != 0 && change.Old.Task.Board.Name != "" {
		boardDisplay := change.Old.Task.Board.Name
		if change.Old.Task.Board.ID != 0 {
			boardDisplay = fmt.Sprintf("<a href=\"%s/boards/%s\">%s</a>", appUrl, change.Old.Task.Board.Slug, change.Old.Task.Board.Name)
		}
		body += fmt.Sprintf("<p><strong>Board:</strong> %s</p>", boardDisplay)
	}

	externalLink := fmt.Sprintf("<a href=\"%s\" target=\"_blank\">%s</a> (%s)", change.New.URL, change.New.Title, change.New.URL)
	if change.New.URL == "" {
		externalLink = fmt.Sprintf("<a href=\"%s\" target=\"_blank\">%s</a> (%s)", change.Old.URL, change.Old.Title, change.Old.URL)
	}

	linkTitle := change.New.Title
	if linkTitle == "" {
		linkTitle = change.Old.Title
	}

	linkID := change.New.ID
	if linkID == 0 {
		linkID = change.Old.ID
	}

	switch event {
	case "externallink.created":
		subject := fmt.Sprintf("New External Link Added: %s", linkTitle)
		body += fmt.Sprintf("<strong>@%s</strong> added a new external link: %s", user.Username, externalLink)
		return subject, template.HTML(body)

	case "externallink.updated":
		subject := fmt.Sprintf("External Link Updated: %s", linkTitle)
		body += fmt.Sprintf("<strong>@%s</strong> updated the external link to: %s", user.Username, externalLink)
		return subject, template.HTML(body)

	case "externallink.deleted":
		subject := fmt.Sprintf("External Link Removed: %s", linkTitle)
		body += fmt.Sprintf("<strong>@%s</strong> removed the external link: %s", user.Username, externalLink)
		return subject, template.HTML(body)

	default:
		subject := "External Link Notification"
		body += fmt.Sprintf(
			"<strong>@%s</strong> triggered an unrecognized event (<em>%s</em>) for external link (ID: %d) associated with task <strong>%s</strong>.",
			user.Username, event, linkID, linkTitle,
		)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherExternalLinkNotificationConfigurations(change eventBus.Change[models.TaskExternalLink], event string) ([]models.NotificationConfiguration, error) {
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

func (ns *NotificationSubscriber) SendExternalLinkNotifications(change eventBus.Change[models.TaskExternalLink], event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherExternalLinkNotificationConfigurations(change, event)
	if err != nil {
		log.Println("Error gathering external link notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for external link:", change.New.ID, "in boards:", change.New.Task.BoardID)
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

		taskID := change.New.Task.ID
		if taskID == 0 {
			taskID = change.Old.Task.ID
		}

		taskTitle := change.New.Task.Title
		if taskTitle == "" {
			taskTitle = change.Old.Task.Title
		}

		subject, body := ns.GetExternalLinkGenericTemplate(event, change, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"taskId":    taskID,
			"taskTitle": taskTitle,
			"appUrl":    ns.cfg.AppUrl,
		}
		err := ns.SendNotification(event, subject, tmplData, config)
		if err != nil {
			log.Println("Error sending notification:", err)
			errRes = append(errRes, err)
		}
	}
	return errors.Join(errRes...)
}
