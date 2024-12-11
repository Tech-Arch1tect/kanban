package models

type Task struct {
	Model
	BoardID     uint      `json:"board_id"`
	Swimlane    Swimlane  `gorm:"foreignKey:SwimlaneID" json:"swimlane"`
	SwimlaneID  uint      `json:"swimlane_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Comments    []Comment `gorm:"foreignKey:TaskID" json:"comments"`
}
