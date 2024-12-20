package taskController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

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
func CreateTask(c *gin.Context) {
	var request CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, user := permissions.Can(c, permissions.CanCreateTask, request.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	_, err := database.DB.SwimlaneRepository.GetByID(request.SwimlaneID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = database.DB.ColumnRepository.GetByID(request.ColumnID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	position, err := database.DB.TaskRepository.GetPosition(request.ColumnID, request.SwimlaneID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var assignee models.User
	if request.AssigneeID != 0 {
		assignee, err = database.DB.UserRepository.GetByID(request.AssigneeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	task := models.Task{
		BoardID:     request.BoardID,
		Title:       request.Title,
		Description: request.Description,
		SwimlaneID:  request.SwimlaneID,
		Status:      request.Status,
		ColumnID:    request.ColumnID,
		Position:    position + 1,
		CreatorID:   user.ID,
		AssigneeID:  assignee.ID,
	}

	err = task.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.TaskRepository.Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateTaskResponse{Task: task})
}
