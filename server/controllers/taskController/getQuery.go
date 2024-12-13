package taskController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

type GetTaskQueryRequest struct {
	Query string `json:"query"`
}

type GetTaskQueryResponse struct {
	Tasks []models.Task `json:"tasks"`
}

// @Summary Get tasks with a query
// @Description Get tasks with a query
// @Tags tasks
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param query body GetTaskQueryRequest true "Query"
// @Success 200 {object} GetTaskQueryResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/get-query [get]
func GetTaskQuery(c *gin.Context) {
	var request GetTaskQueryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tasks, err := database.DB.TaskRepository.GetWithQuery(request.Query, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetTaskQueryResponse{Tasks: tasks})
}
