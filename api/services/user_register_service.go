package services

import (
	"{{cookiecutter.app_name}}/api/serializer"
	"{{cookiecutter.app_name}}/auth"
	"{{cookiecutter.app_name}}/dao"
	"{{cookiecutter.app_name}}/datasource"
	"{{cookiecutter.app_name}}/models"
	"fmt"
)

type UserRegisterService struct {
	NickName        string `form:"nick_name" json:"nick_name"`
	PhoneNumber     string `form:"phone_number" json:"phone_number" binding:"required,min=11,max=21"`
	Password        string `form:"password" json:"password" binding:"required,min=3,max=45"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=3,max=45"`
}

func (service *UserRegisterService) Vaild() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return serializer.ParamaterErrorResponse(
			"两次输入的密码不匹配.",
			nil,
		).Result()
	}

	count := 0
	datasource.SlaveDatabase().Model(&models.User{}).Where("nick_name = ?", service.NickName).Count(&count)
	if count > 0 {
		return serializer.ParamaterErrorResponse(
			"昵称已经被占用了.",
			nil,
		).Result()
	}

	count = 0
	datasource.SlaveDatabase().Model(&models.User{}).Where("phone_number = ?", service.PhoneNumber).Count(&count)
	if count > 0 {
		return serializer.ParamaterErrorResponse(
			"电话号码被占用.",
			nil,
		).Result()
	}

	return nil
}

func (service *UserRegisterService) Replace() *serializer.Response {
	if service.NickName == "" {
		service.NickName = fmt.Sprintf("用户 %s", service.PhoneNumber)
	}

	return nil
}

func (service *UserRegisterService) Register() *serializer.Response {
	if err := service.Vaild(); err != nil {
		return err
	}

	if err := service.Replace(); err != nil {
		return err
	}

	user := models.User{
		NickName:    service.NickName,
		PhoneNumber: service.PhoneNumber,
		Bio:         "这个人很懒, 什么都没写....",
		Avatar:      "null",
		Status:      models.Active,
		Role:        auth.User[0],
	}

	if err := user.SetPassword(service.Password); err != nil {
		return serializer.ErrorResponse(
			serializer.CodePasswordCryptError,
			"密码加密失败.",
			err,
		).Result()
	}

	if _, err := dao.RegisterUser(&user); err != nil {
		return serializer.DatabaseErrorResponse(
			"注册失败, 数据库写入失败.",
			err,
		).Result()
	}

	return serializer.BuildUserResponse(user).Result()
}
