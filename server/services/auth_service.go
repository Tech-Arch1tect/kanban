package services

import (
	"errors"
	"server/config"
	"server/database"
	e "server/internal/email"
	"server/internal/helpers"
	"server/models"
	"time"

	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) Register(email, password string) error {
	user := models.User{
		Email:    email,
		Password: password,
		Role:     models.RoleUser,
	}

	// If this is the first user to register, set the role to admin
	count, err := database.DB.UserRepository.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		user.Role = models.RoleAdmin
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	if err := database.DB.UserRepository.Create(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(email, password string) (uint, error) {
	user, err := database.DB.UserRepository.GetByEmail(email)
	if err != nil {
		return user.ID, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user.ID, err
	}

	if user.Role == models.RoleDisabled {
		return user.ID, errors.New("user is disabled")
	}

	if user.TotpEnabled {
		return user.ID, errors.New("totp_required")
	}

	return user.ID, nil
}

func (s *AuthService) ChangePassword(userID uint, currentPassword, newPassword string) error {
	user, err := database.DB.UserRepository.GetByID(userID)
	if err != nil {
		return err
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword)); err != nil {
		return err
	}

	// Hash and set new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	if err := database.DB.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) RequestPasswordReset(email string) error {
	token := helpers.GenerateRandomToken()

	// find user by email
	user, err := database.DB.UserRepository.GetByEmail(email)
	if err != nil {
		return err
	}

	// update user with reset token
	user.PasswordResetToken = token
	user.PasswordResetSentAt = time.Now()
	if err := database.DB.UserRepository.Update(&user); err != nil {
		return err
	}

	// send email with code to reset password
	err = e.SendPlainText(user.Email, "Password Reset", "A password reset request has been received for your account. Please use the following code to reset your password: "+token)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ResetPassword(email, token, newPassword string) error {
	// find user by email
	user, err := database.DB.UserRepository.GetByEmail(email)
	if err != nil {
		return err
	}

	// verify code
	if user.PasswordResetToken != token {
		return errors.New("invalid code")
	}

	// check if code has expired
	if time.Since(user.PasswordResetSentAt) > 60*time.Minute {
		return errors.New("code expired")
	}

	// update user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := database.DB.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GenerateTOTP(userID uint) (string, error) {
	user, err := database.DB.UserRepository.GetByID(userID)
	if err != nil {
		return "", err
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      config.CFG.AppName,
		AccountName: user.Email,
	})
	if err != nil {
		return "", err
	}

	user.TotpSecret = key.Secret()
	if err := database.DB.UserRepository.Update(&user); err != nil {
		return "", err
	}

	return key.Secret(), nil
}

func (s *AuthService) EnableTOTP(userID uint, code string) error {
	user, err := database.DB.UserRepository.GetByID(userID)
	if err != nil {
		return err
	}

	valid := totp.Validate(code, user.TotpSecret)
	if !valid {
		return errors.New("invalid TOTP code")
	}

	user.TotpEnabled = true
	if err := database.DB.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) DisableTOTP(userID uint, code string) error {
	user, err := database.DB.UserRepository.GetByID(userID)
	if err != nil {
		return err
	}

	if !user.TotpEnabled {
		return errors.New("TOTP is not enabled")
	}

	valid := totp.Validate(code, user.TotpSecret)
	if !valid {
		return errors.New("invalid TOTP code")
	}

	user.TotpEnabled = false
	user.TotpSecret = ""
	if err := database.DB.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ConfirmTOTP(userID uint, code string) error {
	// find the user by ID
	user, err := database.DB.UserRepository.GetByID(userID)
	if err != nil {
		return err
	}

	valid := totp.Validate(code, user.TotpSecret)
	if !valid {
		return errors.New("invalid TOTP code")
	}

	return nil
}
