package task

import (
	"errors"
	"sort"

	"server/database/repository"
	"server/models"
	"server/services/role"

	"gorm.io/gorm"
)

func (ts *TaskService) CommonPreloadFields() []string {
	return []string{
		"Board", "Swimlane", "Column", "Creator", "Assignee", "Comments", "Comments.User", "Files",
		"SrcLinks", "DstLinks", "SrcLinks.SrcTask", "SrcLinks.DstTask", "DstLinks.DstTask", "DstLinks.SrcTask",
		"ExternalLinks", "ParentTask", "Subtasks", "Subtasks.Board", "Subtasks.Swimlane", "Subtasks.Column",
		"Subtasks.Creator", "Subtasks.Assignee", "Subtasks.Comments", "Subtasks.Comments.User", "Subtasks.Files",
		"Subtasks.SrcLinks", "Subtasks.DstLinks", "Subtasks.SrcLinks.SrcTask", "Subtasks.SrcLinks.DstTask",
		"Subtasks.DstLinks.DstTask", "Subtasks.DstLinks.SrcTask", "Subtasks.ExternalLinks", "Subtasks.ParentTask",
		"Comments.Reactions", "Comments.Reactions.User",
	}
}

func (ts *TaskService) GetTask(userID, taskID uint) (models.Task, error) {
	task, err := ts.db.TaskRepository.GetFirst(
		repository.WithWhere("id = ?", taskID),
		repository.WithPreload(ts.CommonPreloadFields()...),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	var newDstLinks []models.TaskLinks
	for _, link := range task.DstLinks {
		linkType := repository.InverseLinkTypeMap[link.LinkType]
		newDstLinks = append(newDstLinks, models.TaskLinks{
			Model:     models.Model{ID: link.ID},
			SrcTaskID: link.DstTaskID,
			SrcTask:   link.DstTask,
			DstTaskID: link.SrcTaskID,
			DstTask:   link.SrcTask,
			LinkType:  string(linkType),
		})
	}
	task.DstLinks = newDstLinks

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return models.Task{}, errors.New("forbidden")
	}

	return task, nil
}

func (ts *TaskService) GetTasksWithQuery(userID uint, boardID uint, query string) ([]models.Task, error) {
	can, _ := ts.rs.CheckRole(userID, boardID, role.ReaderRole)
	if !can {
		return nil, errors.New("forbidden")
	}

	qopts, err := ts.tqs.BuildQuery(boardID, query)
	if err != nil {
		return nil, err
	}

	tasks, err := ts.db.TaskRepository.GetAll(qopts...)
	if err != nil {
		return nil, err
	}

	ts.SortTasks(tasks)
	return tasks, nil
}

func (ts *TaskService) SortTasks(tasks []models.Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Position < tasks[j].Position
	})
}

func (ts *TaskService) GetTaskActivities(userID uint, taskID uint, page, pageSize int) (taskActivities []models.TaskActivity, totalRecords int, totalPages int, err error) {
	task, err := ts.db.TaskRepository.GetFirst(repository.WithWhere("id = ?", taskID))
	if err != nil {
		return nil, 0, 0, err
	}

	can, _ := ts.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if !can {
		return nil, 0, 0, errors.New("forbidden")
	}

	taskActivities, totalRecords, totalPages, err = ts.tas.GetPaginatedTaskActivities(page, pageSize)
	if err != nil {
		return nil, 0, 0, err
	}
	return taskActivities, totalRecords, totalPages, nil
}
