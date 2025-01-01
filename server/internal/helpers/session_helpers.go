package helpers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateLoginSession(c *gin.Context, userID uint) error {
	if err := ClearLoginSession(c); err != nil {
		return err
	}
	session := sessions.Default(c)
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func ClearLoginSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func CreateTOTPSession(c *gin.Context, userID uint) error {
	session := sessions.Default(c)
	session.Set("totpUserID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func ClearTOTPSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Delete("totpUserID")
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func GenerateRandomToken() string {
	return uuid.New().String()
}
