package models

import (
	"errors"
	"slices"
	"time"
)

type Task struct {
	Model
	ParentTaskID  *uint              `json:"parent_task_id"`
	ParentTask    *Task              `gorm:"foreignKey:ParentTaskID" json:"parent_task"`
	Subtasks      []*Task            `gorm:"foreignKey:ParentTaskID" json:"subtasks"`
	BoardID       uint               `json:"board_id"`
	Board         Board              `gorm:"foreignKey:BoardID" json:"board"`
	Swimlane      Swimlane           `gorm:"foreignKey:SwimlaneID" json:"swimlane"`
	SwimlaneID    uint               `json:"swimlane_id"`
	ColumnID      uint               `json:"column_id"`
	Column        Column             `gorm:"foreignKey:ColumnID" json:"column"`
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	Status        string             `json:"status"`
	Comments      []Comment          `gorm:"foreignKey:TaskID" json:"comments"`
	Position      float64            `json:"position"`
	CreatorID     uint               `json:"creator_id"`
	Creator       User               `gorm:"foreignKey:CreatorID" json:"creator"`
	AssigneeID    uint               `json:"assignee_id"`
	Assignee      User               `gorm:"foreignKey:AssigneeID" json:"assignee"`
	Files         []File             `gorm:"foreignKey:TaskID" json:"files"`
	DstLinks      []TaskLinks        `gorm:"foreignKey:DstTaskID" json:"dst_links"`
	SrcLinks      []TaskLinks        `gorm:"foreignKey:SrcTaskID" json:"src_links"`
	ExternalLinks []TaskExternalLink `gorm:"foreignKey:TaskID" json:"external_links"`
	DueDate       *time.Time         `json:"due_date"`
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
	User      User       `gorm:"foreignKey:UserID" json:"user"`
	UserID    uint       `json:"user_id"`
	Task      Task       `gorm:"foreignKey:TaskID" json:"task"`
	TaskID    uint       `json:"task_id"`
	Text      string     `json:"text"`
	Reactions []Reaction `gorm:"foreignKey:CommentID" json:"reactions"`
}

type Reaction struct {
	Model
	CommentID uint    `json:"comment_id" gorm:"not null; uniqueIndex:idx_reactions"`
	Comment   Comment `gorm:"foreignKey:CommentID" json:"comment"`
	User      User    `gorm:"foreignKey:UserID" json:"user"`
	UserID    uint    `json:"user_id" gorm:"not null; uniqueIndex:idx_reactions"`
	Reaction  string  `json:"reaction" gorm:"not null; uniqueIndex:idx_reactions"`
}

type File struct {
	Model
	Name           string `json:"name"`
	Task           Task   `gorm:"foreignKey:TaskID" json:"task"`
	TaskID         uint   `json:"task_id"`
	Path           string `json:"-"`
	Type           string `json:"type"` // image, file, etc
	UploadedBy     uint   `json:"uploaded_by"`
	UploadedByUser User   `gorm:"foreignKey:UploadedBy" json:"uploaded_by_user"`
}

type TaskLinks struct {
	Model
	SrcTaskID uint   `json:"src_task_id" gorm:"not null; uniqueIndex:idx_task_links"`
	SrcTask   Task   `gorm:"foreignKey:SrcTaskID" json:"src_task"`
	DstTaskID uint   `json:"dst_task_id" gorm:"not null; uniqueIndex:idx_task_links"`
	DstTask   Task   `gorm:"foreignKey:DstTaskID" json:"dst_task"`
	LinkType  string `json:"link_type" gorm:"not null; uniqueIndex:idx_task_links"`
}

type TaskExternalLink struct {
	Model
	URL    string `json:"url"`
	Title  string `json:"title"`
	Task   Task   `gorm:"foreignKey:TaskID" json:"task"`
	TaskID uint   `json:"task_id"`
}
