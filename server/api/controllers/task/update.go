package task

import (
	"net/http"
	"server/models"
	"time"

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

	task, oldTask, err := tc.ts.UpdateTaskTitle(user.ID, request.TaskID, request.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.title", oldTask, task, user)

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

	task, oldTask, err := tc.ts.UpdateTaskDescription(user.ID, request.TaskID, request.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.description", oldTask, task, user)

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

	task, oldTask, err := tc.ts.UpdateTaskStatus(user.ID, request.TaskID, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.status", oldTask, task, user)

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

	task, oldTask, err := tc.ts.UpdateTaskAssignee(user.ID, request.TaskID, request.AssigneeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.assignee", oldTask, task, user)

	c.JSON(http.StatusOK, UpdateTaskAssigneeResponse{Task: task})
}

type UpdateTaskDueDateRequest struct {
	TaskID  uint       `json:"task_id" binding:"required"`
	DueDate *time.Time `json:"due_date" binding:"omitempty"`
}

type UpdateTaskDueDateResponse struct {
	Task models.Task `json:"task"`
}

// @Summary Update a task due date
// @Description Update a task due date
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UpdateTaskDueDateRequest true "Update task due date request"
// @Success 200 {object} UpdateTaskDueDateResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/update-due-date [post]
func (tc *TaskController) UpdateTaskDueDate(c *gin.Context) {
	var request UpdateTaskDueDateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, oldTask, err := tc.ts.UpdateTaskDueDate(user.ID, request.TaskID, request.DueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.te.Publish("task.updated.due-date", oldTask, task, user)

	c.JSON(http.StatusOK, UpdateTaskDueDateResponse{Task: task})
}

type UpdateTaskColourRequest struct {
	TaskID uint   `json:"task_id" binding:"required"`
	Colour string `json:"colour" binding:"oneof=slate gray zinc neutral stone red orange amber yellow lime green emerald teal cyan sky blue indigo violet purple fuchsia pink rose,required"`
}

type UpdateTaskColourResponse struct {
	Task models.Task `json:"task"`
}

// @Summary Update a task colour
// @Description Update a task colour
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body UpdateTaskColourRequest true "Update task colour request"
// @Success 200 {object} UpdateTaskColourResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/update-colour [post]
func (tc *TaskController) UpdateTaskColour(c *gin.Context) {
	var request UpdateTaskColourRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, _, err := tc.ts.UpdateTaskColour(user.ID, request.TaskID, request.Colour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UpdateTaskColourResponse{Task: task})
}
