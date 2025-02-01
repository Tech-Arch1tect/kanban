package task

import (
	"errors"

	"server/database/repository"
	"server/models"
)

type MoveTaskRequest struct {
	TaskID     uint    `json:"task_id"`
	ColumnID   uint    `json:"column_id"`
	SwimlaneID uint    `json:"swimlane_id"`
	Position   float64 `json:"position"`
}

func (ts *TaskService) MoveTask(userID uint, request MoveTaskRequest) (models.Task, error) {
	task, err := ts.GetTask(userID, request.TaskID)
	if err != nil {
		return models.Task{}, err
	}

	column, err := ts.db.ColumnRepository.GetByID(request.ColumnID)
	if err != nil {
		return models.Task{}, err
	}
	swimlane, err := ts.db.SwimlaneRepository.GetByID(request.SwimlaneID)
	if err != nil {
		return models.Task{}, err
	}
	if column.BoardID != swimlane.BoardID || column.BoardID != task.BoardID {
		return models.Task{}, errors.New("column, swimlane and task must belong to the same board")
	}

	tasks, err := ts.db.TaskRepository.GetAll(
		repository.WithWhere("column_id = ? AND swimlane_id = ?", request.ColumnID, request.SwimlaneID),
		repository.WithOrder("position ASC"),
	)
	if err != nil {
		return models.Task{}, err
	}

	var filtered []models.Task
	for _, t := range tasks {
		if t.ID != task.ID {
			filtered = append(filtered, t)
		}
	}
	tasks = filtered

	var newPos float64
	if len(tasks) == 0 {
		newPos = request.Position
	} else {
		var nextPos float64
		foundNext := false
		for _, t := range tasks {
			if t.Position > request.Position {
				nextPos = t.Position
				foundNext = true
				break
			}
		}
		if !foundNext {
			newPos = request.Position + 1.0
		} else {
			newPos = (request.Position + nextPos) / 2.0
		}
	}

	task.Position = newPos
	task.ColumnID = request.ColumnID
	task.SwimlaneID = request.SwimlaneID

	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, err
	}

	return task, nil
}
