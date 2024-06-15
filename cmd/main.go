package main

import (
	"tgwp/initalize"
	"tgwp/log/zlog"
)

func main() {

	initalize.Init()
	// 工程进入前夕，释放资源
	defer initalize.Eve()
	zlog.Infof("程序运行成功！")

}
