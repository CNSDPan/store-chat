package socket

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"store-chat/rpc/socket/client/broadcast"
	"store-chat/rpc/socket/client/socket"
	"store-chat/tools/yamls"
)

func New() broadcast.Broadcast {
	c := yamls.RpcSocketClientConf
	conn := zrpc.MustNewClient(c.RpcClientConf)
	_, err := socket.NewSocket(conn).Ping(context.Background(), &socket.ReqPing{})
	if err != nil {
		panic("初始化 rpc.socket 失败：" + err.Error())
	}
	rpcSocketBroadcast := broadcast.NewBroadcast(conn)
	fmt.Println("初始化 rpc.socket ok")
	return rpcSocketBroadcast
}
