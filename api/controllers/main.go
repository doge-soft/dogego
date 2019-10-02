package controllers

import (
	"dogego/api/services"
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	service := services.PingService{}

	context.JSON(200, service.Ping().Result())
}
