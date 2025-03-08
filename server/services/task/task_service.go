package task

import (
	"server/config"
	"server/database/repository"
	"server/services/comment"
	"server/services/role"
	"server/services/taskActivity"
	taskquery "server/services/taskQuery"
	"time"
)

type TaskService struct {
	db     *repository.Database
	rs     *role.RoleService
	config *config.Config
	cs     *comment.CommentService
	tas    *taskActivity.TaskActivityService
	tqs    *taskquery.TaskQueryService
}

type CreateTaskRequest struct {
	ParentTaskID *uint      `json:"parent_task_id"`
	BoardID      uint       `json:"board_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	SwimlaneID   uint       `json:"swimlane_id"`
	ColumnID     uint       `json:"column_id"`
	Status       string     `json:"status"`
	AssigneeID   uint       `json:"assignee_id"`
	DueDate      *time.Time `json:"due_date"`
	Colour       string     `json:"colour"`
}

func NewTaskService(db *repository.Database, rs *role.RoleService, config *config.Config, cs *comment.CommentService, tas *taskActivity.TaskActivityService, tqs *taskquery.TaskQueryService) *TaskService {
	return &TaskService{db: db, rs: rs, config: config, cs: cs, tas: tas, tqs: tqs}
}
