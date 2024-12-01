package models

import "time"

type User struct {
	Model
	Email               string    `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Password            string    `gorm:"not null" json:"-" binding:"required,min=6"`
	Role                Role      `gorm:"default:user" json:"role"`
	TotpSecret          string    `json:"-"`
	TotpEnabled         bool      `gorm:"default:false" json:"totp_enabled"`
	PasswordResetToken  string    `json:"-"`
	PasswordResetSentAt time.Time `json:"-"`
}
