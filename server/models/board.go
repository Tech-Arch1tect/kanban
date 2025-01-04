package models

type Board struct {
	Model
	Name      string     `gorm:"not null" json:"name"`
	Swimlanes []Swimlane `gorm:"foreignKey:BoardID" json:"swimlanes"`
	Tasks     []Task     `gorm:"foreignKey:BoardID" json:"tasks"`
	Columns   []Column   `gorm:"foreignKey:BoardID" json:"columns"`
	Slug      string     `gorm:"not null;unique" json:"slug"`
}

type BoardPermission struct {
	Model
	Name string `gorm:"not null" json:"name" binding:"oneof=admin member viewer"`
}

type UserBoardPermission struct {
	Model
	UserID            uint            `gorm:"not null" json:"user_id"`
	User              User            `gorm:"foreignKey:UserID" json:"user"`
	BoardID           uint            `gorm:"not null" json:"board_id"`
	Board             Board           `gorm:"foreignKey:BoardID" json:"board"`
	BoardPermissionID uint            `gorm:"not null" json:"board_permission_id"`
	BoardPermission   BoardPermission `gorm:"foreignKey:BoardPermissionID" json:"board_permission"`
}

type Swimlane struct {
	Model
	BoardID uint   `json:"board_id"`
	Board   Board  `gorm:"foreignKey:BoardID" json:"board"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
}

type Column struct {
	Model
	Name    string `gorm:"not null" json:"name"`
	BoardID uint   `json:"board_id"`
	Order   int    `json:"order"`
}
