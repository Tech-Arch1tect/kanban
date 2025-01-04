package board

import (
	"net/http"
	"server/database/repository"
	"server/services"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	bs *services.BoardService
	ps *services.PermissionService
	db *repository.Database
}

func NewBoardController(bs *services.BoardService, ps *services.PermissionService, db *repository.Database) *BoardController {
	return &BoardController{
		bs: bs,
		ps: ps,
		db: db,
	}
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
func (bc *BoardController) CreateBoard(c *gin.Context) {
	var request CreateBoardRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	board, err := bc.bs.CreateBoard(request.Name, request.Slug, request.Swimlanes, request.Columns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateBoardResponse{Board: board})
}
