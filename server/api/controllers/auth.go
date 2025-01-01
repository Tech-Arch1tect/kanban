package controllers

import (
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/internal/email"
	"server/internal/helpers"
	"server/models"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"

	"golang.org/x/crypto/bcrypt"
)

type AuthRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
		Role:     models.RoleUser,
	}

	// If this is the first user to register, set the role to admin
	count, err := database.DB.UserRepository.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if count == 0 {
		user.Role = models.RoleAdmin
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	user.Password = string(hashedPassword)
	if err := database.DB.UserRepository.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Message string `json:"message"`
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

	user, err := database.DB.UserRepository.GetByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "invalid credentials",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "invalid credentials",
		})
		return
	}

	if user.Role == models.RoleDisabled {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "user is disabled",
		})
		return
	}

	if user.TotpEnabled {
		if err := helpers.CreateTOTPSession(c, user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		resp := AuthLoginResponse{
			Message: "totp_required",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := helpers.CreateLoginSession(c, user.ID); err != nil {
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
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user found"})
		return
	}
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

type AuthChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}

type AuthChangePasswordResponse struct {
	Message string `json:"message"`
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

	user, err := database.DB.UserRepository.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal server error"})
		return
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.CurrentPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "invalid current password"})
		return
	}

	// Hash and set new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal server error"})
		return
	}

	user.Password = string(hashedPassword)
	if err := database.DB.UserRepository.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, AuthChangePasswordResponse{
		Message: "password changed successfully",
	})
}

type AuthGetCSRFTokenResponse struct {
	CSRFToken string `json:"csrf_token"`
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

type AuthPasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type AuthPasswordResetResponse struct {
	Message string `json:"message"`
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
	token := helpers.GenerateRandomToken()

	var req AuthPasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// find user by email
	user, err := database.DB.UserRepository.GetByEmail(req.Email)
	if err != nil {
		// respond with 200 and message
		c.JSON(200, AuthPasswordResetResponse{Message: "If you have an account with us, you will receive a password reset link shortly."})
		return
	}

	// update user with reset token
	user.PasswordResetToken = token
	user.PasswordResetSentAt = time.Now()
	if err := database.DB.UserRepository.Update(&user); err != nil {
		c.JSON(500, AuthPasswordResetResponse{Message: "Internal server error"})
		return
	}

	// send email with code to reset password
	err = email.SendPlainText(user.Email, "Password Reset", "A password reset request has been received for your account. Please use the following code to reset your password: "+token)
	if err != nil {
		log.Println("Error sending password reset email:", err)
		c.JSON(500, AuthPasswordResetResponse{Message: "Failed to send password reset email. Please try again."})
		return
	}

	c.JSON(200, AuthPasswordResetResponse{Message: "If you have an account with us, you will receive a password reset link shortly."})
}

type AuthResetPasswordRequest struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type AuthResetPasswordResponse struct {
	Message string `json:"message"`
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

	// find user by email
	user, err := database.DB.UserRepository.GetByEmail(req.Email)
	if err != nil {
		c.JSON(500, AuthResetPasswordResponse{Message: "Invalid code"})
		return
	}

	// verify code
	if user.PasswordResetToken != req.Code {
		c.JSON(500, AuthResetPasswordResponse{Message: "Invalid code"})
		return
	}

	// check if code has expired
	if time.Since(user.PasswordResetSentAt) > 60*time.Minute {
		c.JSON(500, AuthResetPasswordResponse{Message: "Invalid code"})
		return
	}

	// update user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, AuthResetPasswordResponse{Message: "Internal server error"})
		return
	}
	user.Password = string(hashedPassword)
	if err := database.DB.UserRepository.Update(&user); err != nil {
		c.JSON(500, AuthResetPasswordResponse{Message: "Internal server error"})
		return
	}

	c.JSON(200, AuthResetPasswordResponse{Message: "Password reset successful"})
}

// TOTP related routes using github.com/pquerna/otp/totp

type AuthGenerateTOTPResponse struct {
	Secret string `json:"secret"`
}

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

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      config.CFG.AppName,
		AccountName: user.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate TOTP secret"})
		return
	}

	user.TotpSecret = key.Secret()
	if err := database.DB.UserRepository.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save TOTP secret"})
		return
	}

	response := AuthGenerateTOTPResponse{Secret: key.Secret()}
	c.JSON(http.StatusOK, response)
}

type AuthEnableTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type AuthEnableTOTPResponse struct {
	Message string `json:"message"`
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
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := database.DB.UserRepository.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	var req AuthEnableTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	valid := totp.Validate(req.Code, user.TotpSecret)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid TOTP code"})
		return
	}

	user.TotpEnabled = true
	if err := database.DB.UserRepository.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enable TOTP"})
		return
	}

	response := AuthEnableTOTPResponse{Message: "TOTP enabled"}
	c.JSON(http.StatusOK, response)
}

type AuthDisableTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type AuthDisableTOTPResponse struct {
	Message string `json:"message"`
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
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := database.DB.UserRepository.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if !user.TotpEnabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TOTP is not enabled"})
		return
	}

	var req AuthDisableTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	valid := totp.Validate(req.Code, user.TotpSecret)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid TOTP code"})
		return
	}

	user.TotpEnabled = false
	user.TotpSecret = ""
	if err := database.DB.UserRepository.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to disable TOTP"})
		return
	}

	response := AuthDisableTOTPResponse{Message: "TOTP disabled"}
	c.JSON(http.StatusOK, response)
}

type AuthConfirmTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type AuthConfirmTOTPResponse struct {
	Message string `json:"message"`
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
	session := sessions.Default(c)
	userID := session.Get("totpUserID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// find the user by ID
	user, err := database.DB.UserRepository.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// verify the TOTP code
	var req AuthConfirmTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	valid := totp.Validate(req.Code, user.TotpSecret)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid TOTP code"})
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
	response := AuthConfirmTOTPResponse{Message: "totp_confirmed"}
	c.JSON(http.StatusOK, response)
}

