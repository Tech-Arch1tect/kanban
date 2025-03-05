package helpers

import (
	"errors"
	"fmt"
	"server/database/repository"
	"server/models"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HelperService struct {
	db *repository.Database
}

func NewHelperService(db *repository.Database) *HelperService {
	return &HelperService{db: db}
}

func (h *HelperService) ParseValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", e.Field()))
		}
		return strings.Join(errorMessages, ", ")
	}
	return "bad request"
}

func (h *HelperService) GetUserFromSession(c *gin.Context) (models.User, error) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		return models.User{}, errors.New("unauthenticated")
	}
	user, err := h.db.UserRepository.GetByID(userID.(uint))
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
