package middlewares

import (
	"dogego/api/serializer"
	"dogego/dao"
	"dogego/models"
	jwtmodels "github.com/doge-soft/dogego_module_jwt/models"
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		claims, ok := context.Get("claims")

		if ok {
			if claim, ok := claims.(*jwtmodels.UserClaim); ok {
				user, err := dao.GetUserById(claim.UserId)

				if err == nil {
					context.Set("user", user)
				}
			}
		}

		context.Next()
	}
}

func AuthRequired(role []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if user, _ := context.Get("user"); user != nil {
			if v, ok := user.(*models.User); ok {
				for _, r := range role {
					if v.Role == r {
						context.Next()
						return
					}
				}
			}
		}

		context.JSON(200, serializer.ErrorResponse(
			serializer.CodeAuthorizationError,
			"您没有登陆或者没有权限.",
			nil,
		).Result())

		context.Abort()
	}
}
