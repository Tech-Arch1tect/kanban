package taskController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type GetTaskRequest struct {
	ID uint `json:"id"`
}

type GetTaskResponse struct {
	Task models.Task `json:"task"`
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
func GetTask(c *gin.Context) {
	var request GetTaskRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := database.DB.TaskRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanAccessBoard, task.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	c.JSON(http.StatusOK, GetTaskResponse{Task: task})
}
