package controllers

import (
	"net/http"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AdminController struct{}

// RemoveUser godoc
// @Summary Remove a user by ID
// @Description Remove a user from the database by providing their ID
// @Tags admin
// @Security cookieAuth
// @Security csrf
// @Param id path string true "User ID"
// @Success 200 {object} AdminRemoveUserResponse "message: user removed"
// @Failure 500 {object} models.ErrorResponse "error: failed to remove user"
// @Router /api/v1/admin/users/{id} [delete]
func (a *AdminController) RemoveUser(c *gin.Context) {
	userID := c.Param("id")
	uintUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := adminService.RemoveUser(uint(uintUserID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove user"})
		return
	}
	c.JSON(http.StatusOK, AdminRemoveUserResponse{Message: "user removed"})
}

// ListUsers godoc
// @Summary List users with pagination and optional search
// @Description List all users with optional pagination parameters and search functionality
// @Tags admin
// @Security cookieAuth
// @Param request query AdminSearchUsersRequest true "Search users"
// @Success 200 {object} AdminListUsersResponse
// @Failure 400 {object} models.ErrorResponse "error: Invalid page or page size"
// @Failure 500 {object} models.ErrorResponse "error: failed to list users"
// @Router /api/v1/admin/users [get]
func (a *AdminController) ListUsers(c *gin.Context) {
	var req AdminSearchUsersRequest
	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page or page size"})
		return
	}

	result, err := adminService.ListUsers(req.Page, req.PageSize, req.Search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	c.JSON(http.StatusOK, AdminListUsersResponse{
		Users:        result.Users,
		Page:         req.Page,
		PageSize:     req.PageSize,
		TotalPages:   result.TotalPages,
		TotalRecords: result.TotalRecords,
	})
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
// @Param user body AdminUpdateUserRoleRequest true "New Role"
// @Success 200 {object} AdminUpdateUserRoleResponse "message: user role updated, user: updated user details"
// @Failure 400 {object} models.ErrorResponse "error: invalid input or invalid role"
// @Failure 404 {object} models.ErrorResponse "error: user not found"
// @Failure 500 {object} models.ErrorResponse "error: failed to update user role"
// @Router /api/v1/admin/users/{id}/role [put]
func (a *AdminController) UpdateUserRole(c *gin.Context) {
	var input AdminUpdateUserRoleRequest
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

	user, err := adminService.UpdateUserRole(uint(uintUserID), models.Role(input.Role))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "user not found" {
			status = http.StatusNotFound
		} else if err.Error() == "invalid role" {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AdminUpdateUserRoleResponse{Message: "user role updated", User: user})
}
