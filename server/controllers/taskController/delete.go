package taskController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type DeleteTaskRequest struct {
	ID uint `json:"id"`
}

type DeleteTaskResponse struct {
	Task models.Task `json:"task"`
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
func DeleteTask(c *gin.Context) {
	var request DeleteTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := database.DB.TaskRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanDeleteTask, task.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	err = database.DB.TaskRepository.Delete(task.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeleteTaskResponse{Task: task})
}
