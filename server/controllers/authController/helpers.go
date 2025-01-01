package authController

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createLoginSession(c *gin.Context, userID uint) error {
	if err := clearLoginSession(c); err != nil {
		return err
	}
	session := sessions.Default(c)
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func clearLoginSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func createTOTPSession(c *gin.Context, userID uint) error {
	session := sessions.Default(c)
	session.Set("totpUserID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func clearTOTPSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Delete("totpUserID")
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func generateRandomToken() string {
	return uuid.New().String()
}
