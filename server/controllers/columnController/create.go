package columnController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type CreateColumnRequest struct {
	BoardID uint   `json:"board_id"`
	Name    string `json:"name"`
}

type CreateColumnResponse struct {
	Column models.Column `json:"column"`
}

// @Summary Create a column
// @Description Create a column for a board
// @Tags columns
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateColumnRequest true "Column details"
// @Success 200 {object} CreateColumnResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/columns/create [post]
func CreateColumn(c *gin.Context) {
	var request CreateColumnRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanEditBoard, request.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	nextOrder, err := database.DB.ColumnRepository.GetNextOrder(request.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	column := models.Column{
		BoardID: request.BoardID,
		Name:    request.Name,
		Order:   nextOrder,
	}

	err = database.DB.ColumnRepository.Create(&column)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateColumnResponse{Column: column})
}
