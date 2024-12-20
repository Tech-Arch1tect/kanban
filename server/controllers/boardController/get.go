package boardController

import (
	"net/http"
	"server/database"
	"server/models"
	"server/permissions"
	"sort"

	"github.com/gin-gonic/gin"
)

type GetBoardRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type GetBoardResponse struct {
	Board models.Board `json:"board"`
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
func GetBoard(c *gin.Context) {
	var req GetBoardRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	can, _ := permissions.Can(c, permissions.CanAccessBoard, req.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	board, err := database.DB.BoardRepository.GetWithPreload(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorised"})
		return
	}

	sort.Slice(board.Swimlanes, func(i, j int) bool { return board.Swimlanes[i].Order < board.Swimlanes[j].Order })
	sort.Slice(board.Columns, func(i, j int) bool { return board.Columns[i].Order < board.Columns[j].Order })

	c.JSON(http.StatusOK, GetBoardResponse{Board: board})
}

type GetBoardBySlugRequest struct {
	Slug string `uri:"slug" binding:"required"`
}

type GetBoardBySlugResponse struct {
	Board models.Board `json:"board"`
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
func GetBoardBySlug(c *gin.Context) {
	var req GetBoardBySlugRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	board, err := database.DB.BoardRepository.GetBySlug(req.Slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorised"})
		return
	}

	can, _ := permissions.Can(c, permissions.CanAccessBoard, board.ID)
	if !can {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	sort.Slice(board.Swimlanes, func(i, j int) bool { return board.Swimlanes[i].Order < board.Swimlanes[j].Order })
	sort.Slice(board.Columns, func(i, j int) bool { return board.Columns[i].Order < board.Columns[j].Order })

	c.JSON(http.StatusOK, GetBoardBySlugResponse{Board: board})
}
