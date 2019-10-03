package utils

import (
	"dogego/models"
	"github.com/gin-gonic/gin"
)

func CurrentUser(context *gin.Context) *models.User {
	if user, _ := context.Get("user"); user != nil {
		if u, ok := user.(*models.User); ok {
			return u
		}
	}
	return nil
}
