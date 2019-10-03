package services

import (
	"dogego/api/serializer"
	"dogego/dao"
	"dogego/modules"
	"github.com/doge-soft/dogego_module_jwt/models"
)

type UserLoginService struct {
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"required,min=11,max=21"`
	Password    string `form:"password" json:"password" binding:"required,min=3,max=45"`
}

func (service *UserLoginService) Login() *serializer.Response {
	user, err, isLogin := dao.LoginUser(service.PhoneNumber, service.Password)

	if err != nil {
		return serializer.DatabaseErrorResponse(
			"数据库操作出错.",
			err,
		).Result()
	}

	if !isLogin {
		return serializer.Response{
			Code:    serializer.CodeAuthorizationError,
			Message: "用户名或密码错误.",
		}.Result()
	}

	token, err := modules.JWTModule.GenerateToken(&models.UserClaim{
		UserId: user.ID,
	})

	if err != nil {
		return serializer.ErrorResponse(
			serializer.JWTGenerateError,
			"JWT Token 生成失败",
			err,
		).Result()
	}

	return serializer.Response{
		Code:    serializer.CodeOK,
		Message: "登陆成功.",
		Data:    token,
	}.Result()
}
