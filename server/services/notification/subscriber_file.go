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

func (ns *NotificationSubscriber) HandleFileEvent(event string, change eventBus.Change[models.File], user models.User) {
	err := ns.SendFileNotifications(change, event, user)
	if err != nil {
		ns.logger.Error("Error sending file notifications", zap.Error(err))
	}
}

func (ns *NotificationSubscriber) GetFileGenericTemplate(event string, change eventBus.Change[models.File], user models.User) (string, template.HTML) {
	var body string

	if change.New.Task.Board.Name != "" {
		body += "<p><strong>Board:</strong> " + change.New.Task.Board.Name + "</p>"
	}
	if change.New.Task.Board.Name == "" && change.Old.Task.Board.Name != "" {
		body += "<p><strong>Board:</strong> " + change.Old.Task.Board.Name + "</p>"
	}
	body += "User <strong>@" + user.Username + "</strong> "

	switch event {
	case "file.created":
		subject := fmt.Sprintf("New File Uploaded: %s", change.New.Name)
		body += "has uploaded a new file to the task <strong>" + change.New.Task.Title + "</strong>.<br>"
		body += fmt.Sprintf(
			"<p><strong>File Name:</strong> %s</p>"+
				"<p><strong>File Type:</strong> %s</p>",
			change.New.Name, change.New.Type,
		)
		return subject, template.HTML(body)

	case "file.updated":
		subject := fmt.Sprintf("File Updated: %s", change.New.Name)
		body += fmt.Sprintf(
			"has updated the file (ID: %d) in the task <strong>%s</strong>.<br>"+
				"<p><strong>File Name:</strong> %s (old name: %s)</p>"+
				"<p><strong>File Type:</strong> %s (old type: %s)</p>",
			change.New.ID, change.New.Task.Title, change.New.Name, change.Old.Name, change.New.Type, change.Old.Type,
		)
		return subject, template.HTML(body)

	case "file.deleted":
		subject := fmt.Sprintf("File Deleted: %s", change.Old.Name)
		body += "has deleted the file from the task <strong>" + change.Old.Task.Title + "</strong>.<br><br>"
		body += fmt.Sprintf(
			"<p><strong>File Name:</strong> %s</p>"+
				"<p><strong>File Type:</strong> %s</p>",
			change.Old.Name, change.Old.Type,
		)
		return subject, template.HTML(body)

	default:
		subject := "File Notification"
		body += fmt.Sprintf("has triggered an unrecognised event (%s) for file (ID: %d) in the task <strong>%s</strong>.", event, change.Old.ID, change.Old.Task.Title)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherFileNotificationConfigurations(change eventBus.Change[models.File], event string) ([]models.NotificationConfiguration, error) {
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

func (ns *NotificationSubscriber) SendFileNotifications(change eventBus.Change[models.File], event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherFileNotificationConfigurations(change, event)
	if err != nil {
		ns.logger.Error("Error gathering file notification configurations", zap.Error(err))
		return err
	}
	if len(configs) == 0 {
		ns.logger.Info("No notification configurations found for file", zap.Uint("id", change.Old.ID), zap.Uint("board_id", change.Old.Task.BoardID))
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
		subject, body := ns.GetFileGenericTemplate(event, change, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"taskId":    taskID,
			"taskTitle": taskTitle,
			"appUrl":    ns.cfg.AppUrl,
		}
		err := ns.SendNotification(event, subject, tmplData, config)
		if err != nil {
			ns.logger.Error("Error sending notification", zap.Error(err))
			errRes = append(errRes, err)
		}
	}
	return errors.Join(errRes...)
}
