package task

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (ts *TaskService) UpdateTaskTitle(userID uint, taskID uint, title string) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	task.Title = title
	if err = task.Validate(); err != nil {
		return models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (ts *TaskService) UpdateTaskDescription(userID uint, taskID uint, description string) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	task.Description = description
	if err = task.Validate(); err != nil {
		return models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (ts *TaskService) UpdateTaskStatus(userID uint, taskID uint, status string) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	task.Status = status
	if err = task.Validate(); err != nil {
		return models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (ts *TaskService) UpdateTaskAssignee(userID uint, taskID uint, assigneeID uint) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	if assigneeID != 0 {
		can, _ = ts.rs.CheckRole(assigneeID, task.BoardID, role.MemberRole)
		if !can {
			return models.Task{}, errors.New("forbidden")
		}
	}

	task.AssigneeID = assigneeID
	if err = task.Validate(); err != nil {
		return models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, err
	}
	return task, nil
}
