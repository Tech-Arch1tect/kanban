package task

import "server/models"

type CreateTaskRequest struct {
	BoardID     uint   `json:"board_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SwimlaneID  uint   `json:"swimlane_id"`
	ColumnID    uint   `json:"column_id"`
	Status      string `json:"status"`
	AssigneeID  uint   `json:"assignee_id"`
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

type EditTaskRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	AssigneeID  uint   `json:"assignee_id"`
}

type EditTaskResponse struct {
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
