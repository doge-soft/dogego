package controllers

import (
	"dogego/api/services"
	"dogego/utils"
	"github.com/gin-gonic/gin"
)

func UserRegister(context *gin.Context) {
	service := services.UserRegisterService{}

	if err := context.ShouldBind(&service); err == nil {
		context.JSON(200, service.Register())
	} else {
		context.JSON(200, utils.ParamaterErrorResponse(err))
	}
}
