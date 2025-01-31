package auth

import (
	"net/http"
	"server/database/repository"
	"server/internal/helpers"
	"server/models"
	"server/services/auth"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService   *auth.AuthService
	db            *repository.Database
	helperService *helpers.HelperService
}

func NewAuthController(authService *auth.AuthService, db *repository.Database, helperService *helpers.HelperService) *AuthController {
	return &AuthController{authService: authService, db: db, helperService: helperService}
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
func (a *AuthController) Register(c *gin.Context) {
	var input AuthRegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": a.helperService.ParseValidationError(err)})
		return
	}

	err := a.authService.Register(input.Username, input.Email, input.Password)
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
func (a *AuthController) Login(c *gin.Context) {
	var input AuthLoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": a.helperService.ParseValidationError(err)})
		return
	}

	id, err := a.authService.Login(input.Email, input.Password)
	if err != nil && err.Error() == "totp_required" {
		if err := a.helperService.CreateTOTPSession(c, id); err != nil {
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

	if err := a.helperService.CreateLoginSession(c, id); err != nil {
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
func (a *AuthController) Logout(c *gin.Context) {
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
func (a *AuthController) Profile(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")

	user, err := a.db.UserRepository.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
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
func (a *AuthController) GetCSRFToken(c *gin.Context) {
	session := sessions.Default(c)
	csrfToken := session.Get("csrfToken")

	if csrfToken == nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, AuthGetCSRFTokenResponse{
		CSRFToken: csrfToken.(string),
	})
}
