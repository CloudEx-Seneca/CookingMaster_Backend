package main

import (
	"flag"
	"fmt"

	"CookingMaster_Backend/app/usercenter/rpc/internal/config"
	"CookingMaster_Backend/app/usercenter/rpc/internal/server"
	"CookingMaster_Backend/app/usercenter/rpc/internal/svc"
	"CookingMaster_Backend/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		usercenter.RegisterUsercenterServer(grpcServer, server.NewUsercenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
