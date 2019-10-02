package datasource

import (
	"dogego/migrations"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var (
	masterDB *gorm.DB
	slaveDB *gorm.DB
)

func ConnectDatabase(masterDSN string, slaveDSN string) {
	var err error

	masterDB, err = gorm.Open("mysql", masterDSN)

	//设置连接池
	//空闲
	masterDB.DB().SetMaxIdleConns(50)
	//打开
	masterDB.DB().SetMaxOpenConns(100)
	//超时
	masterDB.DB().SetConnMaxLifetime(time.Second * 30)

	if err != nil {
		log.Fatal(err)
	}

	slaveDB, err = gorm.Open("mysql", slaveDSN)

	//设置连接池
	//空闲
	slaveDB.DB().SetMaxIdleConns(50)
	//打开
	slaveDB.DB().SetMaxOpenConns(100)
	//超时
	slaveDB.DB().SetConnMaxLifetime(time.Second * 30)

	if err != nil {
		log.Fatal(err)
	}

	err = migrations.MigrationModels(masterDB)

	if err != nil {
		log.Fatal(err)
	}
}
