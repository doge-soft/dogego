package inits

import "{{cookiecutter.app_name}}/modules"

// 初始化'模块'
func InitModule() {
	// 初始化所有模块
	modules.InitAllModules()
}
