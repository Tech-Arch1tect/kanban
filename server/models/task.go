package models

import (
	"errors"
	"slices"
)

type Task struct {
	Model
	BoardID     uint      `json:"board_id"`
	Board       Board     `gorm:"foreignKey:BoardID" json:"board"`
	Swimlane    Swimlane  `gorm:"foreignKey:SwimlaneID" json:"swimlane"`
	SwimlaneID  uint      `json:"swimlane_id"`
	ColumnID    uint      `json:"column_id"`
	Column      Column    `gorm:"foreignKey:ColumnID" json:"column"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Comments    []Comment `gorm:"foreignKey:TaskID" json:"comments"`
	Position    int       `json:"position"`
	CreatorID   uint      `json:"creator_id"`
	Creator     User      `gorm:"foreignKey:CreatorID" json:"creator"`
	AssigneeID  uint      `json:"assignee_id"`
	Assignee    User      `gorm:"foreignKey:AssigneeID" json:"assignee"`
}

var allowedStatuses = []string{"open", "closed"}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}

	if !slices.Contains(allowedStatuses, t.Status) {
		return errors.New("status is not valid")
	}

	return nil
}

type Comment struct {
	Model
	User   User   `gorm:"foreignKey:UserID" json:"user"`
	UserID uint   `json:"user_id"`
	Task   Task   `gorm:"foreignKey:TaskID" json:"task"`
	TaskID uint   `json:"task_id"`
	Text   string `json:"text"`
}
