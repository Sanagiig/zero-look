package main

import (
	"flag"
	"fmt"

	"zero-look/app/travel/cmd/api/internal/config"
	"zero-look/app/travel/cmd/api/internal/handler"
	"zero-look/app/travel/cmd/api/internal/svc"

	"github.com/Sanagiig/zero-look/common/middleware"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/travel.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
