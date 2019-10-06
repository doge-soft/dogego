package inits

// 初始化所有模块
func Init() {
	// 初始化数据连接
	InitDataConnection()
	// 初始化协议
	InitProtocol()
	// 初始化'模块'
	InitModule()
}
