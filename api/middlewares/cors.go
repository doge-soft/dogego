package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域中间件
func Cors(cors_domain string) gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Authorizion"}
	config.AllowCredentials = true

	// 检查是否是开发环境
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境需要配置跨域域名，否则403
		config.AllowOrigins = []string{cors_domain}
	} else {
		config.AllowAllOrigins = true
	}

	return cors.New(config)
}
