package models

type Settings struct {
	Model
	UserID uint   `gorm:"unique"`
	Theme  string `json:"theme"`
}
