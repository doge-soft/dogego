package datasource

import "github.com/jinzhu/gorm"

func SlaveDatabase() *gorm.DB {
	return slaveDB
}
