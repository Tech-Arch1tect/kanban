package notification

import (
	"errors"
	"fmt"
	"log"
	"server/database/repository"
	"server/models"
)

func (ns *NotificationSubscriber) HandleTaskEvent(event string, task models.Task) {
	sendErr := ns.SendTaskNotifications(task, event)
	if sendErr != nil {
		log.Println("Error sending task notifications:", sendErr)
	}
}

func (ns *NotificationSubscriber) GetTaskGenericTemplate(event string, task models.Task) (string, string) {
	switch event {
	case "task.created":
		subject := fmt.Sprintf("New Task Created: %s", task.Title)
		body := fmt.Sprintf("A new task has been created.\n\nTitle: %s\nDescription: %s\nStatus: %s", task.Title, task.Description, task.Status)
		return subject, body

	case "task.updated.title":
		subject := "Task Updated: Title Changed"
		body := fmt.Sprintf("The title for task (ID: %d) has been updated.\n\nNew Title: %s", task.ID, task.Title)
		return subject, body

	case "task.updated.description":
		subject := "Task Updated: Description Changed"
		body := fmt.Sprintf("The description for task (ID: %d) has been updated.\n\nTitle: %s\nNew Description: %s", task.ID, task.Title, task.Description)
		return subject, body

	case "task.updated.status":
		subject := "Task Updated: Status Changed"
		body := fmt.Sprintf("The status for task (ID: %d) has been updated.\n\nTitle: %s\nNew Status: %s", task.ID, task.Title, task.Status)
		return subject, body

	case "task.updated.assignee":
		subject := "Task Updated: Assignee Changed"
		body := fmt.Sprintf("The assignee for task (ID: %d) has been updated.\n\nTitle: %s\nNew Assignee: %s", task.ID, task.Title, task.Assignee.Username)
		return subject, body

	case "task.deleted":
		subject := fmt.Sprintf("Task Deleted: %s", task.Title)
		body := fmt.Sprintf("The following task has been deleted:\n\nTitle: %s\nDescription: %s", task.Title, task.Description)
		return subject, body

	case "task.moved":
		subject := fmt.Sprintf("Task Moved: %s", task.Title)
		body := fmt.Sprintf("The task has been moved to a new location.\n\nTitle: %s\nNew Board: %s\nNew Column: %s\nNew Swimlane: %s", task.Title, task.Board.Name, task.Column.Name, task.Swimlane.Name)
		return subject, body

	default:
		// Fallback for unknown events
		subject := "Task Notification"
		body := fmt.Sprintf("An event (%s) occurred for task (ID: %d).", event, task.ID)
		return subject, body
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

func (ns *NotificationSubscriber) SendTaskNotifications(task models.Task, event string) error {
	errRes := make([]error, 0)
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
		if config.OnlyAssignee {
			if config.UserID != task.AssigneeID {
				continue
			}
		}

		subject, body := ns.GetTaskGenericTemplate(event, task)
		err := ns.SendNotification(subject, body, config)
		if err != nil {
			// dont stop sending notifications just because of one failed one
			log.Println("Error sending notification:", err)
			errRes = append(errRes, err)
		}
	}
	return errors.Join(errRes...)
}
