package task

import (
	"errors"
	"server/models"
	"server/services/role"

	"gorm.io/gorm"
)

func (ts *TaskService) DeleteTaskRequest(userID, taskID uint) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	err = ts.DeleteTask(task.ID)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) DeleteTask(taskID uint) error {
	return ts.db.TaskRepository.Delete(taskID)
}
