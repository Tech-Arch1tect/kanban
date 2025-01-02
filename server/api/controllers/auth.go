package controllers

import (
	"net/http"
	"server/database"
	"server/internal/helpers"
	"server/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body AuthRegisterRequest true "User registration details"
// @Success 201 {object} map[string]string "message: user created"
// @Failure 400 {object} models.ErrorResponse "error: bad request"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/register [post]
func AuthRegister(c *gin.Context) {
	var input AuthRegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.ParseValidationError(err)})
		return
	}

	err := authService.Register(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body AuthLoginRequest true "User login details"
// @Success 200 {object} AuthLoginResponse
// @Failure 400 {object} models.ErrorResponse "error: bad request"
// @Failure 401 {object} models.ErrorResponse "error: invalid credentials"
// @Router /api/v1/auth/login [post]
func AuthLogin(c *gin.Context) {
	var input AuthLoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.ParseValidationError(err)})
		return
	}

	id, err := authService.Login(input.Email, input.Password)
	if err != nil && err.Error() == "totp_required" {
		if err := helpers.CreateTOTPSession(c, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		resp := AuthLoginResponse{
			Message: "totp_required",
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := helpers.CreateLoginSession(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	resp := AuthLoginResponse{
		Message: "logged in",
	}
	c.JSON(http.StatusOK, resp)
}

// Logout godoc
// @Summary Logout a user
// @Description Logout the current user
// @Tags auth
// @Security csrf
// @Produce json
// @Success 200 {object} map[string]string "message: logged out"
// @Router /api/v1/auth/logout [post]
func AuthLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

// AuthProfile godoc
// @Summary Get user profile
// @Description Get the profile of the logged-in user
// @Tags auth
// @Security cookieAuth
// @Produce json
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/profile [get]
func AuthProfile(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")

	user, err := database.DB.UserRepository.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

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
func AuthChangePassword(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "unauthorized"})
		return
	}

	var input AuthChangePasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.ParseValidationError(err)})
		return
	}

	err := authService.ChangePassword(userID.(uint), input.CurrentPassword, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthChangePasswordResponse{
		Message: "password changed successfully",
	})
}

// AuthGetCSRFToken godoc
// @Summary Get CSRF token
// @Description Get the CSRF token for the logged-in user
// @Tags auth
// @Security cookieAuth
// @Produce json
// @Success 200 {object} AuthGetCSRFTokenResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/auth/csrf-token [get]
func AuthGetCSRFToken(c *gin.Context) {
	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "unauthorized"})
		return
	}

	if user.CSRFToken == "" {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, AuthGetCSRFTokenResponse{
		CSRFToken: user.CSRFToken,
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
func AuthPasswordReset(c *gin.Context) {

	var req AuthPasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := authService.RequestPasswordReset(req.Email)
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
func AuthResetPassword(c *gin.Context) {
	var req AuthResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := authService.ResetPassword(req.Email, req.Code, req.Password)
	if err != nil {
		c.JSON(500, AuthResetPasswordResponse{Message: err.Error()})
		return
	}

	c.JSON(200, AuthResetPasswordResponse{Message: "Password reset successful"})
}

// TOTP related routes using github.com/pquerna/otp/totp

// GenerateTOTP godoc
// @Summary Generate TOTP secret
// @Description Generate a new TOTP secret for the logged-in user
// @Tags auth
// @Security cookieAuth
// @Security csrf
// @Produce json
// @Success 200 {object} AuthGenerateTOTPResponse
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/totp/generate [post]
func AuthGenerateTOTP(c *gin.Context) {
	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	secret, err := authService.GenerateTOTP(user.ID)
	if err != nil || secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate TOTP secret"})
		return
	}

	response := AuthGenerateTOTPResponse{Secret: secret}
	c.JSON(http.StatusOK, response)
}

// EnableTOTP godoc
// @Summary Enable TOTP
// @Description Enable TOTP for the logged-in user
// @Tags auth
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body AuthEnableTOTPRequest true "TOTP code"
// @Success 200 {object} AuthEnableTOTPResponse
// @Failure 400 {object} models.ErrorResponse "error: invalid TOTP code"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/totp/enable [post]
func AuthEnableTOTP(c *gin.Context) {

	var req AuthEnableTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = authService.EnableTOTP(user.ID, req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthEnableTOTPResponse{Message: "TOTP enabled"})
}

// DisableTOTP godoc
// @Summary Disable TOTP
// @Description Disable TOTP for the logged-in user
// @Tags auth
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body AuthDisableTOTPRequest true "TOTP code"
// @Success 200 {object} AuthDisableTOTPResponse
// @Failure 400 {object} models.ErrorResponse "error: invalid TOTP code"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/totp/disable [post]
func AuthDisableTOTP(c *gin.Context) {

	var req AuthDisableTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = authService.DisableTOTP(user.ID, req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthDisableTOTPResponse{Message: "TOTP disabled"})
}

// ConfirmTOTP godoc
// @Summary Confirm TOTP code
// @Description Confirm TOTP code for the user during login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body AuthConfirmTOTPRequest true "TOTP code"
// @Success 200 {object} AuthConfirmTOTPResponse
// @Failure 400 {object} models.ErrorResponse "error: invalid request or invalid TOTP code"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/totp/confirm [post]
func AuthConfirmTOTP(c *gin.Context) {

	var req AuthConfirmTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = authService.ConfirmTOTP(user.ID, req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := helpers.CreateLoginSession(c, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if err := helpers.ClearTOTPSession(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, AuthConfirmTOTPResponse{Message: "totp_confirmed"})
}
