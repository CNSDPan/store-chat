package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"store-chat/dbs"
	"store-chat/socket/mq"
	"store-chat/socket/rpc"
	"store-chat/socket/server"
	"strconv"

	"store-chat/socket/internal/config"
	"store-chat/socket/internal/handler"
	"store-chat/socket/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("socket.file", "etc/socket-api.yaml", "the config file")
var Module = "socket服务:"

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	Module = c.ServerName

	s := rest.MustNewServer(c.RestConf)
	defer s.Stop()

	/********初始化********/
	var (
		err      error
		serverIp string
		nodeId   int64
		node     *snowflake.Node
		l        logx.Logger
		cont     = context.Background()
	)
	l = logx.WithContext(cont)
	if serverIp, err = os.Hostname(); err != nil {
		panic(Module + " 获取服务信息 fail:" + err.Error())
	}
	// 服务uuid节点池
	if nodeId, err = strconv.ParseInt(c.ServiceId, 10, 64); err != nil {
		panic(Module + " serverId server string to int64 fail:" + err.Error())
	}
	if node, err = snowflake.NewNode(nodeId); err != nil {
		panic(Module + " start server newNode func fail:" + err.Error())
	}
	// 初始化redis
	if _, err = dbs.NewRedisClient(); err != nil {
		panic(Module + " redisClient init fail:" + err.Error())
	}
	// 初始化rpc.socket.client端
	rpc.NewRpcSocket()
	// 订阅消息
	if subReceive, err := mq.NewSubscribe(); err != nil {
		panic(Module + " mq.NewSubscribe init fail:" + err.Error())
	} else {
		subReceive.SubReceive()
	}
	// 初始化websocket服务
	socket := server.NewServer(c.ServiceId, node, l, c.BucketNumber, serverIp, Module)
	/********初始化********/

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(s, ctx, socket)

	fmt.Printf("%s Starting server at %s:%d...\n", Module, serverIp, c.Port)
	s.Start()
}
