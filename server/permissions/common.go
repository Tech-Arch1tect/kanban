package permissions

import (
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func Can(c *gin.Context, permFunc func(user models.User, id uint) bool, id uint) (bool, models.User) {
	user, err := helpers.GetUserFromSession(c)
	if err != nil {
		return false, models.User{}
	}

	if user.Role == models.RoleAdmin {
		return true, user
	}

	return permFunc(user, id), user
}
