package authController

import (
	"net/http"
	"server/config"
	"server/database"
	"server/helpers"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

// TOTP related routes using github.com/pquerna/otp/totp

type GenerateTOTPResponse struct {
	Secret string `json:"secret"`
}

// GenerateTOTP godoc
// @Summary Generate TOTP secret
// @Description Generate a new TOTP secret for the logged-in user
// @Tags auth
// @Security cookieAuth
// @Produce json
// @Success 200 {object} GenerateTOTPResponse
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/totp/generate [post]
func GenerateTOTP(c *gin.Context) {
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

	response := GenerateTOTPResponse{Secret: key.Secret()}
	c.JSON(http.StatusOK, response)
}

type EnableTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type EnableTOTPResponse struct {
	Message string `json:"message"`
}

// EnableTOTP godoc
// @Summary Enable TOTP
// @Description Enable TOTP for the logged-in user
// @Tags auth
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param request body EnableTOTPRequest true "TOTP code"
// @Success 200 {object} EnableTOTPResponse
// @Failure 400 {object} models.ErrorResponse "error: invalid TOTP code"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/totp/enable [post]
func EnableTOTP(c *gin.Context) {
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

	var req EnableTOTPRequest
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

	response := EnableTOTPResponse{Message: "TOTP enabled"}
	c.JSON(http.StatusOK, response)
}

type DisableTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type DisableTOTPResponse struct {
	Message string `json:"message"`
}

// DisableTOTP godoc
// @Summary Disable TOTP
// @Description Disable TOTP for the logged-in user
// @Tags auth
// @Security cookieAuth
// @Accept json
// @Produce json
// @Param request body DisableTOTPRequest true "TOTP code"
// @Success 200 {object} DisableTOTPResponse
// @Failure 400 {object} models.ErrorResponse "error: invalid TOTP code"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/totp/disable [post]
func DisableTOTP(c *gin.Context) {
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

	var req DisableTOTPRequest
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

	response := DisableTOTPResponse{Message: "TOTP disabled"}
	c.JSON(http.StatusOK, response)
}

type ConfirmTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type ConfirmTOTPResponse struct {
	Message string `json:"message"`
}

// ConfirmTOTP godoc
// @Summary Confirm TOTP code
// @Description Confirm TOTP code for the user during login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ConfirmTOTPRequest true "TOTP code"
// @Success 200 {object} ConfirmTOTPResponse
// @Failure 400 {object} models.ErrorResponse "error: invalid request or invalid TOTP code"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/totp/confirm [post]
func ConfirmTOTP(c *gin.Context) {
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
	var req ConfirmTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	valid := totp.Validate(req.Code, user.TotpSecret)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid TOTP code"})
		return
	}

	createLoginSession(c, user.ID)
	clearTOTPSession(c)
	response := ConfirmTOTPResponse{Message: "totp_confirmed"}
	c.JSON(http.StatusOK, response)
}
