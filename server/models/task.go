package models

import (
	"errors"
	"server/helpers/commonHelpers"
)

type Task struct {
	Model
	BoardID     uint      `json:"board_id"`
	Swimlane    Swimlane  `gorm:"foreignKey:SwimlaneID" json:"swimlane"`
	SwimlaneID  uint      `json:"swimlane_id"`
	ColumnID    uint      `json:"column_id"`
	Column      Column    `gorm:"foreignKey:ColumnID" json:"column"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Comments    []Comment `gorm:"foreignKey:TaskID" json:"comments"`
}

var allowedStatuses = []string{"open", "closed"}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}

	if !commonHelpers.Contains(allowedStatuses, t.Status) {
		return errors.New("status is not valid")
	}

	return nil
}
