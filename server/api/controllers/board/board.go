package board

import (
	"net/http"
	"server/database/repository"
	"server/internal/helpers"
	"server/services"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	bs *services.BoardService
	rs *services.RoleService
	db *repository.Database
	hs *helpers.HelperService
}

func NewBoardController(bs *services.BoardService, rs *services.RoleService, db *repository.Database, hs *helpers.HelperService) *BoardController {
	return &BoardController{
		bs: bs,
		rs: rs,
		db: db,
		hs: hs,
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

// DeleteBoard godoc
// @Summary Delete a board
// @Description Delete a board by ID
// @Tags boards
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteBoardRequest true "Board ID"
// @Success 200 {object} DeleteBoardResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/delete [post]
func (bc *BoardController) DeleteBoard(c *gin.Context) {
	var req DeleteBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.bs.DeleteBoard(req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeleteBoardResponse{Message: "Board deleted"})
}

// GetBoard godoc
// @Summary Get a board
// @Description Get a board by ID
// @Tags boards
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param id path string true "Board ID"
// @Success 200 {object} GetBoardResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/get/{id} [get]
func (bc *BoardController) GetBoard(c *gin.Context) {
	var req GetBoardRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := bc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	board, err := bc.bs.GetBoardWithPermissions(user.ID, req.ID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, GetBoardResponse{Board: board})
}

// GetBoardBySlug godoc
// @Summary Get a board by slug
// @Description Get a board by slug
// @Tags boards
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param slug path string true "Board Slug"
// @Success 200 {object} GetBoardBySlugResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/get-by-slug/{slug} [get]
func (bc *BoardController) GetBoardBySlug(c *gin.Context) {
	var req GetBoardBySlugRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := bc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	board, err := bc.bs.GetBoardBySlugWithPermissions(user.ID, req.Slug)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, GetBoardBySlugResponse{Board: board})
}

// ListBoards godoc
// @Summary List all boards
// @Description List all boards for the current user
// @Tags boards
// @Security cookieAuth
// @Success 200 {object} ListBoardsResponse "Boards"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/boards/list [get]
func (bc *BoardController) ListBoards(c *gin.Context) {
	user, err := bc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	boards, err := bc.bs.ListBoards(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ListBoardsResponse{Boards: boards})
}
