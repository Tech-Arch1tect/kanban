package task

import (
	"errors"
	"fmt"
	"regexp"
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

	statuses, assigneeUsername, title, searchTerm := parseQuery(query)
	var qopts []repository.QueryOption

	qopts = append(qopts, repository.WithWhere("board_id = ?", boardID))
	qopts = append(qopts, repository.WithWhere("parent_task_id IS NULL"))

	if len(statuses) > 0 {
		qopts = append(qopts, repository.WithWhere("LOWER(status) IN ?", statuses))
	}
	if assigneeUsername != "" {
		if assigneeUsername == "unassigned" {
			qopts = append(qopts, repository.WithWhere("assignee_id = 0"))
		} else {
			user, err := ts.db.UserRepository.GetFirst(
				repository.WithWhere("LOWER(username) = ?", strings.ToLower(assigneeUsername)),
			)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return []models.Task{}, nil
				}
				return nil, err
			}
			qopts = append(qopts, repository.WithWhere("assignee_id = ?", user.ID))
		}
	}
	if title != "" {
		titleLike := fmt.Sprintf("%%%s%%", strings.ToLower(title))
		qopts = append(qopts, repository.WithWhere("LOWER(title) LIKE ?", titleLike))
	}
	if searchTerm != "" {
		likeValue := fmt.Sprintf("%%%s%%", strings.ToLower(searchTerm))
		qopts = append(qopts, repository.WithCustom(func(db *gorm.DB) *gorm.DB {
			return db.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", likeValue, likeValue)
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
//   - `status:open|closed` → []string{"open", "closed"}
//   - `assignee:username`  → "username"
//   - `title:"my title"`    → "my title"
//   - Everything else is appended into one search string
func parseQuery(q string) (statuses []string, assignee string, title string, searchTerm string) {
	re := regexp.MustCompile(`\S+:"[^"]+"|\S+`)
	tokens := re.FindAllString(q, -1)
	var searchParts []string
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		lowerToken := strings.ToLower(token)
		switch {
		case strings.HasPrefix(lowerToken, "status:"):
			raw := token[len("status:"):]
			raw = strings.Trim(raw, "\"")
			rawStatuses := strings.Split(raw, "|")
			for _, s := range rawStatuses {
				statuses = append(statuses, strings.ToLower(strings.TrimSpace(s)))
			}
		case strings.HasPrefix(lowerToken, "assignee:"):
			assignee = strings.Trim(token[len("assignee:"):], "\"")
		case strings.HasPrefix(lowerToken, "title:"):
			title = strings.Trim(token[len("title:"):], "\"")
		default:
			searchParts = append(searchParts, token)
		}
	}
	if len(searchParts) > 0 {
		searchTerm = strings.Join(searchParts, " ")
	}
	return statuses, assignee, title, searchTerm
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
