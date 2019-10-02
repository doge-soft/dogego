package dao

import (
	"dogego/datasource"
	"dogego/models"
)

// 注册用户
func RegisterUser(user *models.User) (*models.User, error) {
	result := datasource.MasterDatabase().Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Value.(*models.User), nil
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
