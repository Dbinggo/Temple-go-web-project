package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tgwp/configs"
	"tgwp/log/zlog"
)

type Mysql struct {
}

// InitDataBases 初始化
func (m *Mysql) initDataBases(config configs.Config) (*gorm.DB, error) {
	dsn := m.getDsn(config)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		zlog.Panicf("MySQL无法连接数据库！: %v", err)
		return nil, err
	}
	zlog.Infof("MySQL连接数据库成功！")
	return db, nil
}
func (m *Mysql) getDsn(config configs.Config) string {
	return config.DB.Dsn
}
func NewMySql() DataBase {
	return &Mysql{}
}
