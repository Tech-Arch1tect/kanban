package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Username            string    `gorm:"uniqueIndex;not null" json:"username" binding:"required"`
	Email               string    `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Password            string    `gorm:"not null" json:"-" binding:"required,min=6"`
	Role                Role      `gorm:"default:user" json:"role"`
	TotpSecret          string    `json:"-"`
	TotpEnabled         bool      `gorm:"default:false" json:"totp_enabled"`
	PasswordResetToken  string    `json:"-"`
	PasswordResetSentAt time.Time `json:"-"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
