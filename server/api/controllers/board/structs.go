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

type DeleteBoardRequest struct {
	ID uint `json:"id" binding:"required"`
}

type DeleteBoardResponse struct {
	Message string `json:"message"`
}

type GetBoardRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type GetBoardResponse struct {
	Board models.Board `json:"board"`
}

type GetBoardBySlugRequest struct {
	Slug string `uri:"slug" binding:"required"`
}

type GetBoardBySlugResponse struct {
	Board models.Board `json:"board"`
}

type ListBoardsResponse struct {
	Boards []models.Board `json:"boards"`
}

type GetUsersWithAccessToBoardRequest struct {
	BoardID uint `uri:"board_id" binding:"required"`
}

type GetUsersWithAccessToBoardResponse struct {
	Users []models.User `json:"users"`
}
