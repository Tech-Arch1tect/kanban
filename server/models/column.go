package models

type Column struct {
	Model
	Name    string `gorm:"not null" json:"name" binding:"required"`
	BoardID uint   `json:"board_id"`
	Order   int    `json:"order"`
}
