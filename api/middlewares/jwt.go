package middlewares

import (
	"{{cookiecutter.app_name}}/modules"
	jwtmiddle "github.com/doge-soft/dogego_module_jwt/middlewares"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	jwt := jwtmiddle.NewJwtMiddleware(modules.JWTModule)

	return jwt.New()
}
