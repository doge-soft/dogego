package servers

import (
	"github.com/doge-soft/dogego_server_http"
	"github.com/gin-gonic/gin"
	"log"
)

var Servers []func(router *gin.Engine) error

// 注册所有的服务器
func RegisterServers() {
	// 注册HTTP服务器
	RegisterServer(dogego_server_http.HTTPServerProtocol)
}

func RegisterServer(s func(router *gin.Engine) error) {
	Servers = append(Servers, s)
}

func StartServers(router *gin.Engine) {
	for _, server := range Servers {
		go func(serverCpy func(router *gin.Engine) error) {
			err := serverCpy(router)

			if err != nil {
				log.Fatal(err)
			}
		}(server)
	}
}
