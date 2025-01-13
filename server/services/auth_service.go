package services

import (
	"errors"
	"server/config"
	"server/database/repository"
	e "server/internal/email"
	"server/internal/helpers"
	"server/models"
	"time"

	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	config *config.Config
	email  *e.EmailService
	db     *repository.Database
	helper *helpers.HelperService
	rs     *RoleService
}

func NewAuthService(config *config.Config, email *e.EmailService, db *repository.Database, helper *helpers.HelperService, rs *RoleService) *AuthService {
	return &AuthService{
		config: config,
		email:  email,
		db:     db,
		helper: helper,
		rs:     rs,
	}
}

func (s *AuthService) Register(username, email, password string) error {
	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     models.RoleUser,
	}

	// If this is the first user to register, set the role to admin
	count, err := s.db.UserRepository.Count()
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
	if err := s.db.UserRepository.Create(&user); err != nil {
		return err
	}

	pendingInvites, err := s.db.BoardInviteRepository.GetAll(repository.WithWhere("email = ?", email))
	if err != nil {
		return err
	}

	for _, invite := range pendingInvites {
		err := s.rs.AssignRole(user.ID, invite.BoardID, AppRole{Name: invite.RoleName})
		if err != nil {
			return err
		}
		if err := s.db.BoardInviteRepository.Delete(invite.ID); err != nil {
			return err
		}
	}

	return nil
}

func (s *AuthService) Login(email, password string) (uint, error) {
	user, err := s.db.UserRepository.GetFirst(repository.WithWhere("email = ?", email))
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
	user, err := s.db.UserRepository.GetByID(userID)
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
	if err := s.db.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ChangeUsername(userID uint, username string) error {
	user, err := s.db.UserRepository.GetByID(userID)
	if err != nil {
		return err
	}

	user.Username = username
	if err := s.db.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) RequestPasswordReset(email string) error {
	token := s.helper.GenerateRandomToken()

	// find user by email
	user, err := s.db.UserRepository.GetFirst(repository.WithWhere("email = ?", email))
	if err != nil {
		return err
	}

	// update user with reset token
	user.PasswordResetToken = token
	user.PasswordResetSentAt = time.Now()
	if err := s.db.UserRepository.Update(&user); err != nil {
		return err
	}

	// send email with code to reset password
	err = s.email.SendHTMLTemplate(user.Email, "Password Reset", "passwordReset.tmpl", map[string]string{
		"code":      token,
		"appUrl":    s.config.AppUrl,
		"appName":   s.config.AppName,
		"userEmail": user.Email,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ResetPassword(email, token, newPassword string) error {
	// find user by email
	user, err := s.db.UserRepository.GetFirst(repository.WithWhere("email = ?", email))
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

	// reset password reset token
	user.PasswordResetToken = ""
	if err := s.db.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GenerateTOTP(userID uint) (string, error) {
	user, err := s.db.UserRepository.GetByID(userID)
	if err != nil {
		return "", err
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      s.config.AppName,
		AccountName: user.Email,
	})
	if err != nil {
		return "", err
	}

	user.TotpSecret = key.Secret()
	if err := s.db.UserRepository.Update(&user); err != nil {
		return "", err
	}

	return key.Secret(), nil
}

func (s *AuthService) EnableTOTP(userID uint, code string) error {
	user, err := s.db.UserRepository.GetByID(userID)
	if err != nil {
		return err
	}

	valid := totp.Validate(code, user.TotpSecret)
	if !valid {
		return errors.New("invalid TOTP code")
	}

	user.TotpEnabled = true
	if err := s.db.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) DisableTOTP(userID uint, code string) error {
	user, err := s.db.UserRepository.GetByID(userID)
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
	if err := s.db.UserRepository.Update(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ConfirmTOTP(userID uint, code string) error {
	// find the user by ID
	user, err := s.db.UserRepository.GetByID(userID)
	if err != nil {
		return err
	}

	valid := totp.Validate(code, user.TotpSecret)
	if !valid {
		return errors.New("invalid TOTP code")
	}

	return nil
}
