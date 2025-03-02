package task

import (
	"errors"
	"server/database/repository"
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
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return err
	}

	subtasks, err := ts.db.TaskRepository.GetAll(repository.WithWhere("parent_task_id = ?", task.ID))
	if err != nil {
		return err
	}

	for _, subtask := range subtasks {
		err = ts.DeleteTask(subtask.ID)
		if err != nil {
			return err
		}
	}

	comments, err := ts.db.CommentRepository.GetAll(repository.WithWhere("task_id = ?", task.ID))
	if err != nil {
		return err
	}

	for _, comment := range comments {
		err = ts.cs.DeleteComment(comment.ID)
		if err != nil {
			return err
		}
	}

	dstLinks, err := ts.db.TaskLinkRepository.GetAll(repository.WithWhere("dst_task_id = ?", task.ID))
	if err != nil {
		return err
	}

	for _, dstLink := range dstLinks {
		err = ts.DeleteTaskLink(dstLink.ID)
		if err != nil {
			return err
		}
	}

	srcLinks, err := ts.db.TaskLinkRepository.GetAll(repository.WithWhere("src_task_id = ?", task.ID))
	if err != nil {
		return err
	}

	for _, srcLink := range srcLinks {
		err = ts.DeleteTaskLink(srcLink.ID)
		if err != nil {
			return err
		}
	}

	externalLinks, err := ts.db.TaskExternalLinkRepository.GetAll(repository.WithWhere("task_id = ?", task.ID))
	if err != nil {
		return err
	}

	for _, externalLink := range externalLinks {
		err = ts.DeleteTaskExternalLink(externalLink.ID)
		if err != nil {
			return err
		}
	}

	files, err := ts.db.FileRepository.GetAll(repository.WithWhere("task_id = ?", task.ID))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = ts.DeleteFile(file.ID)
		if err != nil {
			return err
		}
	}

	activities, err := ts.db.TaskActivityRepository.GetAll(repository.WithWhere("task_id = ?", task.ID))
	if err != nil {
		return err
	}

	for _, activity := range activities {
		err = ts.DeleteTaskActivity(activity.ID)
		if err != nil {
			return err
		}
	}

	return ts.db.TaskRepository.Delete(taskID)
}

func (ts *TaskService) DeleteTaskActivity(activityID uint) error {
	return ts.db.TaskActivityRepository.Delete(activityID)
}
