package sampleDataController

import (
	"fmt"
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

// this is a controller which is used to insert sample data into the database for a specific board
// Fake tasks are inserted into the database with either status "open" or "closed"

type InsertSampleDataRequest struct {
	BoardID uint `json:"board_id"`
	NumTasks int `json:"num_tasks"`
}

type InsertSampleDataResponse struct {
	Message string `json:"message"`
}

// @Summary Insert sample data into the database for a specific board
// @Description Insert sample data into the database for a specific board
// @Tags sample-data
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body InsertSampleDataRequest true "Insert sample data request"
// @Success 200 {object} InsertSampleDataResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/sample-data/insert [post]
func InsertSampleData(c *gin.Context) {
	var request InsertSampleDataRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.BoardRepository.GetByID(request.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	simlanes, err := database.DB.SwimlaneRepository.GetByBoardID(request.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	columns, err := database.DB.ColumnRepository.GetByBoardID(request.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < request.NumTasks; i++ {
		status := "open"
		if rand.Intn(2) == 1 {
			status = "closed"
		}
		column := columns[rand.Intn(len(columns))]
		swimlane := simlanes[rand.Intn(len(simlanes))]
		position, err := database.DB.TaskRepository.GetPosition(column.ID, swimlane.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fakeTask := models.Task{
			Title: fmt.Sprintf("Fake Task %d", i),
			Status: status,
			BoardID: request.BoardID,
			Position: position + 1,
			Description: "This is a fake task created by the sample data controller for testing purposes. It should be deleted after testing is complete. lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			ColumnID: column.ID,
			SwimlaneID: swimlane.ID,
		}
		database.DB.TaskRepository.Create(&fakeTask)
	}

	// reposition all tasks
	for _, column := range columns {
		for _, swimlane := range simlanes {
			database.DB.TaskRepository.RePositionAll(column.ID, swimlane.ID)
		}
	}

	c.JSON(http.StatusOK, InsertSampleDataResponse{Message: "Sample data inserted"})
}
