package helpers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *HelperService) CreateLoginSession(c *gin.Context, userID uint) error {
	if err := h.ClearLoginSession(c); err != nil {
		return err
	}
	session := sessions.Default(c)
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (h *HelperService) ClearLoginSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (h *HelperService) CreateTOTPSession(c *gin.Context, userID uint) error {
	session := sessions.Default(c)
	session.Set("totpUserID", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (h *HelperService) ClearTOTPSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Delete("totpUserID")
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (h *HelperService) GenerateRandomToken() string {
	return uuid.New().String()
}
