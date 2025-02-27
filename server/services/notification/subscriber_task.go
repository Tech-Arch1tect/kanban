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

func (ns *NotificationSubscriber) HandleTaskEvent(event string, change eventBus.Change[models.Task], user models.User) {
	err := ns.SendTaskNotifications(change, event, user)
	if err != nil {
		log.Println("Error sending task notifications:", err)
	}
}

func (ns *NotificationSubscriber) GetTaskGenericTemplate(event string, change eventBus.Change[models.Task], user models.User) (string, template.HTML) {
	var body string

	if change.New.Board.Name != "" {
		body += "<p><strong>Board:</strong> " + change.New.Board.Name + "</p>"
	}
	if change.New.Board.Name == "" && change.Old.Board.Name != "" {
		body += "<p><strong>Board:</strong> " + change.Old.Board.Name + "</p>"
	}
	body += "User <strong>@" + user.Username + "</strong> "

	switch event {
	case "task.created":
		subject := fmt.Sprintf("New Task Created: %s", change.New.Title)
		body += "has created a new task with the following details:<br>"
		body += fmt.Sprintf(
			"<p><strong>Title:</strong> %s</p>"+
				"<p><strong>Description:</strong></p><blockquote>%s</blockquote>"+
				"<p><strong>Status:</strong> %s</p>",
			change.New.Title, change.New.Description, change.New.Status,
		)
		return subject, template.HTML(body)

	case "task.updated.title":
		subject := "Task Updated: Title Modified"
		body += fmt.Sprintf(
			"has updated the title of task (ID: %d).<br><p><strong>New Title:</strong> %s</p>",
			change.New.ID, change.New.Title,
		)
		return subject, template.HTML(body)

	case "task.updated.description":
		subject := "Task Updated: Description Modified"
		body += fmt.Sprintf(
			"has updated the description of task (ID: %d).<br><p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Description:</strong></p><blockquote>%s</blockquote>"+
				"<p><strong>Old Description:</strong></p><blockquote>%s</blockquote>",
			change.New.ID, change.New.Title, change.New.Description, change.Old.Description,
		)
		return subject, template.HTML(body)

	case "task.updated.status":
		subject := "Task Updated: Status Modified"
		body += fmt.Sprintf(
			"has updated the status of task (ID: %d).<br><p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Status:</strong> %s</p>"+
				"<p><strong>Old Status:</strong> %s</p>",
			change.New.ID, change.New.Title, change.New.Status, change.Old.Status,
		)
		return subject, template.HTML(body)

	case "task.updated.assignee":
		subject := "Task Updated: Assignee Modified"
		body += fmt.Sprintf(
			"has updated the assignee of task (ID: %d).<br><p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Assignee:</strong> %s</p>"+
				"<p><strong>Old Assignee:</strong> %s</p>",
			change.New.ID, change.New.Title, change.New.Assignee.Username, change.Old.Assignee.Username,
		)
		return subject, template.HTML(body)

	case "task.updated.due-date":
		subject := "Task Updated: Due Date Modified"
		body += fmt.Sprintf(
			"has updated the due date of task (ID: %d).<br><p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Due Date:</strong> %s</p>"+
				"<p><strong>Old Due Date:</strong> %s</p>",
			change.New.ID, change.New.Title, change.New.DueDate, change.Old.DueDate,
		)
		return subject, template.HTML(body)

	case "task.deleted":
		subject := fmt.Sprintf("Task Deleted: %s", change.Old.Title)
		body += "has deleted the following task:<br><br>"
		body += fmt.Sprintf(
			"<p><strong>Title:</strong> %s</p>"+
				"<p><strong>Description:</strong></p><blockquote>%s</blockquote>",
			change.Old.Title, change.Old.Description,
		)
		return subject, template.HTML(body)

	case "task.moved":
		subject := fmt.Sprintf("Task Moved: %s", change.New.Title)
		body += fmt.Sprintf(
			"has moved task (ID: %d) to a new location.<br><br>"+
				"<p><strong>Title:</strong> %s</p>"+
				"<p><strong>New Board:</strong> %s (old board: %s)</p>"+
				"<p><strong>New Column:</strong> %s (old column: %s)</p>"+
				"<p><strong>New Swimlane:</strong> %s (old swimlane: %s)</p>",
			change.New.ID, change.New.Title, change.New.Board.Name, change.Old.Board.Name, change.New.Column.Name, change.Old.Column.Name, change.New.Swimlane.Name, change.Old.Swimlane.Name,
		)
		return subject, template.HTML(body)

	default:
		subject := "Task Notification"
		body += fmt.Sprintf("has triggered an unrecognised event (%s) for task (ID: %d).", event, change.New.ID)
		return subject, template.HTML(body)
	}
}

func (ns *NotificationSubscriber) GatherTaskNotificationConfigurations(change eventBus.Change[models.Task], event string) ([]models.NotificationConfiguration, error) {
	boardID := change.New.BoardID
	if boardID == 0 {
		boardID = change.Old.BoardID
	}
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", boardID, event),
	)
}

func (ns *NotificationSubscriber) SendTaskNotifications(change eventBus.Change[models.Task], event string, user models.User) error {
	var errRes []error
	configs, err := ns.GatherTaskNotificationConfigurations(change, event)
	if err != nil {
		log.Println("Error gathering task notification configurations:", err)
		return err
	}
	if len(configs) == 0 {
		log.Println("No notification configurations found for task:", change.New.ID, "in board:", change.New.BoardID)
		return nil
	}
	for _, config := range configs {
		assigneeID := change.New.AssigneeID
		if assigneeID == 0 {
			assigneeID = change.Old.AssigneeID
		}
		if config.OnlyAssignee && config.UserID != assigneeID {
			continue
		}
		taskID := change.New.ID
		if taskID == 0 {
			taskID = change.Old.ID
		}
		taskTitle := change.New.Title
		if taskTitle == "" {
			taskTitle = change.Old.Title
		}
		subject, body := ns.GetTaskGenericTemplate(event, change, user)
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
