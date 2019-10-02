package main

import (
	"dogego/api/routers"
	"dogego/api/servers"
	_ "dogego/conf"
)

func main() {
	router := routers.NewRouter()
	channel := make(chan bool)

	// 启动所有注册过的服务器
	servers.StartServers(router)

	<-channel
}
