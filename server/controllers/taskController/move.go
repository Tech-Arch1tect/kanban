package taskController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type MoveTaskRequest struct {
	TaskID     uint `json:"task_id" binding:"required"`
	ColumnID   uint `json:"column_id" binding:"required"`
	SwimlaneID uint `json:"swimlane_id" binding:"required"`
	Position   int  `json:"position" binding:"number"`
}

type MoveTaskResponse struct {
	Task models.Task `json:"task"`
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
func MoveTask(c *gin.Context) {
	var request MoveTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := database.DB.TaskRepository.GetByID(request.TaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanAccessBoard, task.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	column, err := database.DB.ColumnRepository.GetByID(request.ColumnID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	swimlane, err := database.DB.SwimlaneRepository.GetByID(request.SwimlaneID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if column.BoardID != swimlane.BoardID || column.BoardID != task.BoardID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Column, swimlane, and task must belong to the same board"})
		return
	}

	tasks, err := database.DB.TaskRepository.GetByColumnAndSwimlane(request.ColumnID, request.SwimlaneID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var filteredTasks []models.Task
	for _, t := range tasks {
		if t.ID != task.ID {
			filteredTasks = append(filteredTasks, t)
		}
	}
	tasks = filteredTasks

	if request.Position > len(tasks) {
		request.Position = len(tasks)
	}

	tasks = append(tasks, models.Task{})
	copy(tasks[request.Position+1:], tasks[request.Position:])
	tasks[request.Position] = task

	for i := range tasks {
		tasks[i].Position = i
		tasks[i].ColumnID = request.ColumnID
		tasks[i].SwimlaneID = request.SwimlaneID
	}

	for i := range tasks {
		if err := database.DB.TaskRepository.Update(&tasks[i]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	updatedTask, err := database.DB.TaskRepository.GetByID(task.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.TaskRepository.RePositionAll(task.ColumnID, task.SwimlaneID)

	c.JSON(http.StatusOK, MoveTaskResponse{Task: updatedTask})
}
