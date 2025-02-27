package task

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (ts *TaskService) GetTaskExternalLink(linkID uint) (models.TaskExternalLink, error) {
	link, err := ts.db.TaskExternalLinkRepository.GetByID(linkID, repository.WithPreload("Task"), repository.WithPreload("Task.Board"))
	if err != nil {
		return models.TaskExternalLink{}, err
	}
	return link, nil
}

func (ts *TaskService) CreateTaskExternalLink(userID uint, taskID uint, title string, url string) (models.TaskExternalLink, error) {
	task, err := ts.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.TaskExternalLink{}, err
	}
	if can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole); !can {
		return models.TaskExternalLink{}, errors.New("forbidden")
	}

	link := models.TaskExternalLink{
		URL:    url,
		Title:  title,
		TaskID: taskID,
	}
	if err = ts.db.TaskExternalLinkRepository.Create(&link); err != nil {
		return models.TaskExternalLink{}, err
	}
	return ts.GetTaskExternalLink(link.ID)
}

func (ts *TaskService) UpdateTaskExternalLink(userID uint, linkID uint, title string, url string) (models.TaskExternalLink, models.TaskExternalLink, error) {
	link, err := ts.db.TaskExternalLinkRepository.GetByID(linkID, repository.WithPreload("Task"))
	oldLink := link
	if err != nil {
		return models.TaskExternalLink{}, models.TaskExternalLink{}, err
	}
	if can, _ := ts.rs.CheckRole(userID, link.Task.BoardID, role.MemberRole); !can {
		return models.TaskExternalLink{}, models.TaskExternalLink{}, errors.New("forbidden")
	}

	link.URL = url
	link.Title = title
	if err = ts.db.TaskExternalLinkRepository.Update(&link); err != nil {
		return models.TaskExternalLink{}, models.TaskExternalLink{}, err
	}

	link, err = ts.GetTaskExternalLink(linkID)
	if err != nil {
		return models.TaskExternalLink{}, models.TaskExternalLink{}, err
	}
	return link, oldLink, nil
}

func (ts *TaskService) DeleteTaskExternalLinkRequest(userID uint, linkID uint) (models.TaskExternalLink, error) {
	link, err := ts.db.TaskExternalLinkRepository.GetByID(linkID, repository.WithPreload("Task"))
	if err != nil {
		return models.TaskExternalLink{}, err
	}
	if can, _ := ts.rs.CheckRole(userID, link.Task.BoardID, role.MemberRole); !can {
		return models.TaskExternalLink{}, errors.New("forbidden")
	}

	err = ts.DeleteTaskExternalLink(link.ID)
	if err != nil {
		return models.TaskExternalLink{}, err
	}
	return link, nil
}

func (ts *TaskService) DeleteTaskExternalLink(linkID uint) error {
	return ts.db.TaskExternalLinkRepository.Delete(linkID)
}
