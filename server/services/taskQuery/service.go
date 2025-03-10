// service.go - Provides the service for building SQL queries from query strings.
package taskquery

import (
	"errors"
	"fmt"
	"server/database/repository"
	"strings"

	"gorm.io/gorm"
)

type TaskQueryService struct {
	db *repository.Database
}

func NewTaskQueryService(db *repository.Database) *TaskQueryService {
	return &TaskQueryService{db: db}
}

func (tqs *TaskQueryService) BuildQuery(boardID uint, query string) ([]repository.QueryOption, error) {
	parser := NewParser(query)
	ast, err := parser.ParseQuery()
	if err != nil {
		return nil, fmt.Errorf("failed to parse query: %w", err)
	}

	condition, args, err := tqs.convertAST(ast)
	if err != nil {
		return nil, fmt.Errorf("failed to convert query AST: %w", err)
	}

	var qopts []repository.QueryOption
	qopts = append(qopts, repository.WithWhere("board_id = ?", boardID))
	qopts = append(qopts, repository.WithWhere("parent_task_id IS NULL"))

	if strings.TrimSpace(condition) != "" {
		qopts = append(qopts, repository.WithWhere(condition, args...))
	}

	qopts = append(qopts, repository.WithPreload("Assignee", "Subtasks", "Subtasks.Assignee"))
	return qopts, nil
}

func (tqs *TaskQueryService) convertAST(ast Expression) (string, []interface{}, error) {
	switch node := ast.(type) {
	case *ComparisonExpr:
		if strings.ToLower(node.Field) == "search" {
			if strings.TrimSpace(strings.ToLower(node.Operator)) != "like" {
				return "", nil, fmt.Errorf("unsupported operator for search field: %s", node.Operator)
			}
			val, ok := node.Value.(string)
			if !ok {
				return "", nil, fmt.Errorf("search value must be a string")
			}
			searchVal := fmt.Sprintf("%%%s%%", strings.ToLower(val))
			condition := "(LOWER(title) LIKE ? OR LOWER(description) LIKE ?)"
			return condition, []interface{}{searchVal, searchVal}, nil
		}

		if strings.ToLower(node.Field) == "assignee" {
			return tqs.convertUserComparison("assignee_id", node)
		}
		if strings.ToLower(node.Field) == "creator" {
			return tqs.convertUserComparison("creator_id", node)
		}

		op := strings.TrimSpace(strings.ToLower(node.Operator))
		var sqlOp string
		switch op {
		case "==":
			sqlOp = "="
		case "!=":
			sqlOp = "!="
		case ">":
			sqlOp = ">"
		case "<":
			sqlOp = "<"
		case ">=":
			sqlOp = ">="
		case "<=":
			sqlOp = "<="
		case "like":
			sqlOp = "LIKE"
		default:
			return "", nil, fmt.Errorf("unsupported operator: %s", node.Operator)
		}
		condition := fmt.Sprintf("LOWER(%s) %s ?", node.Field, sqlOp)
		if valStr, ok := node.Value.(string); ok {
			if op == "like" {
				return condition, []interface{}{fmt.Sprintf("%%%s%%", strings.ToLower(valStr))}, nil
			}
			return condition, []interface{}{strings.ToLower(valStr)}, nil
		}
		return condition, []interface{}{node.Value}, nil

	case *LogicalExpr:
		leftCond, leftArgs, err := tqs.convertAST(node.Left)
		if err != nil {
			return "", nil, err
		}
		rightCond, rightArgs, err := tqs.convertAST(node.Right)
		if err != nil {
			return "", nil, err
		}
		combinedCond := fmt.Sprintf("(%s %s %s)", leftCond, node.Operator, rightCond)
		args := append(leftArgs, rightArgs...)
		return combinedCond, args, nil

	default:
		return "", nil, errors.New("unsupported expression type")
	}
}

func (tqs *TaskQueryService) convertUserComparison(dbField string, node *ComparisonExpr) (string, []interface{}, error) {
	valStr, ok := node.Value.(string)
	if !ok {
		return "", nil, fmt.Errorf("%s query value must be a string", dbField)
	}
	if strings.ToLower(valStr) == "unassigned" && strings.ToLower(dbField) == "assignee_id" {
		switch node.Operator {
		case "==":
			return "assignee_id = 0", nil, nil
		case "!=":
			return "assignee_id != 0", nil, nil
		default:
			return "", nil, fmt.Errorf("unsupported operator for assignee: %s", node.Operator)
		}
	}

	user, err := tqs.db.UserRepository.GetFirst(
		repository.WithWhere("LOWER(username) = ?", strings.ToLower(valStr)),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "1 = 0", nil, nil
		}
		return "", nil, err
	}

	var condition string
	switch node.Operator {
	case "==":
		condition = fmt.Sprintf("%s = ?", dbField)
	case "!=":
		condition = fmt.Sprintf("%s != ?", dbField)
	default:
		return "", nil, fmt.Errorf("unsupported operator for user field: %s", node.Operator)
	}
	return condition, []interface{}{user.ID}, nil
}
