package notification

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"server/database/repository"
	"server/models"
)

func (ns *NotificationSubscriber) HandleTaskEvent(event string, task models.Task, user models.User) {
	err := ns.SendTaskNotifications(task, event, user)
	if err != nil {
		log.Println("Error sending task notifications:", err)
	}
}

func (ns *NotificationSubscriber) GetTaskGenericTemplate(event string, task models.Task, user models.User) (string, template.HTML) {
	var body string

	if task.Board.Name != "" {
		body += "<p><strong>Board:</strong> " + task.Board.Name + "</p>"
	}
	body += "User <strong>@" + user.Username + "</strong> "

	switch event {
	case "task.created":
		subject := fmt.Sprintf("New Task Created: %s", task.Title)
		body += "has created a new task with the following details:<br>"
		body += fmt.Sprintf(
			"<p><strong>Title:</strong> %s</p>"+
				"<p><strong>Description:</strong></p><blockquote>%s</blockquote>"+
				"<p><strong>Status:</strong> %s</p>",
			task.Title, task.Description, task.Status,
		)
		return subject, template.HTML(body)

	case "task.updated.title":
		subject := "Task Updated: Title Modified"
		body += fmt.Sprintf(
			"has updated the title of task (ID: %d).<br><p><strong>New Title:</strong> %s</p>",
			task.ID, task.Title,
		)
		return subject, template.HTML(body)

	case "task.updated.description":
		subject := "Task Updated: Description Modified"
		body += fmt.Sprintf(
			"has updated the description of task (ID: %d).<br><p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Description:</strong></p><blockquote>%s</blockquote>",
			task.ID, task.Title, task.Description,
		)
		return subject, template.HTML(body)

	case "task.updated.status":
		subject := "Task Updated: Status Modified"
		body += fmt.Sprintf(
			"has updated the status of task (ID: %d).<br><p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Status:</strong> %s</p>",
			task.ID, task.Title, task.Status,
		)
		return subject, template.HTML(body)

	case "task.updated.assignee":
		subject := "Task Updated: Assignee Modified"
		body += fmt.Sprintf(
			"has updated the assignee of task (ID: %d).<br><p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Assignee:</strong> %s</p>",
			task.ID, task.Title, task.Assignee.Username,
		)
		return subject, template.HTML(body)

	case "task.deleted":
		subject := fmt.Sprintf("Task Deleted: %s", task.Title)
		body += "has deleted the following task:<br><br>"
		body += fmt.Sprintf(
			"<p><strong>Title:</strong> %s</p>"+
				"<p><strong>Description:</strong></p><blockquote>%s</blockquote>",
			task.Title, task.Description,
		)
		return subject, template.HTML(body)

	case "task.moved":
		subject := fmt.Sprintf("Task Moved: %s", task.Title)
		body += fmt.Sprintf(
			"has moved task (ID: %d) to a new location.<br><br>"+
				"<p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Board:</strong> %s</p>"+
				"<p><strong>New Column:</strong> %s</p>"+
				"<p><strong>New Swimlane:</strong> %s</p>",
			task.ID, task.Title, task.Board.Name, task.Column.Name, task.Swimlane.Name,
		)
		return subject, template.HTML(body)

	default:
		subject := "Task Notification"
		body += fmt.Sprintf("has triggered an unrecognised event (%s) for task (ID: %d).", event, task.ID)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherTaskNotificationConfigurations(task models.Task, event string) ([]models.NotificationConfiguration, error) {
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", task.BoardID, event),
	)
}

func (ns *NotificationSubscriber) SendTaskNotifications(task models.Task, event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherTaskNotificationConfigurations(task, event)
	if err != nil {
		log.Println("Error gathering task notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for task:", task.ID, "in board:", task.BoardID)
		return nil
	}
	for _, config := range configs {
		if config.OnlyAssignee && config.UserID != task.AssigneeID {
			continue
		}

		subject, body := ns.GetTaskGenericTemplate(event, task, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"taskId":    task.ID,
			"taskTitle": task.Title,
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
