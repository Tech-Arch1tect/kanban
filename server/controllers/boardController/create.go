package boardController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type CreateBoardRequest struct {
	Name      string   `json:"name"`
	Swimlanes []string `json:"swimlanes"`
	Columns   []string `json:"columns"`
}

type CreateBoardResponse struct {
	Board models.Board `json:"board"`
}

// CreateBoard godoc
// @Summary Create a new board
// @Description Create a new board with the given name
// @Tags boards
// @Security cookieAuth
// @Security csrf
// @Param request body CreateBoardRequest true "Board name"
// @Success 200 {object} CreateBoardResponse "Board created"
// @Failure 400 {object} models.ErrorResponse "Bad request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/boards/create [post]
func CreateBoard(c *gin.Context) {
	var request CreateBoardRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, user := permissions.Can(c, permissions.CanCreateBoard, 0)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	swimlanes := make([]models.Swimlane, len(request.Swimlanes))
	for i, name := range request.Swimlanes {
		swimlanes[i] = models.Swimlane{Name: name, Order: i}
	}

	columns := make([]models.Column, len(request.Columns))
	for i, name := range request.Columns {
		columns[i] = models.Column{Name: name, Order: i}
	}

	board := models.Board{
		Name:      request.Name,
		OwnerID:   user.ID,
		Swimlanes: swimlanes,
		Columns:   columns,
	}

	err := database.DB.BoardRepository.Create(&board)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	board, err = database.DB.BoardRepository.GetByID(board.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateBoardResponse{Board: board})
}
