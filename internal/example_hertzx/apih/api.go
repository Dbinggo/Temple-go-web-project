package apih

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/route"
	manager "tgwp/internal/example_hertzx/managerh"
	"tgwp/log/zlog"
)

// only for test
func init() {
	manager.RouteHandler.RegisterRouter(manager.LEVEL_GLOBAL, func(r *route.RouterGroup) {
		r.GET("/test01", Test)
	})
}
func Test(ctx context.Context, c *app.RequestContext) {
	zlog.Infof("load - test")
	zlog.CtxInfof(ctx, "load ctx info - test")
	zlog.CtxWarnf(ctx, "load ctx warn - test")
	zlog.CtxErrorf(ctx, "load ctx error - test")
	c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
}
