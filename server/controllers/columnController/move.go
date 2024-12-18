package columnController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"

	"github.com/gin-gonic/gin"
)

type MoveColumnRequest struct {
	ID         uint   `json:"id" binding:"required"`
	RelativeID uint   `json:"relative_id" binding:"required"`
	Direction  string `json:"direction" binding:"required,oneof=before after"`
}

type MoveColumnResponse struct {
	Column models.Column `json:"column"`
}

// @Summary Move a column
// @Description Move a column relative to another column
// @Tags columns
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body MoveColumnRequest true "Column ID and relative column ID and direction"
// @Success 200 {object} MoveColumnResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/columns/move [post]
func MoveColumn(c *gin.Context) {
	var request MoveColumnRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	column, err := database.DB.ColumnRepository.GetByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch column"})
		return
	}
	relativeColumn, err := database.DB.ColumnRepository.GetByID(request.RelativeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch relative column"})
		return
	}

	if column.BoardID != relativeColumn.BoardID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Columns are not on the same board"})
		return
	}

	if can, _ := permissions.Can(c, permissions.CanEditBoard, column.BoardID); !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	allBoardColumns, err := database.DB.ColumnRepository.GetByBoardID(column.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch columns"})
		return
	}

	columnMap := make(map[uint]int)
	for i, c := range allBoardColumns {
		columnMap[c.ID] = i
	}

	currentIdx, relativeIdx := columnMap[column.ID], columnMap[relativeColumn.ID]

	targetIdx := relativeIdx
	if request.Direction == "before" {
		targetIdx++
	}

	allBoardColumns = append(allBoardColumns[:currentIdx], allBoardColumns[currentIdx+1:]...)

	if currentIdx < targetIdx {
		targetIdx--
	}

	allBoardColumns = append(allBoardColumns[:targetIdx], append([]models.Column{column}, allBoardColumns[targetIdx:]...)...)

	for i, col := range allBoardColumns {
		col.Order = i + 1
		if err := database.DB.ColumnRepository.Update(&col); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update column order"})
			return
		}
	}

	c.JSON(http.StatusOK, MoveColumnResponse{Column: column})
}
