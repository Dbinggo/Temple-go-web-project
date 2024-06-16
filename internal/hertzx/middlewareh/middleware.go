package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"tgwp/internal/hertzx/managerh"
	"tgwp/log/zlog"
)

func init() {
	manager.RouteHandler.RegisterMiddleware(manager.LEVEL_GLOBAL, AddTraceId, false)
}

// AddTraceId
//
//	@Description: add traced in logger
//	@return app.HandlerFunc
func AddTraceId() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 假设 Trace ID 存在于 HTTP Header "X-Trace-ID" 中
		traceID := ctx.Request.Header.Get("X-Request-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		c = zlog.NewContext(c, zap.String("traceId", traceID))
		ctx.Next(c)
	}
}
