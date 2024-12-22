package taskController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

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
func EditTask(c *gin.Context) {
	var request EditTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := database.DB.TaskRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanEditTask, task.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	if request.AssigneeID != 0 {
		_, err = database.DB.UserRepository.GetByID(request.AssigneeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	task.Title = request.Title
	task.Description = request.Description
	task.Status = request.Status
	task.AssigneeID = request.AssigneeID

	err = task.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.TaskRepository.Update(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, EditTaskResponse{Task: task})
}
