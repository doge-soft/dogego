package servers

import (
	"github.com/doge-soft/dogego_server_http"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
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

func StartServers(group *sync.WaitGroup, router *gin.Engine) {
	for _, server := range Servers {
		group.Add(1)
		go func(server func(router *gin.Engine) error) {
			err := server(router)

			defer group.Done()

			if err != nil {
				group.Done()
				log.Fatal(err)
			}
		}(server)
	}

	group.Wait()
}
