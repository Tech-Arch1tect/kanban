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

func (ns *NotificationSubscriber) HandleMentionEvent(event string, change eventBus.Change[eventBus.TaskOrComment], user models.User) {
	err := ns.SendMentionNotifications(change, user)
	if err != nil {
		ns.logger.Error("Error sending mention notifications", zap.Error(err))
	}
}

func (ns *NotificationSubscriber) GetBoardFromTaskOrComment(change eventBus.Change[eventBus.TaskOrComment]) (models.Board, error) {
	if change.New.Task != nil && change.New.Task.BoardID != 0 {
		return ns.db.BoardRepository.GetFirst(repository.WithWhere("id = ?", change.New.Task.BoardID))
	}
	if change.New.Comment != nil && change.New.Comment.Task.BoardID != 0 {
		return ns.db.BoardRepository.GetFirst(repository.WithWhere("id = ?", change.New.Comment.Task.BoardID))
	}
	return models.Board{}, errors.New("no board found")
}

func (ns *NotificationSubscriber) GetMentionGenericTemplate(change eventBus.Change[eventBus.TaskOrComment], user models.User) (string, template.HTML) {
	var body string
	board, err := ns.GetBoardFromTaskOrComment(change)
	if err != nil {
		ns.logger.Error("Error getting board from task or comment", zap.Error(err))
	}
	if board.Name != "" {
		body += "<p><strong>Board:</strong> " + board.Name + "</p>"
	}
	body += "User <strong>@" + user.Username + "</strong> "

	t := "Comment"
	if change.New.Task != nil {
		t = "Task Description"
	}

	subject := fmt.Sprintf("You have been mentioned in a %s", t)

	if t == "Comment" {
		body += fmt.Sprintf(
			"<p><strong>Comment:</strong></p><blockquote>%s</blockquote>",
			change.New.Comment.Text,
		)
	}

	if t == "Task Description" {
		body += fmt.Sprintf(
			"<p><strong>Task Title:</strong> %s</p>"+
				"<p><strong>Description:</strong></p><blockquote>%s</blockquote>",
			change.New.Task.Title, change.New.Task.Description,
		)
	}

	return subject, template.HTML(body)
}

func (ns *NotificationSubscriber) GatherMentionNotificationConfigurations(change eventBus.Change[eventBus.TaskOrComment]) ([]models.NotificationConfiguration, error) {
	board, err := ns.GetBoardFromTaskOrComment(change)
	if err != nil {
		ns.logger.Error("Error getting board from task or comment", zap.Error(err))
	}
	return ns.db.NotificationConfigurationRepository.GetAll(
		repository.WithJoin("JOIN notification_configuration_boards ON notification_configuration_boards.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_configuration_events ON notification_configuration_events.notification_configuration_id = notification_configurations.id"),
		repository.WithJoin("JOIN notification_events ON notification_events.id = notification_configuration_events.notification_event_id"),
		repository.WithWhere("notification_configuration_boards.board_id = ? AND notification_events.name = ?", board.ID, "mentioned"),
	)
}

func (ns *NotificationSubscriber) SendMentionNotifications(change eventBus.Change[eventBus.TaskOrComment], user models.User) error {
	var errRes []error
	configs, err := ns.GatherMentionNotificationConfigurations(change)
	if err != nil {
		ns.logger.Error("Error gathering mention notification configurations", zap.Error(err))
		return err
	}
	if len(configs) == 0 {
		return nil
	}
	for _, config := range configs {
		if config.UserID != change.New.MentionedUser.ID {
			continue
		}
		taskID := uint(0)
		if change.New.Task != nil {
			taskID = change.New.Task.ID
		} else if change.New.Comment != nil {
			taskID = change.New.Comment.Task.ID
		}
		taskTitle := ""
		if change.New.Task != nil {
			taskTitle = change.New.Task.Title
		} else {
			taskTitle = change.New.Comment.Task.Title
		}
		subject, body := ns.GetMentionGenericTemplate(change, user)
		tmplData := map[string]interface{}{
			"subject":   subject,
			"body":      body,
			"taskId":    taskID,
			"taskTitle": taskTitle,
			"appUrl":    ns.cfg.AppUrl,
		}
		err := ns.SendNotification("mentioned", subject, tmplData, config)
		if err != nil {
			ns.logger.Error("Error sending notification", zap.Error(err))
			errRes = append(errRes, err)
		}
	}
	return errors.Join(errRes...)
}
