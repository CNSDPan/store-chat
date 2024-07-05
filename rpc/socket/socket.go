package main

import (
	"flag"
	"fmt"
	"store-chat/dbs"
	"store-chat/rpc/socket/internal/config"
	broadcastServer "store-chat/rpc/socket/internal/server/broadcast"
	pingServer "store-chat/rpc/socket/internal/server/ping"
	"store-chat/rpc/socket/internal/svc"
	"store-chat/rpc/socket/pb/socket"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/socket.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	// 初始化redis
	if _, err := dbs.NewRedisClient(); err != nil {
		panic(c.ServerName + " redisClient init fail:" + err.Error())
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		socket.RegisterPingServer(grpcServer, pingServer.NewPingServer(ctx))
		socket.RegisterBroadcastServer(grpcServer, broadcastServer.NewBroadcastServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
