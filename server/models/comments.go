package models

type Comment struct {
	Model
	User   User   `gorm:"foreignKey:UserID" json:"user"`
	UserID uint   `json:"user_id"`
	TaskID uint   `json:"task_id"`
	Text   string `json:"text"`
}
