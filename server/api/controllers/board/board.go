package board

import (
	"net/http"
	"server/database/repository"
	"server/internal/helpers"
	"server/models"
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

// @Summary Get users with access to a board
// @Description Get users with access to a board
// @Tags boards
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param board_id path uint true "Board ID"
// @Success 200 {object} GetUsersWithAccessToBoardResponse
// @Router /api/v1/boards/permissions/{board_id} [get]
func (bc *BoardController) GetUsersWithAccessToBoard(c *gin.Context) {
	var req GetUsersWithAccessToBoardRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := bc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users, err := bc.bs.GetUsersWithAccess(user.ID, req.BoardID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetUsersWithAccessToBoardResponse{Users: users})
}

type AddOrInviteUserToBoardRequest struct {
	BoardID uint   `json:"board_id" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Role    string `json:"role" binding:"required,oneof=admin member reader"`
}

type AddOrInviteUserToBoardResponse struct {
	Message string `json:"message"`
}

// @Summary Add or invite a user to a board
// @Description Add or invite a user to a board
// @Tags boards
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body AddOrInviteUserToBoardRequest true "Add or invite user to board request"
// @Success 200 {object} AddOrInviteUserToBoardResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/add-or-invite [post]
func (bc *BoardController) AddOrInviteUserToBoard(c *gin.Context) {
	var req AddOrInviteUserToBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := bc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	role := services.AppRole{
		Name: req.Role,
	}

	err = bc.bs.AddOrInviteUserToBoard(user.ID, req.BoardID, req.Email, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AddOrInviteUserToBoardResponse{Message: "User added or invited"})
}

type GetPendingInvitesRequest struct {
	BoardID uint `uri:"board_id" binding:"required"`
}

type GetPendingInvitesResponse struct {
	Invites []models.BoardInvite `json:"invites"`
}

// @Summary Get pending invites
// @Description Get pending invites
// @Tags boards
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param board_id path uint true "Board ID"
// @Success 200 {object} GetPendingInvitesResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/pending-invites/{board_id} [get]
func (bc *BoardController) GetPendingInvites(c *gin.Context) {
	var req GetPendingInvitesRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := bc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	invites, err := bc.bs.GetPendingInvites(user.ID, req.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetPendingInvitesResponse{Invites: invites})
}

type RemoveUserFromBoardRequest struct {
	BoardID uint `json:"board_id" binding:"required"`
	UserID  uint `json:"user_id" binding:"required"`
}

type RemoveUserFromBoardResponse struct {
	Message string `json:"message"`
}

// @Summary Remove a user from a board
// @Description Remove a user from a board
// @Tags boards
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body RemoveUserFromBoardRequest true "Remove user from board request"
// @Success 200 {object} RemoveUserFromBoardResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/remove-user [post]
func (bc *BoardController) RemoveUserFromBoard(c *gin.Context) {
	var req RemoveUserFromBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := bc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = bc.bs.RemoveUserFromBoard(user.ID, req.UserID, req.BoardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RemoveUserFromBoardResponse{Message: "User removed"})
}

type ChangeBoardRoleRequest struct {
	BoardID uint   `json:"board_id" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
	Role    string `json:"role" binding:"required,oneof=admin member reader"`
}

type ChangeBoardRoleResponse struct {
	Message string `json:"message"`
}

// @Summary Change a user's role in a board
// @Description Change a user's role in a board
// @Tags boards
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body ChangeBoardRoleRequest true "Change board role request"
// @Success 200 {object} ChangeBoardRoleResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/boards/change-role [post]
func (bc *BoardController) ChangeBoardRole(c *gin.Context) {
	var req ChangeBoardRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := bc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = bc.bs.ChangeBoardRole(user.ID, req.UserID, req.BoardID, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ChangeBoardRoleResponse{Message: "Role changed"})
}
