package boardController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

type CreateBoardRequest struct {
	Name string `json:"name"`
}

type CreateBoardResponse struct {
	Board models.Board `json:"board"`
}

// CreateBoard godoc
// @Summary Create a new board
// @Description Create a new board with the given name
// @Tags board
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

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	board := models.Board{
		Name:    request.Name,
		OwnerID: user.ID,
	}

	err = database.DB.BoardRepository.Create(&board)
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