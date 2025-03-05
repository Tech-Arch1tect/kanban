package middleware

import (
	"net/http"
	"server/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m *Middleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}
		c.Next()
	}
}

func (m *Middleware) EnsureRole(role models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}
		user, err := m.db.UserRepository.GetByID(userID.(uint))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		if user.Role != role {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}
		c.Next()
	}
}

func (m *Middleware) TOTPTempAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("totpUserID")
		if userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}
	}
}

func (m *Middleware) EnsureCSRFTokenExistsInSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		csrfToken := session.Get("csrfToken")
		if csrfToken == nil {
			if err := m.updateCSRFToken(c, m.generateCSRFToken()); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		}
		c.Next()
	}
}

func (m *Middleware) CSRFTokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-CSRF-Token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		if !m.validateCSRFToken(c, token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		if err := m.updateCSRFToken(c, m.generateCSRFToken()); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.Next()
	}
}

func (m *Middleware) generateCSRFToken() string {
	return uuid.New().String()
}

func (m *Middleware) validateCSRFToken(c *gin.Context, token string) bool {
	session := sessions.Default(c)
	csrfToken := session.Get("csrfToken")
	return csrfToken == token
}

func (m *Middleware) updateCSRFToken(c *gin.Context, token string) error {
	session := sessions.Default(c)
	session.Set("csrfToken", token)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
