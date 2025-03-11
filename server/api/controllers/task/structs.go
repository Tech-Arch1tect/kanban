package task

import (
	"server/models"
	"time"
)

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
	Colour       string     `json:"colour" binding:"omitempty,oneof=slate gray zinc neutral stone red orange amber yellow lime green emerald teal cyan sky blue indigo violet purple fuchsia pink rose"`
}

type CreateTaskResponse struct {
	Task models.Task `json:"task"`
}

type DeleteTaskRequest struct {
	ID uint `json:"id"`
}

type DeleteTaskResponse struct {
	Task models.Task `json:"task"`
}

type GetTaskRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type GetTaskResponse struct {
	Task models.Task `json:"task"`
}

type MoveTaskRequest struct {
	TaskID     uint    `json:"task_id" binding:"required"`
	ColumnID   uint    `json:"column_id" binding:"required"`
	SwimlaneID uint    `json:"swimlane_id" binding:"required"`
	Position   float64 `json:"position" binding:"number"`
	MoveAfter  bool    `json:"move_after"`
}

type MoveTaskResponse struct {
	Task models.Task `json:"task"`
}

type GetTaskQueryRequest struct {
	BoardID uint   `uri:"board_id" binding:"required"`
	Query   string `uri:"query" binding:"required"`
}

type GetTaskQueryResponse struct {
	Tasks []models.Task `json:"tasks"`
}
