package initalize

import (
	"tgwp/configs"
	"tgwp/db/databases"
	"tgwp/db/myRedis"
	"tgwp/global"
	"tgwp/log/zlog"
)

func InitDataBase(config configs.Config) {
	switch config.DB.Driver {
	case "mysql":
		databases.InitDataBases(databases.NewMySql(), config)
		break
	}
	if config.App.Env != "pro" {
		err := global.DB.AutoMigrate()
		if err != nil {
			zlog.Fatalf("数据库迁移失败！")
		}
	}
	zlog.Infof("数据库初始化成功！")
}
func InitRedis(config configs.Config) {
	if config.Redis.Enable {
		var err error
		global.Rdb, err = myRedis.GetRedisClient(config)
		if err != nil {
			zlog.Errorf("无法初始化Redis : %v", err)
		}
	} else {
		zlog.Warnf("不使用Redis")
	}

}
