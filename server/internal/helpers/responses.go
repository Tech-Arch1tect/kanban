package helpers

import (
	"errors"
	"fmt"
	"server/database"
	"server/models"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ParseValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", e.Field()))
		}
		return strings.Join(errorMessages, ", ")
	}
	return "bad request"
}

func GetUserFromSession(c *gin.Context) (models.User, error) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		return models.User{}, errors.New("unauthenticated")
	}
	user, err := database.DB.UserRepository.GetByID(userID.(uint))
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
