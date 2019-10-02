package conf

import (
	"dogego/inits"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	// 初始化
	inits.Init()
}
