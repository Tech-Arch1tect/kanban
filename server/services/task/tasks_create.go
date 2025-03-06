package task

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
	"time"

	"gorm.io/gorm"
)

func (ts *TaskService) CreateTask(userID uint, request CreateTaskRequest) (models.Task, error) {
	can, _ := ts.rs.CheckRole(userID, request.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	if _, err := ts.db.SwimlaneRepository.GetByID(request.SwimlaneID); err != nil {
		return models.Task{}, err
	}
	if _, err := ts.db.ColumnRepository.GetByID(request.ColumnID); err != nil {
		return models.Task{}, err
	}

	taskPosition, err := ts.db.TaskRepository.GetFirst(
		repository.WithWhere("column_id = ? AND swimlane_id = ?", request.ColumnID, request.SwimlaneID),
		repository.WithOrder("position DESC"),
	)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Task{}, err
	}

	var assignee models.User
	if request.AssigneeID != 0 {
		assignee, err = ts.db.UserRepository.GetByID(request.AssigneeID)
		if err != nil {
			return models.Task{}, err
		}
	}

	dueDate := request.DueDate
	if dueDate != nil {
		t := time.Time(*dueDate)
		dueDate = &t
	}

	task := models.Task{
		ParentTaskID: request.ParentTaskID,
		BoardID:      request.BoardID,
		Title:        request.Title,
		Description:  request.Description,
		SwimlaneID:   request.SwimlaneID,
		Status:       request.Status,
		ColumnID:     request.ColumnID,
		Position:     taskPosition.Position + 1,
		CreatorID:    userID,
		AssigneeID:   assignee.ID,
		DueDate:      dueDate,
		Colour:       request.Colour,
	}

	if err = task.Validate(); err != nil {
		return models.Task{}, err
	}
	if err = ts.db.TaskRepository.Create(&task); err != nil {
		return models.Task{}, err
	}

	return task, nil
}
