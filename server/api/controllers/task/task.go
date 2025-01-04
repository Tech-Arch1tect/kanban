package task

import (
	"net/http"
	"server/database/repository"
	"server/internal/helpers"
	"server/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	db *repository.Database
	ts *services.TaskService
	rs *services.RoleService
	hs *helpers.HelperService
}

func NewTaskController(db *repository.Database, ts *services.TaskService, rs *services.RoleService, hs *helpers.HelperService) *TaskController {
	return &TaskController{db: db, ts: ts, rs: rs, hs: hs}
}

// @Summary Create a task
// @Description Create a task
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateTaskRequest true "Create task request"
// @Success 200 {object} CreateTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/create [post]
func (tc *TaskController) CreateTask(c *gin.Context) {
	var request CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.CreateTask(user.ID, services.CreateTaskRequest{
		BoardID:     request.BoardID,
		Title:       request.Title,
		Description: request.Description,
		SwimlaneID:  request.SwimlaneID,
		ColumnID:    request.ColumnID,
		Status:      request.Status,
		AssigneeID:  request.AssigneeID,
	})
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, CreateTaskResponse{Task: task})
}

// @Summary Delete a task
// @Description Delete a task
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteTaskRequest true "Delete task request"
// @Success 200 {object} DeleteTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/delete [post]
func (tc *TaskController) DeleteTask(c *gin.Context) {
	var request DeleteTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.DeleteTask(user.ID, request.ID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, DeleteTaskResponse{Task: task})
}

// @Summary Edit a task
// @Description Edit a task
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body EditTaskRequest true "Edit task request"
// @Success 200 {object} EditTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/edit [post]
func (tc *TaskController) EditTask(c *gin.Context) {
	var request EditTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.EditTask(user.ID, services.EditTaskRequest{
		ID:          request.ID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		AssigneeID:  request.AssigneeID,
	})
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, EditTaskResponse{Task: task})
}

// @Summary Get a task
// @Description Get a task
// @Tags tasks
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param id path uint true "Task ID"
// @Success 200 {object} GetTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/get/{id} [get]
func (tc *TaskController) GetTask(c *gin.Context) {
	var request GetTaskRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.GetTask(user.ID, request.ID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, GetTaskResponse{Task: task})
}

// @Summary Move a task
// @Description Move a task to a different column and swimlane, and update its position
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body MoveTaskRequest true "Move task request"
// @Success 200 {object} MoveTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/move [post]
func (tc *TaskController) MoveTask(c *gin.Context) {
	var request MoveTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.ts.MoveTask(user.ID, services.MoveTaskRequest{
		TaskID:     request.TaskID,
		ColumnID:   request.ColumnID,
		SwimlaneID: request.SwimlaneID,
		Position:   request.Position,
	})
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, MoveTaskResponse{Task: task})
}

// @Summary Get tasks with a query
// @Description Get tasks with a query
// @Tags tasks
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param query path string true "Query"
// @Success 200 {object} GetTaskQueryResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/get-query/{query} [get]
func (tc *TaskController) GetTaskQuery(c *gin.Context) {
	var request GetTaskQueryRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tasks, err := tc.ts.GetTasksWithQuery(user.ID, request.BoardID, request.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetTaskQueryResponse{Tasks: tasks})
}
