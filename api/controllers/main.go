package controllers

import (
	"dogego/api/serializer"
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(200, serializer.Response{
		Code:    serializer.CodeOK,
		Message: "Pong",
	}.Result())
}
