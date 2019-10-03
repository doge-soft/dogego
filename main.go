package main

import (
	"dogego/api/routers"
	"dogego/api/servers"
	_ "dogego/conf"
	"sync"
)

func main() {
	router := routers.NewRouter()
	group := sync.WaitGroup{}

	// 启动所有注册过的服务器
	servers.StartServers(&group, router)
}
