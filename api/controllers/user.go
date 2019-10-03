package controllers

import (
	"dogego/api/serializer"
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

func UserLogin(context *gin.Context) {
	service := services.UserLoginService{}

	if err := context.ShouldBind(&service); err == nil {
		context.JSON(200, service.Login())
	} else {
		context.JSON(200, utils.ParamaterErrorResponse(err))
	}
}

func UserMe(context *gin.Context) {
	user := utils.CurrentUser(context)

	context.JSON(200, serializer.BuildUserResponse(*user))
}
