package board

import "server/models"

type CreateBoardRequest struct {
	Name      string   `json:"name" binding:"required"`
	Swimlanes []string `json:"swimlanes" binding:"required"`
	Columns   []string `json:"columns" binding:"required"`
	Slug      string   `json:"slug" binding:"required"`
}

type CreateBoardResponse struct {
	Board models.Board `json:"board"`
}
