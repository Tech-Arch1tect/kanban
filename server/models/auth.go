package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Model
	Email               string    `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Password            string    `gorm:"not null" json:"-" binding:"required,min=6"`
	Role                Role      `gorm:"default:user" json:"role"`
	TotpSecret          string    `json:"-"`
	TotpEnabled         bool      `gorm:"default:false" json:"totp_enabled"`
	PasswordResetToken  string    `json:"-"`
	PasswordResetSentAt time.Time `json:"-"`
	CSRFToken           string    `json:"-" gorm:"default:''"`
}

func (u *User) GenerateCSRFToken() {
	u.CSRFToken = uuid.New().String()
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.GenerateCSRFToken()
	return
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.GenerateCSRFToken()
	return
}
