package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"chatskn/app/ws/rpc/internal/config"
	"chatskn/app/ws/rpc/internal/handler"
	"chatskn/app/ws/rpc/internal/mqs"
	"chatskn/app/ws/rpc/internal/svc"
)

var configFile = flag.String("f", "etc/ws.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := context.Background()
	svcCtx := svc.NewServiceContext(c)

	//s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
	//	pb.RegisterWsServer(grpcServer, server.NewWsServer(svcCtx))
	//
	//	if c.Mode == service.DevMode || c.Mode == service.TestMode {
	//		reflection.Register(grpcServer)
	//	}
	//})
	s := rest.MustNewServer(c.RestConf)
	defer s.Stop()
	handler.RegisterHandlers(s, svcCtx)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	serviceGroup.Add(s)

	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
		serviceGroup.Add(mq)
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	serviceGroup.Start()
}
