package rpc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"store-chat/rpc/socket/client/broadcast"
	"store-chat/rpc/socket/client/ping"
	"store-chat/tools/yamls"
)

var GrpcSocket *GrpcSocketLogic

type GrpcSocketLogic struct {
	Ping      ping.Ping
	Broadcast broadcast.Broadcast
}

func NewRpcSocket() {
	conn := zrpc.MustNewClient(yamls.RpcSocketClientConf.RpcClientConf)
	GrpcSocket = new(GrpcSocketLogic)
	GrpcSocket.Ping = ping.NewPing(conn)
	GrpcSocket.Broadcast = broadcast.NewBroadcast(conn)
}
