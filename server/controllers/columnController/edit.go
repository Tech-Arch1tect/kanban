package columnController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type EditColumnRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type EditColumnResponse struct {
	Column models.Column `json:"column"`
}

// @Summary Edit a column
// @Description Edit a column by ID
// @Tags columns
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body EditColumnRequest true "Column details"
// @Success 200 {object} EditColumnResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/columns/edit [post]
func EditColumn(c *gin.Context) {
	var request EditColumnRequest
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

	column.Name = request.Name

	err = database.DB.ColumnRepository.Update(&column)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, EditColumnResponse{Column: column})
}
