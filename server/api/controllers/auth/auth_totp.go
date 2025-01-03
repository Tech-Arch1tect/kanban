package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
func (a *AuthController) GenerateTOTP(c *gin.Context) {
	user, err := a.helperService.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	secret, err := a.authService.GenerateTOTP(user.ID)
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
func (a *AuthController) EnableTOTP(c *gin.Context) {

	var req AuthEnableTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := a.helperService.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = a.authService.EnableTOTP(user.ID, req.Code)
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
func (a *AuthController) DisableTOTP(c *gin.Context) {

	var req AuthDisableTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := a.helperService.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = a.authService.DisableTOTP(user.ID, req.Code)
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
func (a *AuthController) ConfirmTOTP(c *gin.Context) {

	var req AuthConfirmTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := a.helperService.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = a.authService.ConfirmTOTP(user.ID, req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := a.helperService.CreateLoginSession(c, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if err := a.helperService.ClearTOTPSession(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, AuthConfirmTOTPResponse{Message: "totp_confirmed"})
}
