package migrations

import (
	"{{cookiecutter.app_name}}/models"
	"github.com/jinzhu/gorm"
)

// 迁移数据库
func MigrationModels(context *gorm.DB) error {
	context.AutoMigrate(&models.User{})
	return nil
}
