package dao

import (
	"{{cookiecutter.app_name}}/datasource"
	"{{cookiecutter.app_name}}/models"
	"errors"
)

// 注册用户
func RegisterUser(user *models.User) (*models.User, error) {
	result := datasource.MasterDatabase().Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Value.(*models.User), nil
}

// 登录用户
func LoginUser(phone_number string, password string) (*models.User, error, bool) {
	user := models.User{}

	if err := datasource.SlaveDatabase().Where("phone_number = ?", phone_number).First(&user).Error; err != nil {
		return nil, err, false
	}

	if user.CheckPassword(password) == false {
		return nil, errors.New("账号或密码错误."), false
	}

	return &user, nil, true
}

// 获取用户
func GetUserById(ID interface{}) (*models.User, error) {
	var user models.User

	err := datasource.SlaveDatabase().First(&user, ID).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// 更新用户信息
func UpdateUserProfile(user *models.User) error {
	err := datasource.MasterDatabase().Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}
