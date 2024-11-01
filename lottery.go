package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"net/http"

	"lottery/internal/config"
	"lottery/internal/handler"
	"lottery/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/lottery.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.MustSetUp()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	ctx := svc.NewServiceContext(c)

	// http服务
	httpServer := rest.MustNewServer(c.RestConf, rest.WithCors(), rest.WithFileServer("/api/file/v1/static", http.Dir("static")))
	handler.RegisterHandlers(httpServer, ctx)
	serviceGroup.Add(httpServer)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	serviceGroup.Start()
}
