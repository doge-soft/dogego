package routers

import (
	"dogego/api/controllers"
	"dogego/api/middlewares"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 中间件, 顺序不能乱
	router.Use(middlewares.Cors(os.Getenv("CORS_DOMAIN")))

	v1 := router.Group("/api/v1")
	{
		// 心跳检查接口
		v1.POST("/ping", controllers.Ping)
	}

	return router
}
