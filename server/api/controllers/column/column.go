package column

import (
	"net/http"
	"server/database/repository"
	"server/internal/helpers"
	"server/services/column"
	"server/services/role"

	"github.com/gin-gonic/gin"
)

type ColumnController struct {
	db *repository.Database
	cs *column.ColumnService
	rs *role.RoleService
	hs *helpers.HelperService
}

func NewColumnController(db *repository.Database, cs *column.ColumnService, rs *role.RoleService, hs *helpers.HelperService) *ColumnController {
	return &ColumnController{db: db, cs: cs, rs: rs, hs: hs}
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
func (cc *ColumnController) CreateColumn(c *gin.Context) {
	var request CreateColumnRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	column, err := cc.cs.CreateColumn(user.ID, request.BoardID, request.Name)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, CreateColumnResponse{Column: column})
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
func (cc *ColumnController) DeleteColumn(c *gin.Context) {
	var request DeleteColumnRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	column, err := cc.cs.DeleteColumn(user.ID, request.ID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, DeleteColumnResponse{Column: column})
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
func (cc *ColumnController) EditColumn(c *gin.Context) {
	var request EditColumnRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	column, err := cc.cs.EditColumn(user.ID, request.ID, request.Name)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, EditColumnResponse{Column: column})
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
func (cc *ColumnController) MoveColumn(c *gin.Context) {
	var request MoveColumnRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	column, err := cc.cs.MoveColumn(user.ID, request.ID, request.RelativeID, request.Direction)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "columns are not on the same board" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, MoveColumnResponse{Column: column})
}
