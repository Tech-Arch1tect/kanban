package models

type Board struct {
	Model
	Name        string            `gorm:"not null" json:"name" binding:"required"`
	Permissions []BoardPermission `gorm:"foreignKey:BoardID" json:"permissions"`
	Owner       User              `gorm:"foreignKey:OwnerID" json:"owner"`
	OwnerID     uint              `gorm:"not null" json:"owner_id"`
	Swimlanes   []Swimlane        `gorm:"foreignKey:BoardID" json:"swimlanes"`
	Tasks       []Task            `gorm:"foreignKey:BoardID" json:"tasks"`
	Columns     []Column          `gorm:"foreignKey:BoardID" json:"columns"`
	Slug        string            `gorm:"not null;unique" json:"slug"`
}

type BoardPermission struct {
	Model
	BoardID       uint `json:"board_id"`
	User          User `gorm:"foreignKey:UserID" json:"user"`
	UserID        uint `json:"user_id"`
	Edit          bool `json:"edit"`
	Delete        bool `json:"delete"`
	GeneralAccess bool `json:"general_access"`
}
