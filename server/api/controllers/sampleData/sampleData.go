package sampleData

import (
	"fmt"
	"net/http"
	"server/database/repository"
	"server/models"
	"server/services"

	"math/rand"

	"github.com/gin-gonic/gin"
)

type SampleDataController struct {
	db *repository.Database
	ts *services.TaskService
}

func NewSampleDataController(db *repository.Database, ts *services.TaskService) *SampleDataController {
	return &SampleDataController{db: db, ts: ts}
}

// this is a controller which is used to insert sample data into the database for a specific board
// Fake tasks are inserted into the database with either status "open" or "closed"

type InsertSampleDataRequest struct {
	BoardID     uint `json:"board_id"`
	NumTasks    int  `json:"num_tasks"`
	NumComments int  `json:"num_comments"`
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
func (sc *SampleDataController) InsertSampleData(c *gin.Context) {
	var request InsertSampleDataRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := sc.db.BoardRepository.GetByID(request.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	simlanes, err := sc.db.SwimlaneRepository.GetAll(repository.WithWhere("board_id = ?", request.BoardID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	columns, err := sc.db.ColumnRepository.GetAll(repository.WithWhere("board_id = ?", request.BoardID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tasks []models.Task

	for i := 0; i < request.NumTasks; i++ {
		status := "open"
		if rand.Intn(2) == 1 {
			status = "closed"
		}
		column := columns[rand.Intn(len(columns))]
		swimlane := simlanes[rand.Intn(len(simlanes))]
		taskPosition, err := sc.db.TaskRepository.GetFirst(
			repository.WithWhere("column_id = ? AND swimlane_id = ?", column.ID, swimlane.ID),
			repository.WithOrder("position DESC"),
		)
		if err != nil && err.Error() != "record not found" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		position := 0
		if taskPosition.Position != 0 {
			position = taskPosition.Position + 1
		}

		fakeTask := models.Task{
			Title:       fmt.Sprintf("Fake Task %d", i),
			Status:      status,
			BoardID:     request.BoardID,
			Position:    position,
			Description: "This is a fake task created by the sample data controller for testing purposes. It should be deleted after testing is complete. lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			ColumnID:    column.ID,
			SwimlaneID:  swimlane.ID,
		}
		sc.db.TaskRepository.Create(&fakeTask)
		tasks = append(tasks, fakeTask)
	}

	// reposition all tasks
	for _, column := range columns {
		for _, swimlane := range simlanes {
			sc.ts.RePositionAll(column.ID, swimlane.ID)
		}
	}

	for i := 0; i < request.NumComments; i++ {
		task := tasks[rand.Intn(len(tasks))]
		comment := models.Comment{
			TaskID: task.ID,
			Text:   fmt.Sprintf("Fake Comment %d", i),
			UserID: task.CreatorID,
		}
		sc.db.CommentRepository.Create(&comment)
	}

	c.JSON(http.StatusOK, InsertSampleDataResponse{Message: "Sample data inserted"})
}
