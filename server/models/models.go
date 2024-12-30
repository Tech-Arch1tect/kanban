package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleUser     Role = "user"
	RoleAdmin    Role = "admin"
	RoleDisabled Role = "disabled"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// everything below this line is for swagger / openapi spec purposes. NOT ALWAYS USED WITHIN THE APPLICATION (but sometimes)

// generic
type PaginationParamsRequest struct {
	Page     int `form:"page" binding:"omitempty,min=1"`
	PageSize int `form:"page_size" binding:"omitempty,min=1"`
}

// ErrorResponse is a generic error response struct for the API
type ErrorResponse struct {
	Error string `json:"error"`
}
