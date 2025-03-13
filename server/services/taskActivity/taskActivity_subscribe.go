package taskActivity

import (
	"server/models"
	"server/services/eventBus"
	"strconv"

	"go.uber.org/zap"
)

func (s *TaskActivityService) Subscribe() {
	s.te.SubscribeGlobal(func(event string, change eventBus.Change[models.Task], user models.User) {
		var err error
		switch event {
		// we ignore task.deleted as there is no task to log activity for!
		case "task.created":
			err = s.CreateTaskActivity(change.New.ID, user.ID, event, "", change.New.Title+"\n"+change.New.Description)
		case "task.updated.title":
			err = s.CreateTaskActivity(change.New.ID, user.ID, event, change.Old.Title, change.New.Title)
		case "task.updated.description":
			err = s.CreateTaskActivity(change.New.ID, user.ID, event, change.Old.Description, change.New.Description)
		case "task.updated.status":
			err = s.CreateTaskActivity(change.New.ID, user.ID, event, change.Old.Status, change.New.Status)
		case "task.updated.assignee":
			err = s.CreateTaskActivity(change.New.ID, user.ID, event, change.Old.Assignee.Username, change.New.Assignee.Username)
		case "task.updated.due-date":
			oldDate := ""
			if change.Old.DueDate != nil {
				oldDate = change.Old.DueDate.Format("2006-01-02 15:04:05")
			}
			newDate := ""
			if change.New.DueDate != nil {
				newDate = change.New.DueDate.Format("2006-01-02 15:04:05")
			}
			err = s.CreateTaskActivity(change.New.ID, user.ID, event, oldDate, newDate)
		case "task.moved":
			newData := "Board: " + change.New.Board.Name + "\n" + "Column: " + change.New.Column.Name + "\n" + "Position: " + strconv.FormatFloat(change.New.Position, 'f', -1, 64) + "\n" + "Column: " + change.New.Column.Name
			oldData := "Board: " + change.Old.Board.Name + "\n" + "Column: " + change.Old.Column.Name + "\n" + "Position: " + strconv.FormatFloat(change.Old.Position, 'f', -1, 64) + "\n" + "Column: " + change.Old.Column.Name
			err = s.CreateTaskActivity(change.New.ID, user.ID, event, oldData, newData)
		}
		if err != nil {
			s.logger.Error("error creating task activity", zap.Error(err))
		}
	})
	s.ce.SubscribeGlobal(func(event string, change eventBus.Change[models.Comment], user models.User) {
		var err error
		switch event {
		case "comment.created":
			err = s.CreateTaskActivity(change.New.TaskID, user.ID, event, "", change.New.Text)
		case "comment.updated":
			err = s.CreateTaskActivity(change.New.TaskID, user.ID, event, change.Old.Text, change.New.Text)
		case "comment.deleted":
			err = s.CreateTaskActivity(change.Old.TaskID, user.ID, event, change.Old.Text, "")
		}
		if err != nil {
			s.logger.Error("error creating task activity", zap.Error(err))
		}
	})
	s.cre.SubscribeGlobal(func(event string, change eventBus.Change[models.Reaction], user models.User) {
		var err error
		switch event {
		case "reaction.created":
			err = s.CreateTaskActivity(change.New.Comment.TaskID, user.ID, event, "", change.New.Reaction)
		case "reaction.deleted":
			err = s.CreateTaskActivity(change.Old.Comment.TaskID, user.ID, event, change.Old.Reaction, "")
		}
		if err != nil {
			s.logger.Error("error creating task activity", zap.Error(err))
		}
	})
	s.fe.SubscribeGlobal(func(event string, change eventBus.Change[models.File], user models.User) {
		var err error
		switch event {
		case "file.created":
			err = s.CreateTaskActivity(change.New.TaskID, user.ID, event, "", change.New.Name)
		case "file.updated":
			err = s.CreateTaskActivity(change.New.TaskID, user.ID, event, change.Old.Name, change.New.Name)
		case "file.deleted":
			err = s.CreateTaskActivity(change.Old.TaskID, user.ID, event, change.Old.Name, "")
		}
		if err != nil {
			s.logger.Error("error creating task activity", zap.Error(err))
		}
	})
	s.le.SubscribeGlobal(func(event string, change eventBus.Change[models.TaskLinks], user models.User) {
		var err error
		switch event {
		case "link.created":
			err = s.CreateTaskActivity(change.New.SrcTaskID, user.ID, event, "", change.New.LinkType)
			// todo inverse link
		case "link.deleted":
			err = s.CreateTaskActivity(change.Old.SrcTaskID, user.ID, event, change.Old.LinkType, "")
			// todo inverse link
		}
		if err != nil {
			s.logger.Error("error creating task activity", zap.Error(err))
		}
	})
	s.lee.SubscribeGlobal(func(event string, change eventBus.Change[models.TaskExternalLink], user models.User) {
		var err error
		switch event {
		case "externallink.created":
			err = s.CreateTaskActivity(change.New.TaskID, user.ID, event, "", change.New.Title+" -> "+change.New.URL)
		case "externallink.updated":
			err = s.CreateTaskActivity(change.New.TaskID, user.ID, event, change.Old.Title+" -> "+change.Old.URL, change.New.Title+" -> "+change.New.URL)
		case "externallink.deleted":
			err = s.CreateTaskActivity(change.Old.TaskID, user.ID, event, change.Old.Title+" -> "+change.Old.URL, "")
		}
		if err != nil {
			s.logger.Error("error creating task activity", zap.Error(err))
		}
	})
}
