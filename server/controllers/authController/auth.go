package authController

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration details"
// @Success 201 {object} map[string]string "message: user created"
// @Failure 400 {object} models.ErrorResponse "error: bad request"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/register [post]
func Register(c *gin.Context) {
	var input RegisterRequest
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User login details"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} models.ErrorResponse "error: bad request"
// @Failure 401 {object} models.ErrorResponse "error: invalid credentials"
// @Router /api/v1/auth/login [post]
func Login(c *gin.Context) {
	var input LoginRequest
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
		if err := createTOTPSession(c, user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		resp := LoginResponse{
			Message: "totp_required",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := createLoginSession(c, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	resp := LoginResponse{
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
func Logout(c *gin.Context) {
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

// Profile godoc
// @Summary Get user profile
// @Description Get the profile of the logged-in user
// @Tags auth
// @Security cookieAuth
// @Produce json
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/profile [get]
func Profile(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")

	user, err := database.DB.UserRepository.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}

type ChangePasswordResponse struct {
	Message string `json:"message"`
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change the password of the logged-in user
// @Tags auth
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param passwordChange body ChangePasswordRequest true "Password change details"
// @Success 200 {object} ChangePasswordResponse "message: password changed successfully"
// @Failure 400 {object} models.ErrorResponse "error: bad request"
// @Failure 401 {object} models.ErrorResponse "error: unauthorized"
// @Failure 500 {object} models.ErrorResponse "error: internal server error"
// @Router /api/v1/auth/change-password [post]
func ChangePassword(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "unauthorized"})
		return
	}

	var input ChangePasswordRequest
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

	c.JSON(http.StatusOK, ChangePasswordResponse{
		Message: "password changed successfully",
	})
}

type GetCSRFTokenResponse struct {
	CSRFToken string `json:"csrf_token"`
}

// GetCSRFToken godoc
// @Summary Get CSRF token
// @Description Get the CSRF token for the logged-in user
// @Tags auth
// @Security cookieAuth
// @Produce json
// @Success 200 {object} GetCSRFTokenResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/auth/csrf-token [get]
func GetCSRFToken(c *gin.Context) {
	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "unauthorized"})
		return
	}

	if user.CSRFToken == "" {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, GetCSRFTokenResponse{
		CSRFToken: user.CSRFToken,
	})
}
