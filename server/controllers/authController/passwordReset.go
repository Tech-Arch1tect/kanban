package authController

import (
	"fmt"
	"log"
	"server/database"
	"server/email"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type PasswordResetResponse struct {
	Message string `json:"message"`
}

// PasswordReset godoc
// @Summary Request a password reset
// @Description Request a password reset for a user
// @Tags auth
// @Accept json
// @Produce json
// @Param passwordReset body PasswordResetRequest true "Password reset details"
// @Success 200 {object} PasswordResetResponse "message: If you have an account with us, you will receive a password reset link shortly."
// @Failure 400 {object} models.ErrorResponse "error: invalid request"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/password-reset [post]
func PasswordReset(c *gin.Context) {
	token := generateRandomToken()

	var req PasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// find user by email
	user, err := database.DB.GetUserByEmail(req.Email)
	if err != nil {
		// respond with 200 and message
		c.JSON(200, PasswordResetResponse{Message: "If you have an account with us, you will receive a password reset link shortly."})
		return
	}

	// update user with reset token
	user.PasswordResetToken = token
	user.PasswordResetSentAt = time.Now()
	if err := database.DB.UpdateUserByID(fmt.Sprint(user.ID), user); err != nil {
		c.JSON(500, PasswordResetResponse{Message: "Internal server error"})
		return
	}

	// send email with code to reset password
	err = email.SendPlainText(user.Email, "Password Reset", "A password reset request has been received for your account. Please use the following code to reset your password: "+token)
	if err != nil {
		log.Println("Error sending password reset email:", err)
		c.JSON(500, PasswordResetResponse{Message: "Failed to send password reset email. Please try again."})
		return
	}

	c.JSON(200, PasswordResetResponse{Message: "If you have an account with us, you will receive a password reset link shortly."})
}

type ResetPasswordRequest struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}

// ResetPassword godoc
// @Summary Reset a user's password
// @Description Reset a user's password
// @Tags auth
// @Accept json
// @Produce json
// @Param resetPassword body ResetPasswordRequest true "Reset password details"
// @Success 200 {object} ResetPasswordResponse "message: Password reset successful"
// @Failure 400 {object} models.ErrorResponse "error: invalid request"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/reset-password [post]
func ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// find user by email
	user, err := database.DB.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(500, PasswordResetResponse{Message: "Invalid code"})
		return
	}

	// verify code
	if user.PasswordResetToken != req.Code {
		c.JSON(500, PasswordResetResponse{Message: "Invalid code"})
		return
	}

	// check if code has expired
	if time.Since(user.PasswordResetSentAt) > 60*time.Minute {
		c.JSON(500, PasswordResetResponse{Message: "Invalid code"})
		return
	}

	// update user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, PasswordResetResponse{Message: "Internal server error"})
		return
	}
	user.Password = string(hashedPassword)
	if err := database.DB.UpdateUserByID(fmt.Sprint(user.ID), user); err != nil {
		c.JSON(500, PasswordResetResponse{Message: "Internal server error"})
		return
	}

	c.JSON(200, ResetPasswordResponse{Message: "Password reset successful"})
}
