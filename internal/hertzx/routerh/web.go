package router

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"tgwp/global"
	_ "tgwp/internal/hertzx/apih"
	manager "tgwp/internal/hertzx/managerh"
	_ "tgwp/internal/hertzx/middlewareh"
	"tgwp/log/zlog"
)

func RunServer() {
	h, err := listen()
	if err != nil {
		zlog.Errorf("Listen error: %v", err)
		panic(err.Error())
	}
	h.Spin()
}

func listen() (*server.Hertz, error) {

	h := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", global.Config.App.Host, global.Config.App.Port)))
	manager.RouteHandler.Register(h)
	return h, nil
}
