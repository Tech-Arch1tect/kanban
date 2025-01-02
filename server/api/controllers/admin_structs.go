package controllers

import "server/models"

type AdminRemoveUserResponse struct {
	Message string `json:"message"`
}

type AdminListUsersResponse struct {
	Users        []models.User `json:"users"`
	Page         int           `json:"page"`
	PageSize     int           `json:"page_size"`
	TotalPages   int           `json:"total_pages"`
	TotalRecords int           `json:"total_records"`
}

type AdminSearchUsersRequest struct {
	models.PaginationParamsRequest
	Search string `form:"search"`
}

type AdminUpdateUserRoleRequest struct {
	Role string `json:"role" binding:"required,oneof=admin user"`
}

type AdminUpdateUserRoleResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}
