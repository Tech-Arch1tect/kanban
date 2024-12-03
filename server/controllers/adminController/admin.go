package adminController

import (
	"net/http"
	"server/database"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RemoveUserResponse struct {
	Message string `json:"message"`
}

// RemoveUser godoc
// @Summary Remove a user by ID
// @Description Remove a user from the database by providing their ID
// @Tags admin
// @Security cookieAuth
// @Security csrf
// @Param id path string true "User ID"
// @Success 200 {object} RemoveUserResponse "message: user removed"
// @Failure 500 {object} models.ErrorResponse "error: failed to remove user"
// @Router /api/v1/admin/users/{id} [delete]
func RemoveUser(c *gin.Context) {
	userID := c.Param("id")
	uintUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	if err := database.DB.UserRepository.Delete(uint(uintUserID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove user"})
		return
	}
	c.JSON(http.StatusOK, RemoveUserResponse{Message: "user removed"})
}

type ListUsersResponse struct {
	Users        []models.User `json:"users"`
	Page         int           `json:"page"`
	PageSize     int           `json:"page_size"`
	TotalPages   int           `json:"total_pages"`
	TotalRecords int           `json:"total_records"`
}

type SearchUsersRequest struct {
	models.PaginationParamsRequest
	Search string `form:"search"`
}

// ListUsers godoc
// @Summary List users with pagination and optional search
// @Description List all users with optional pagination parameters and search functionality
// @Tags admin
// @Security cookieAuth
// @Param request query SearchUsersRequest true "Search users"
// @Success 200 {object} ListUsersResponse
// @Failure 400 {object} models.ErrorResponse "error: Invalid page or page size"
// @Failure 500 {object} models.ErrorResponse "error: failed to list users"
// @Router /api/v1/admin/users [get]
func ListUsers(c *gin.Context) {
	var req SearchUsersRequest
	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page or page size"})
		return
	}

	users, totalRecords, err := database.DB.UserRepository.PaginatedSearch(req.Page, req.PageSize, req.Search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	totalPages := (int(totalRecords) + req.PageSize - 1) / req.PageSize

	c.JSON(http.StatusOK, ListUsersResponse{
		Users:        users,
		Page:         req.Page,
		PageSize:     req.PageSize,
		TotalPages:   totalPages,
		TotalRecords: int(totalRecords),
	})
}

type UpdateUserRoleRequest struct {
	Role string `json:"role"`
}

type UpdateUserRoleResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

// UpdateUserRole godoc
// @Summary Update user role by ID
// @Description Update the role of a user identified by their ID with the provided role
// @Tags admin
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRoleRequest true "New Role"
// @Success 200 {object} UpdateUserRoleResponse "message: user role updated, user: updated user details"
// @Failure 400 {object} models.ErrorResponse "error: invalid input or invalid role"
// @Failure 404 {object} models.ErrorResponse "error: user not found"
// @Failure 500 {object} models.ErrorResponse "error: failed to update user role"
// @Router /api/v1/admin/users/{id}/role [put]
func UpdateUserRole(c *gin.Context) {
	var input UpdateUserRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userID := c.Param("id")
	uintUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	user, err := database.DB.UserRepository.GetByID(uint(uintUserID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	newrole := models.Role(input.Role)
	if newrole == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
		return
	}

	user.Role = newrole

	if err := database.DB.UserRepository.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user role"})
		return
	}
	c.JSON(http.StatusOK, UpdateUserRoleResponse{Message: "user role updated", User: user})
}
