package datasource

import "github.com/jinzhu/gorm"

func MasterDatabase() *gorm.DB {
	return masterDB
}
