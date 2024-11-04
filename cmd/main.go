package main

import (
	"tgwp/initalize"
	router "tgwp/internal/example_hertzx/routerh"
	"tgwp/log/zlog"
)

func main() {

	initalize.Init()
	// 工程进入前夕，释放资源
	defer initalize.Eve()
	router.RunServer()
	zlog.Infof("程序运行完成！")

}
