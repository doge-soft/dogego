package inits

import (
	"dogego/cache"
	"dogego/datasource"
	"os"
)

// 数据链接初始化
func InitDataConnection() {
	// 连接数据库
	datasource.ConnectDatabase(
		os.Getenv("MYSQL_DSN_MASTER"),
		os.Getenv("MYSQL_DSN_SLAVE"))

	// 连接Redis
	cache.ConnectRedisCache()
}
