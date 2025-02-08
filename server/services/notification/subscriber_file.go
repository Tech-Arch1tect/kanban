package notification

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"server/database/repository"
	"server/models"
)

func (ns *NotificationSubscriber) HandleFileEvent(event string, file models.File, user models.User) {
	f, _, err := ns.TaskService.GetFile(file.UploadedBy, file.ID)
	if err != nil {
		log.Println("Error getting file:", err)
		f = file
	}
	err = ns.SendFileNotifications(f, event, user)
	if err != nil {
		log.Println("Error sending file notifications:", err)
	}
}

func (ns *NotificationSubscriber) GetFileGenericTemplate(event string, file models.File, user models.User) (string, template.HTML) {
	var body string

	if file.Task.Board.Name != "" {
		body += "<p><strong>Board:</strong> " + file.Task.Board.Name + "</p>"
	}
	body += "User <strong>@" + user.Username + "</strong> "

	switch event {
	case "file.created":
		subject := fmt.Sprintf("New File Uploaded: %s", file.Name)
		body += "has uploaded a new file to the task <strong>" + file.Task.Title + "</strong>.<br>"
		body += fmt.Sprintf(
			"<p><strong>File Name:</strong> %s</p>"+
				"<p><strong>File Type:</strong> %s</p>",
			file.Name, file.Type,
		)
		return subject, template.HTML(body)

	case "file.updated":
		subject := fmt.Sprintf("File Updated: %s", file.Name)
		body += fmt.Sprintf(
			"has updated the file (ID: %d) in the task <strong>%s</strong>.<br>"+
				"<p><strong>File Name:</strong> %s</p>"+
				"<p><strong>File Type:</strong> %s</p>",
			file.ID, file.Task.Title, file.Name, file.Type,
		)
		return subject, template.HTML(body)

	case "file.deleted":
		subject := fmt.Sprintf("File Deleted: %s", file.Name)
		body += "has deleted the file from the task <strong>" + file.Task.Title + "</strong>.<br><br>"
		body += fmt.Sprintf(
			"<p><strong>File Name:</strong> %s</p>"+
				"<p><strong>File Type:</strong> %s</p>",
			file.Name, file.Type,
		)
		return subject, template.HTML(body)

	default:
		subject := "File Notification"
		body += fmt.Sprintf("has triggered an unrecognised event (%s) for file (ID: %d) in the task <strong>%s</strong>.", event, file.ID, file.Task.Title)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherFileNotificationConfigurations(file models.File, event string) ([]models.NotificationConfiguration, error) {
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", file.Task.BoardID, event),
	)
}

func (ns *NotificationSubscriber) SendFileNotifications(file models.File, event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherFileNotificationConfigurations(file, event)
	if err != nil {
		log.Println("Error gathering file notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for file:", file.ID, "in board:", file.Task.BoardID)
		return nil
	}
	for _, config := range configs {
		if config.OnlyAssignee && config.UserID != file.Task.AssigneeID {
			continue
		}

		subject, body := ns.GetFileGenericTemplate(event, file, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"taskId":    file.Task.ID,
			"taskTitle": file.Task.Title,
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
