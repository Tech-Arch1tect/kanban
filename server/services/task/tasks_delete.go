package task

import (
	"errors"
	"server/models"
	"server/services/role"

	"gorm.io/gorm"
)

func (ts *TaskService) DeleteTask(userID, taskID uint) (models.Task, error) {
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

	if err = ts.db.TaskRepository.Delete(task.ID); err != nil {
		return models.Task{}, err
	}

	return task, nil
}
