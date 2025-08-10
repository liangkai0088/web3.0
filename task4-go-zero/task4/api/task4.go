package main

import (
	"flag"
	"fmt"
	"task4-go-zero/task4/api/internal/config"
	"task4-go-zero/task4/api/internal/handler"
	"task4-go-zero/task4/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/task4-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 创建服务上下文
	serverCtx := svc.NewServiceContext(c)

	// 创建HTTP服务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 注册路由
	handler.RegisterHandlers(server, serverCtx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
