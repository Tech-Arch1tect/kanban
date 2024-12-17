package columnController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type DeleteColumnRequest struct {
	ID uint `json:"id"`
}

type DeleteColumnResponse struct {
	Column models.Column `json:"column"`
}

// @Summary Delete a column
// @Description Delete a column by ID
// @Tags columns
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteColumnRequest true "Column ID"
// @Success 200 {object} DeleteColumnResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/columns/delete [post]
func DeleteColumn(c *gin.Context) {
	var request DeleteColumnRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	column, err := database.DB.ColumnRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	can, _ := permissions.Can(c, permissions.CanEditBoard, column.BoardID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	err = database.DB.ColumnRepository.Delete(column.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeleteColumnResponse{Column: column})
}
