package controllers

import (
	"{{cookiecutter.app_name}}/api/services"
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	service := services.PingService{}

	context.JSON(200, service.Ping().Result())
}
