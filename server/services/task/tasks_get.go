package task

import (
	"errors"
	"fmt"
	"sort"
	"strings"

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

	statuses, assigneeEmail, searchTerm := parseQuery(query)
	var qopts []repository.QueryOption

	qopts = append(qopts, repository.WithWhere("board_id = ?", boardID))
	qopts = append(qopts, repository.WithWhere("parent_task_id IS NULL"))

	if len(statuses) > 0 {
		qopts = append(qopts, repository.WithWhere("status IN ?", statuses))
	}
	if assigneeEmail != "" {
		user, err := ts.db.UserRepository.GetFirst(repository.WithWhere("email = ?", assigneeEmail))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return []models.Task{}, nil
			}
			return nil, err
		}
		qopts = append(qopts, repository.WithWhere("assignee_id = ?", user.ID))
	}
	if searchTerm != "" {
		likeValue := fmt.Sprintf("%%%s%%", searchTerm)
		qopts = append(qopts, repository.WithCustom(func(db *gorm.DB) *gorm.DB {
			return db.Where("title LIKE ? OR description LIKE ?", likeValue, likeValue)
		}))
	}

	qopts = append(qopts, repository.WithPreload("Assignee", "Subtasks", "Subtasks.Assignee"))
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

// parseQuery is a naive parser that extracts:
//   - `status:open|closed` → []string{"open","closed"}
//   - `assignee:email`  → "email"
//   - Everything else → appended into one search string
func parseQuery(q string) (statuses []string, assignee string, searchTerm string) {
	tokens := strings.Split(strings.TrimSpace(q), " ")
	var searchParts []string
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		switch {
		case strings.HasPrefix(token, "status:"):
			raw := strings.TrimPrefix(token, "status:")
			statuses = strings.Split(raw, "|")
		case strings.HasPrefix(token, "assignee:"):
			assignee = strings.TrimPrefix(token, "assignee:")
		default:
			searchParts = append(searchParts, token)
		}
	}
	if len(searchParts) > 0 {
		searchTerm = strings.Join(searchParts, " ")
	}
	return statuses, assignee, searchTerm
}
