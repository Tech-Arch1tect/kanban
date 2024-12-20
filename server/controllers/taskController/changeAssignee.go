package taskController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type ChangeAssigneeRequest struct {
	ID         uint `json:"id"`
	AssigneeID uint `json:"assignee_id"`
}

type ChangeAssigneeResponse struct {
	Task models.Task `json:"task"`
}

// @Summary Change assignee of a task
// @Description Change the assignee of a task
// @Tags tasks
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body ChangeAssigneeRequest true "Change assignee request"
// @Success 200 {object} ChangeAssigneeResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/change-assignee [post]
func ChangeAssignee(c *gin.Context) {
	var request ChangeAssigneeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanEditTask, request.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to change the assignee of this task"})
		return
	}

	task, err := database.DB.TaskRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	assignee, err := database.DB.UserRepository.GetByID(request.AssigneeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task.AssigneeID = assignee.ID

	err = database.DB.TaskRepository.Update(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ChangeAssigneeResponse{Task: task})
}
