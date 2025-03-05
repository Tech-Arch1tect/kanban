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

func (ns *NotificationSubscriber) HandleLinkEvent(event string, change eventBus.Change[models.TaskLinks], user models.User) {
	err := ns.SendLinkNotifications(change, event, user)
	if err != nil {
		ns.logger.Error("Error sending link notifications", zap.Error(err))
	}
}

func (ns *NotificationSubscriber) GetLinkGenericTemplate(event string, change eventBus.Change[models.TaskLinks], user models.User) (string, template.HTML) {
	appUrl := ns.cfg.AppUrl

	srcTaskLink := fmt.Sprintf("<a href=\"%s/task/%d\">%s</a>", appUrl, change.New.SrcTask.ID, change.New.SrcTask.Title)
	dstTaskLink := fmt.Sprintf("<a href=\"%s/task/%d\">%s</a>", appUrl, change.New.DstTask.ID, change.New.DstTask.Title)
	if change.New.SrcTask.ID == 0 {
		srcTaskLink = fmt.Sprintf("<a href=\"%s/task/%d\">%s</a>", appUrl, change.Old.SrcTask.ID, change.Old.SrcTask.Title)
	}
	if change.New.DstTask.ID == 0 {
		dstTaskLink = fmt.Sprintf("<a href=\"%s/task/%d\">%s</a>", appUrl, change.Old.DstTask.ID, change.Old.DstTask.Title)
	}

	srcBoard := change.New.SrcTask.Board.Name
	if srcBoard == "" {
		srcBoard = change.Old.SrcTask.Board.Name
	}

	dstBoard := change.New.DstTask.Board.Name
	if dstBoard == "" {
		dstBoard = change.Old.DstTask.Board.Name
	}

	linkType := change.New.LinkType
	if linkType == "" {
		linkType = change.Old.LinkType
	}

	var body string
	switch event {
	case "link.created":
		subject := fmt.Sprintf("New Link Created: %s", linkType)
		body = fmt.Sprintf(
			"<strong>@%s</strong> created a new <em>%s</em> link connecting the following tasks:<br>"+
				"<p><strong>Source Task:</strong> %s (Board: %s)</p>"+
				"<p><strong>Destination Task:</strong> %s (Board: %s)</p>",
			user.Username, linkType, srcTaskLink, srcBoard, dstTaskLink, dstBoard,
		)
		return subject, template.HTML(body)

	case "link.deleted":
		subject := fmt.Sprintf("Link Deleted: %s", linkType)
		body = fmt.Sprintf(
			"<strong>@%s</strong> removed the <em>%s</em> link between the following tasks:<br>"+
				"<p><strong>Source Task:</strong> %s (Board: %s)</p>"+
				"<p><strong>Destination Task:</strong> %s (Board: %s)</p>",
			user.Username, linkType, srcTaskLink, srcBoard, dstTaskLink, dstBoard,
		)
		return subject, template.HTML(body)

	default:
		subject := "Link Notification"
		body = fmt.Sprintf(
			"<strong>@%s</strong> triggered an unrecognized event (<em>%s</em>) for link (ID: %d) associated with task %s.",
			user.Username, event, change.New.ID, dstTaskLink,
		)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherLinkNotificationConfigurations(change eventBus.Change[models.TaskLinks], event string) ([]models.NotificationConfiguration, error) {
	boardID := change.New.DstTask.BoardID
	if boardID == 0 {
		boardID = change.Old.DstTask.BoardID
	}
	dstTaskConfigs, err := ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", boardID, event),
	)
	if err != nil {
		return nil, err
	}
	boardID = change.New.SrcTask.BoardID
	if boardID == 0 {
		boardID = change.Old.SrcTask.BoardID
	}
	srcTaskConfigs, err := ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", boardID, event),
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

func (ns *NotificationSubscriber) SendLinkNotifications(change eventBus.Change[models.TaskLinks], event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherLinkNotificationConfigurations(change, event)
	if err != nil {
		ns.logger.Error("Error gathering link notification configurations", zap.Error(err))
		return err
	}
	if len(configs) == 0 {
		ns.logger.Info("No notification configurations found for link", zap.Uint("id", change.New.ID), zap.Uint("src_board_id", change.New.SrcTask.BoardID), zap.Uint("dst_board_id", change.New.DstTask.BoardID))
		return nil
	}
	for _, config := range configs {
		dstAssigneeID := change.New.DstTask.AssigneeID
		if dstAssigneeID == 0 {
			dstAssigneeID = change.Old.DstTask.AssigneeID
		}
		srcAssigneeID := change.New.SrcTask.AssigneeID
		if srcAssigneeID == 0 {
			srcAssigneeID = change.Old.SrcTask.AssigneeID
		}
		if config.OnlyAssignee && config.UserID != dstAssigneeID && config.UserID != srcAssigneeID {
			continue
		}

		subject, body := ns.GetLinkGenericTemplate(event, change, user)
		tmplData := map[string]interface{}{
			"subject": subject,
			"body":    body,
			"appUrl":  ns.cfg.AppUrl,
		}
		err := ns.SendNotification(event, subject, tmplData, config)
		if err != nil {
			ns.logger.Error("Error sending notification", zap.Error(err))
			errRes = append(errRes, err)
		}
	}
	return errors.Join(errRes...)
}
