package notification

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"server/database/repository"
	"server/models"
)

func (ns *NotificationSubscriber) HandleLinkEvent(event string, link models.TaskLinks, user models.User) {
	l, err := ns.TaskService.GetTaskLink(link.ID)
	if err != nil {
		log.Println("Error getting link:", err)
		l = link
	}
	err = ns.SendLinkNotifications(l, event, user)
	if err != nil {
		log.Println("Error sending link notifications:", err)
	}
}

func (ns *NotificationSubscriber) GetLinkGenericTemplate(event string, link models.TaskLinks, user models.User) (string, template.HTML) {
	appUrl := ns.cfg.AppUrl

	srcTaskLink := fmt.Sprintf("<a href=\"%s/task/%d\">%s</a>", appUrl, link.SrcTask.ID, link.SrcTask.Title)
	dstTaskLink := fmt.Sprintf("<a href=\"%s/task/%d\">%s</a>", appUrl, link.DstTask.ID, link.DstTask.Title)

	srcBoardLink := link.SrcTask.Board.Name
	log.Println(link.SrcTask.Board)
	log.Println(link.DstTask.Board)
	if link.SrcTask.Board.ID != 0 {
		srcBoardLink = fmt.Sprintf("<a href=\"%s/boards/%s\">%s</a>", appUrl, link.SrcTask.Board.Slug, link.SrcTask.Board.Name)
	}

	dstBoardLink := link.DstTask.Board.Name
	if link.DstTask.Board.ID != 0 {
		dstBoardLink = fmt.Sprintf("<a href=\"%s/boards/%s\">%s</a>", appUrl, link.DstTask.Board.Slug, link.DstTask.Board.Name)
	}

	var body string
	switch event {
	case "link.created":
		subject := fmt.Sprintf("New Link Created: %s", link.LinkType)
		body = fmt.Sprintf(
			"<strong>@%s</strong> created a new <em>%s</em> link connecting the following tasks:<br>"+
				"<p><strong>Source Task:</strong> %s (Board: %s)</p>"+
				"<p><strong>Destination Task:</strong> %s (Board: %s)</p>",
			user.Username, link.LinkType, srcTaskLink, srcBoardLink, dstTaskLink, dstBoardLink,
		)
		return subject, template.HTML(body)

	case "link.deleted":
		subject := fmt.Sprintf("Link Deleted: %s", link.LinkType)
		body = fmt.Sprintf(
			"<strong>@%s</strong> removed the <em>%s</em> link between the following tasks:<br>"+
				"<p><strong>Source Task:</strong> %s (Board: %s)</p>"+
				"<p><strong>Destination Task:</strong> %s (Board: %s)</p>",
			user.Username, link.LinkType, srcTaskLink, srcBoardLink, dstTaskLink, dstBoardLink,
		)
		return subject, template.HTML(body)

	default:
		subject := "Link Notification"
		body = fmt.Sprintf(
			"<strong>@%s</strong> triggered an unrecognized event (<em>%s</em>) for link (ID: %d) associated with task %s.",
			user.Username, event, link.ID, dstTaskLink,
		)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherLinkNotificationConfigurations(link models.TaskLinks, event string) ([]models.NotificationConfiguration, error) {
	dstTaskConfigs, err := ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", link.DstTask.BoardID, event),
	)
	if err != nil {
		return nil, err
	}
	srcTaskConfigs, err := ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", link.SrcTask.BoardID, event),
	)
	if err != nil {
		return nil, err
	}

	uniqueConfigs := make(map[uint]models.NotificationConfiguration)

	for _, config := range dstTaskConfigs {
		uniqueConfigs[config.ID] = config
	}

	for _, config := range srcTaskConfigs {
		uniqueConfigs[config.ID] = config
	}

	var result []models.NotificationConfiguration
	for _, config := range uniqueConfigs {
		result = append(result, config)
	}

	return result, nil
}

func (ns *NotificationSubscriber) SendLinkNotifications(link models.TaskLinks, event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherLinkNotificationConfigurations(link, event)
	if err != nil {
		log.Println("Error gathering link notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for link:", link.ID, "in boards:", link.SrcTask.BoardID, link.DstTask.BoardID)
		return nil
	}
	for _, config := range configs {
		if config.OnlyAssignee && config.UserID != link.DstTask.AssigneeID && config.UserID != link.SrcTask.AssigneeID {
			continue
		}

		subject, body := ns.GetLinkGenericTemplate(event, link, user)
		tmplData := map[string]interface{}{
			"subject": subject,
			"body":    body,
			"appUrl":  ns.cfg.AppUrl,
		}
		err := ns.SendNotification(event, subject, tmplData, config)
		if err != nil {
			log.Println("Error sending notification:", err)
			errRes = append(errRes, err)
		}
	}
	return errors.Join(errRes...)
}
