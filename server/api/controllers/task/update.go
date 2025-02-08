package task

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

type UpdateTaskTitleRequest struct {
	TaskID uint   `json:"task_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTaskTitleResponse struct {
	Task models.Task `json:"task"`
}

// @Summary Update a task title
// @Description Update a task title
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UpdateTaskTitleRequest true "Update task title request"
// @Success 200 {object} UpdateTaskTitleResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/update/title [post]
func (tc *TaskController) UpdateTaskTitle(c *gin.Context) {
	var request UpdateTaskTitleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.UpdateTaskTitle(user.ID, request.TaskID, request.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.title", task, user)

	c.JSON(http.StatusOK, UpdateTaskTitleResponse{Task: task})
}

type UpdateTaskDescriptionRequest struct {
	TaskID      uint   `json:"task_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTaskDescriptionResponse struct {
	Task models.Task `json:"task"`
}

// @Summary Update a task description
// @Description Update a task description
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UpdateTaskDescriptionRequest true "Update task description request"
// @Success 200 {object} UpdateTaskDescriptionResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/update/description [post]
func (tc *TaskController) UpdateTaskDescription(c *gin.Context) {
	var request UpdateTaskDescriptionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.UpdateTaskDescription(user.ID, request.TaskID, request.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.description", task, user)

	c.JSON(http.StatusOK, UpdateTaskDescriptionResponse{Task: task})
}

type UpdateTaskStatusRequest struct {
	TaskID uint   `json:"task_id" binding:"required"`
	Status string `json:"status" binding:"oneof=open closed,required"`
}

type UpdateTaskStatusResponse struct {
	Task models.Task `json:"task"`
}

// @Summary Update a task status
// @Description Update a task status
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UpdateTaskStatusRequest true "Update task status request"
// @Success 200 {object} UpdateTaskStatusResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/update/status [post]
func (tc *TaskController) UpdateTaskStatus(c *gin.Context) {
	var request UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.UpdateTaskStatus(user.ID, request.TaskID, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.status", task, user)

	c.JSON(http.StatusOK, UpdateTaskStatusResponse{Task: task})
}

type UpdateTaskAssigneeRequest struct {
	TaskID     uint `json:"task_id" binding:"required"`
	AssigneeID uint `json:"assignee_id"` // not required so we can unassign tasks
}

type UpdateTaskAssigneeResponse struct {
	Task models.Task `json:"task"`
}

// @Summary Update a task assignee
// @Description Update a task assignee
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UpdateTaskAssigneeRequest true "Update task assignee request"
// @Success 200 {object} UpdateTaskAssigneeResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/update/assignee [post]
func (tc *TaskController) UpdateTaskAssignee(c *gin.Context) {
	var request UpdateTaskAssigneeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.UpdateTaskAssignee(user.ID, request.TaskID, request.AssigneeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.assignee", task, user)

	c.JSON(http.StatusOK, UpdateTaskAssigneeResponse{Task: task})
}
