package controllers

import (
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(200, "OK!")
}
