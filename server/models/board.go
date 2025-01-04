package models

type Board struct {
	Model
	Name      string     `gorm:"not null" json:"name"`
	Swimlanes []Swimlane `gorm:"foreignKey:BoardID" json:"swimlanes"`
	Tasks     []Task     `gorm:"foreignKey:BoardID" json:"tasks"`
	Columns   []Column   `gorm:"foreignKey:BoardID" json:"columns"`
	Slug      string     `gorm:"not null;unique" json:"slug"`
}

type BoardRole struct {
	Model
	Name string `gorm:"not null" json:"name" binding:"oneof=admin member viewer"`
}

type UserBoardRole struct {
	Model
	UserID      uint      `gorm:"not null;uniqueIndex:idx_user_board_role" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	BoardID     uint      `gorm:"not null;uniqueIndex:idx_user_board_role" json:"board_id"`
	Board       Board     `gorm:"foreignKey:BoardID" json:"board"`
	BoardRoleID uint      `gorm:"not null;" json:"board_role_id"`
	BoardRole   BoardRole `gorm:"foreignKey:BoardRoleID" json:"role"`
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
