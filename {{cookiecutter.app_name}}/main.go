package main

import (
	"{{cookiecutter.app_name}}/api/routers"
	"{{cookiecutter.app_name}}/api/servers"
	_ "{{cookiecutter.app_name}}/conf"
	"sync"
)

func main() {
	router := routers.NewRouter()
	group := sync.WaitGroup{}

	// 启动所有注册过的服务器
	servers.StartServers(&group, router)
}
