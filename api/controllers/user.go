package controllers

import (
	"dogego/api/serializer"
	"dogego/api/services"
	"dogego/modules"
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

func UserLogout(context *gin.Context) {
	token := context.GetHeader("Authorization")

	err := modules.JWTModule.DieToken(token)

	if err != nil {
		context.JSON(200, serializer.ErrorResponse(
			serializer.JWTLogoutError,
			"登出失败.",
			err,
		).Result())
		return
	}

	context.JSON(200, serializer.Response{
		Code:    serializer.CodeOK,
		Message: "登出成功.",
	}.Result())
}
