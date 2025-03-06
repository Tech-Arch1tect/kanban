package task

import (
	"errors"
	"server/models"
	"server/services/role"
	"time"
)

func (ts *TaskService) UpdateTaskTitle(userID uint, taskID uint, title string) (models.Task, models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, models.Task{}, errors.New("forbidden")
	}
	oldTask := task

	task.Title = title
	if err = task.Validate(); err != nil {
		return models.Task{}, models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, models.Task{}, err
	}

	task, err = ts.GetTask(userID, taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}
	return task, oldTask, nil
}

func (ts *TaskService) UpdateTaskDescription(userID uint, taskID uint, description string) (models.Task, models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, models.Task{}, errors.New("forbidden")
	}
	oldTask := task
	task.Description = description
	if err = task.Validate(); err != nil {
		return models.Task{}, models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, models.Task{}, err
	}

	task, err = ts.GetTask(userID, taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}
	return task, oldTask, nil
}

func (ts *TaskService) UpdateTaskStatus(userID uint, taskID uint, status string) (models.Task, models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, models.Task{}, errors.New("forbidden")
	}
	oldTask := task
	task.Status = status
	if err = task.Validate(); err != nil {
		return models.Task{}, models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, models.Task{}, err
	}

	task, err = ts.GetTask(userID, taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}
	return task, oldTask, nil
}

func (ts *TaskService) UpdateTaskAssignee(userID uint, taskID uint, assigneeID uint) (models.Task, models.Task, error) {
	task, err := ts.GetTask(userID, taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, models.Task{}, errors.New("forbidden")
	}
	oldTask := task
	if assigneeID != 0 {
		can, _ = ts.rs.CheckRole(assigneeID, task.BoardID, role.MemberRole)
		if !can {
			return models.Task{}, models.Task{}, errors.New("forbidden")
		}
	}

	task.AssigneeID = assigneeID
	if err = task.Validate(); err != nil {
		return models.Task{}, models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, models.Task{}, err
	}

	task, err = ts.GetTask(userID, taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}
	return task, oldTask, nil
}

func (ts *TaskService) UpdateTaskDueDate(userID uint, taskID uint, dueDate *time.Time) (models.Task, models.Task, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, models.Task{}, errors.New("forbidden")
	}
	oldTask := task
	task.DueDate = dueDate
	if err = task.Validate(); err != nil {
		return models.Task{}, models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, models.Task{}, err
	}

	task, err = ts.GetTask(userID, taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}
	return task, oldTask, nil
}

func (ts *TaskService) UpdateTaskColour(userID uint, taskID uint, colour string) (models.Task, models.Task, error) {
	task, err := ts.GetTask(userID, taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, models.Task{}, errors.New("forbidden")
	}
	oldTask := task
	task.Colour = colour
	if err = task.Validate(); err != nil {
		return models.Task{}, models.Task{}, err
	}
	if err = ts.db.TaskRepository.Update(&task); err != nil {
		return models.Task{}, models.Task{}, err
	}

	task, err = ts.GetTask(userID, taskID)
	if err != nil {
		return models.Task{}, models.Task{}, err
	}
	return task, oldTask, nil
}
