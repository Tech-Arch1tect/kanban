package notification

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"server/database/repository"
	"server/models"
)

func (ns *NotificationSubscriber) HandleExternalLinkEvent(event string, link models.TaskExternalLink, user models.User) {
	l, err := ns.TaskService.GetTaskExternalLink(link.ID)
	if err != nil {
		log.Println("Error getting link:", err)
		l = link
	}
	err = ns.SendExternalLinkNotifications(l, event, user)
	if err != nil {
		log.Println("Error sending link notifications:", err)
	}
}

func (ns *NotificationSubscriber) GetExternalLinkGenericTemplate(event string, link models.TaskExternalLink, user models.User) (string, template.HTML) {
	var body string
	appUrl := ns.cfg.AppUrl

	if link.Task.BoardID != 0 && link.Task.Board.Name != "" {
		boardDisplay := link.Task.Board.Name
		if link.Task.Board.ID != 0 {
			boardDisplay = fmt.Sprintf("<a href=\"%s/boards/%d\">%s</a>", appUrl, link.Task.Board.ID, link.Task.Board.Name)
		}
		body += fmt.Sprintf("<p><strong>Board:</strong> %s</p>", boardDisplay)
	}

	externalLink := fmt.Sprintf("<a href=\"%s\" target=\"_blank\">%s</a> (%s)", link.URL, link.Title, link.URL)

	switch event {
	case "externallink.created":
		subject := fmt.Sprintf("New External Link Added: %s", link.Title)
		body += fmt.Sprintf("<strong>@%s</strong> added a new external link: %s", user.Username, externalLink)
		return subject, template.HTML(body)

	case "externallink.updated":
		subject := fmt.Sprintf("External Link Updated: %s", link.Title)
		body += fmt.Sprintf("<strong>@%s</strong> updated the external link to: %s", user.Username, externalLink)
		return subject, template.HTML(body)

	case "externallink.deleted":
		subject := fmt.Sprintf("External Link Removed: %s", link.Title)
		body += fmt.Sprintf("<strong>@%s</strong> removed the external link: %s", user.Username, externalLink)
		return subject, template.HTML(body)

	default:
		subject := "External Link Notification"
		body += fmt.Sprintf(
			"<strong>@%s</strong> triggered an unrecognized event (<em>%s</em>) for external link (ID: %d) associated with task <strong>%s</strong>.",
			user.Username, event, link.ID, link.Task.Title,
		)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherExternalLinkNotificationConfigurations(link models.TaskExternalLink, event string) ([]models.NotificationConfiguration, error) {
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", link.Task.BoardID, event),
	)
}

func (ns *NotificationSubscriber) SendExternalLinkNotifications(link models.TaskExternalLink, event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherExternalLinkNotificationConfigurations(link, event)
	if err != nil {
		log.Println("Error gathering external link notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for external link:", link.ID, "in boards:", link.Task.BoardID)
		return nil
	}
	for _, config := range configs {
		if config.OnlyAssignee && config.UserID != link.Task.AssigneeID {
			continue
		}

		subject, body := ns.GetExternalLinkGenericTemplate(event, link, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"taskId":    link.Task.ID,
			"taskTitle": link.Task.Title,
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
