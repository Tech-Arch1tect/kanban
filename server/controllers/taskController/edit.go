package taskController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type EditTaskRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SwimlaneID  uint   `json:"swimlane_id"`
	Status      string `json:"status"`
	ColumnID    uint   `json:"column_id"`
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

	column, err := database.DB.ColumnRepository.GetByID(request.ColumnID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if column.BoardID != task.BoardID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	swimlane, err := database.DB.SwimlaneRepository.GetByID(request.SwimlaneID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if swimlane.BoardID != task.BoardID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	task.Title = request.Title
	task.Description = request.Description
	task.SwimlaneID = request.SwimlaneID
	task.Status = request.Status
	task.ColumnID = request.ColumnID
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
