package authController

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createLoginSession(c *gin.Context, userID uint) {
	clearLoginSession(c)
	session := sessions.Default(c)
	session.Set("userID", userID)
	session.Save()
}

func clearLoginSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func createTOTPSession(c *gin.Context, userID uint) {
	session := sessions.Default(c)
	session.Set("totpUserID", userID)
	session.Save()
}

func clearTOTPSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("totpUserID")
	session.Save()
}

func generateRandomToken() string {
	return uuid.New().String()
}
