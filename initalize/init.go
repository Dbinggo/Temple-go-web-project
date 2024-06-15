package initalize

import (
	"tgwp/global"
	"tgwp/utils"
)

func Init() {
	InitLog(global.Config)
	InitPath()
	InitConfig()
	InitLog(global.Config)
	InitDataBase(*global.Config)
	InitRedis(*global.Config)
}
func InitPath() {
	global.Path = utils.GetRootPath("")
}
