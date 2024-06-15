package databases

import (
	"gorm.io/gorm"
	"tgwp/configs"
	"tgwp/global"
	"tgwp/log/zlog"
)

type DataBase interface {
	getDsn(config configs.Config) string
	initDataBases(config configs.Config) (*gorm.DB, error)
}

func InitDataBases(base DataBase, config configs.Config) {
	var err error
	global.DB, err = base.initDataBases(config)
	if err != nil {
		zlog.Fatalf("无法初始化数据库 %v", err)
		return
	}
	zlog.Infof("初始化数据库成功！")
	return
}
