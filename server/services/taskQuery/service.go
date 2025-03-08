package taskquery

import (
	"errors"
	"fmt"
	"strings"

	"server/database/repository"

	"gorm.io/gorm"
)

type TaskQueryService struct {
	db     *repository.Database
	parser *QueryParser
}

func NewTaskQueryService(db *repository.Database) *TaskQueryService {
	return &TaskQueryService{
		db:     db,
		parser: NewQueryParser(),
	}
}

func (tqs *TaskQueryService) BuildQuery(boardID uint, query string) ([]repository.QueryOption, error) {
	ctx := tqs.parser.Parse(query, boardID)
	var qopts []repository.QueryOption

	qopts = append(qopts, repository.WithWhere("board_id = ?", boardID))
	qopts = append(qopts, repository.WithWhere("parent_task_id IS NULL"))

	if len(ctx.Statuses) > 0 {
		qopts = append(qopts, repository.WithWhere("LOWER(status) IN ?", ctx.Statuses))
	}

	if ctx.Assignee != "" {
		if strings.ToLower(ctx.Assignee) == "unassigned" {
			qopts = append(qopts, repository.WithWhere("assignee_id = 0"))
		} else {
			user, err := tqs.db.UserRepository.GetFirst(
				repository.WithWhere("LOWER(username) = ?", strings.ToLower(ctx.Assignee)),
			)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					qopts = append(qopts, repository.WithWhere("1 = 0"))
				} else {
					return nil, err
				}
			} else {
				qopts = append(qopts, repository.WithWhere("assignee_id = ?", user.ID))
			}
		}
	}

	if ctx.Creator != "" {
		user, err := tqs.db.UserRepository.GetFirst(
			repository.WithWhere("LOWER(username) = ?", strings.ToLower(ctx.Creator)),
		)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				qopts = append(qopts, repository.WithWhere("1 = 0"))
			} else {
				return nil, err
			}
		} else {
			qopts = append(qopts, repository.WithWhere("creator_id = ?", user.ID))
		}
	}

	if ctx.Title != "" {
		titleLike := fmt.Sprintf("%%%s%%", strings.ToLower(ctx.Title))
		qopts = append(qopts, repository.WithWhere("LOWER(title) LIKE ?", titleLike))
	}

	if len(ctx.FreeText) > 0 {
		combined := strings.Join(ctx.FreeText, " ")
		likeValue := fmt.Sprintf("%%%s%%", strings.ToLower(combined))
		qopts = append(qopts, repository.WithCustom(func(db *gorm.DB) *gorm.DB {
			return db.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", likeValue, likeValue)
		}))
	}

	qopts = append(qopts, repository.WithPreload("Assignee", "Subtasks", "Subtasks.Assignee"))
	return qopts, nil
}
