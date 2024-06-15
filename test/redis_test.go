package test

import (
	"context"
	"testing"
	"tgwp/global"
	"tgwp/initalize"
	"tgwp/log/zlog"
)

func TestRedis(T *testing.T) {
	initalize.Init()
	global.Rdb.Set(context.Background(), "test_redis", "1111", 0)
	v1 := global.Rdb.Get(context.Background(), "test_redis")
	global.Rdb.Set(context.Background(), "test_redis", "2222", 0)
	v2 := global.Rdb.Get(context.Background(), "test_redis")
	global.Rdb.Set(context.Background(), "test_redis", "3333", 0)
	v3 := global.Rdb.Get(context.Background(), "test_redis")
	zlog.Infof("%v,%v,%v", v1, v2, v3)
}
