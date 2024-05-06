package mySql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tgwp/configs"
	"tgwp/internal/Model"
	"tgwp/log"
)

var DB *gorm.DB

// InitMyRedis 初始化
func InitMySql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.Conf.MySqlConfig.UserName, configs.Conf.MySqlConfig.Password, configs.Conf.MySqlConfig.Host, configs.Conf.MySqlConfig.Port, configs.Conf.MySqlConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log.MyLogger,
	})
	if err != nil {
		logrus.Fatalf("无法连接数据库！: %v", err)
		return
	}
	err = db.AutoMigrate(Model.User{})
	if err != nil {
		logrus.Fatalf("无法迁移数据库！: %v", err)
		return
	}
	DB = db

	logrus.Infof("数据库连接成功！")

}
