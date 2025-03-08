package taskquery

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"server/database/repository"

	"gorm.io/gorm"
)

type TaskQueryService struct {
	db *repository.Database
}

func NewTaskQueryService(db *repository.Database) *TaskQueryService {
	return &TaskQueryService{db: db}
}

func (tqs *TaskQueryService) BuildQuery(boardID uint, query string) ([]repository.QueryOption, error) {
	statuses, assigneeUsername, title, searchTerm, creatorUsername := parseQuery(query)
	var qopts []repository.QueryOption

	qopts = append(qopts, repository.WithWhere("board_id = ?", boardID))
	qopts = append(qopts, repository.WithWhere("parent_task_id IS NULL"))

	if len(statuses) > 0 {
		qopts = append(qopts, repository.WithWhere("LOWER(status) IN ?", statuses))
	}
	if assigneeUsername != "" {
		if strings.ToLower(assigneeUsername) == "unassigned" {
			qopts = append(qopts, repository.WithWhere("assignee_id = 0"))
		} else {
			user, err := tqs.db.UserRepository.GetFirst(
				repository.WithWhere("LOWER(username) = ?", strings.ToLower(assigneeUsername)),
			)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// If the user doesnt exist, return no results
					qopts = append(qopts, repository.WithWhere("1 = 0"))
				} else {
					return nil, err
				}
			} else {
				qopts = append(qopts, repository.WithWhere("assignee_id = ?", user.ID))
			}
		}
	}
	if creatorUsername != "" {
		user, err := tqs.db.UserRepository.GetFirst(
			repository.WithWhere("LOWER(username) = ?", strings.ToLower(creatorUsername)),
		)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// if the user doesnt exist, return no results
				qopts = append(qopts, repository.WithWhere("1 = 0"))
			} else {
				return nil, err
			}
		} else {
			qopts = append(qopts, repository.WithWhere("creator_id = ?", user.ID))
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
	return qopts, nil
}

// parseQuery is a naive parser that extracts:
//   - `status:open|closed`   → []string{"open", "closed"}
//   - `assignee:username`    → "username"
//   - `creator:username`     → "username"
//   - `title:"my title"`      → "my title"
//   - Everything else is appended into one search string
func parseQuery(q string) (statuses []string, assignee string, title string, searchTerm string, creator string) {
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
		case strings.HasPrefix(lowerToken, "creator:"):
			creator = strings.Trim(token[len("creator:"):], "\"")
		case strings.HasPrefix(lowerToken, "title:"):
			title = strings.Trim(token[len("title:"):], "\"")
		default:
			searchParts = append(searchParts, token)
		}
	}
	if len(searchParts) > 0 {
		searchTerm = strings.Join(searchParts, " ")
	}
	return statuses, assignee, title, searchTerm, creator
}
