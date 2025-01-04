package column

import "server/models"

type CreateColumnRequest struct {
	BoardID uint   `json:"board_id"`
	Name    string `json:"name"`
}

type CreateColumnResponse struct {
	Column models.Column `json:"column"`
}

type DeleteColumnRequest struct {
	ID uint `json:"id"`
}

type DeleteColumnResponse struct {
	Column models.Column `json:"column"`
}

type EditColumnRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type EditColumnResponse struct {
	Column models.Column `json:"column"`
}

type MoveColumnRequest struct {
	ID         uint   `json:"id" binding:"required"`
	RelativeID uint   `json:"relative_id" binding:"required"`
	Direction  string `json:"direction" binding:"required,oneof=before after"`
}

type MoveColumnResponse struct {
	Column models.Column `json:"column"`
}
