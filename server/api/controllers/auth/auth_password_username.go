package auth

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthChangePassword godoc
// @Summary Change user password
// @Description Change the password of the logged-in user
// @Tags auth
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param passwordChange body AuthChangePasswordRequest true "Password change details"
// @Success 200 {object} AuthChangePasswordResponse "message: password changed successfully"
// @Failure 400 {object} models.ErrorResponse "error: bad request"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/change-password [post]
func (a *AuthController) ChangePassword(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "unauthorized"})
		return
	}

	var input AuthChangePasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": a.helperService.ParseValidationError(err)})
		return
	}

	err := a.authService.ChangePassword(userID.(uint), input.CurrentPassword, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthChangePasswordResponse{
		Message: "password changed successfully",
	})
}

// AuthPasswordReset godoc
// @Summary Request a password reset
// @Description Request a password reset for a user
// @Tags auth
// @Accept json
// @Produce json
// @Param passwordReset body AuthPasswordResetRequest true "Password reset details"
// @Success 200 {object} AuthPasswordResetResponse "message: If you have an account with us, you will receive a password reset link shortly."
// @Failure 400 {object} models.ErrorResponse "error: invalid request"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/password-reset [post]
func (a *AuthController) PasswordReset(c *gin.Context) {

	var req AuthPasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := a.authService.RequestPasswordReset(req.Email)
	if err != nil {
		c.JSON(500, AuthPasswordResetResponse{Message: "Internal server error"})
		return
	}

	c.JSON(200, AuthPasswordResetResponse{Message: "If you have an account with us, you will receive a password reset link shortly."})
}

// ResetPassword godoc
// @Summary Reset a user's password
// @Description Reset a user's password
// @Tags auth
// @Accept json
// @Produce json
// @Param resetPassword body AuthResetPasswordRequest true "Reset password details"
// @Success 200 {object} AuthResetPasswordResponse "message: Password reset successful"
// @Failure 400 {object} models.ErrorResponse "error: invalid request"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/reset-password [post]
func (a *AuthController) ResetPassword(c *gin.Context) {
	var req AuthResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := a.authService.ResetPassword(req.Email, req.Code, req.Password)
	if err != nil {
		c.JSON(500, AuthResetPasswordResponse{Message: err.Error()})
		return
	}

	c.JSON(200, AuthResetPasswordResponse{Message: "Password reset successful"})
}

// AuthChangeUsername godoc
// @Summary Change user username
// @Description Change the username of the logged-in user
// @Tags auth
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param usernameChange body AuthChangeUsernameRequest true "Username change details"
// @Success 200 {object} AuthChangeUsernameResponse "message: username changed successfully"
// @Failure 400 {object} models.ErrorResponse "error: bad request"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/change-username [post]
func (a *AuthController) ChangeUsername(c *gin.Context) {
	user, err := a.helperService.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "unauthorized"})
		return
	}

	var input AuthChangeUsernameRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": a.helperService.ParseValidationError(err)})
		return
	}

	err = a.authService.ChangeUsername(user.ID, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthChangeUsernameResponse{Message: "username changed successfully"})
}
