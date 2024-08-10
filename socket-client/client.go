package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store-chat/socket-client/server"
	"store-chat/tools/tools"
)

type ClientMap struct {
	Clients []*server.DefaultUser
	Log     logx.Logger
}

var DefaultClient *ClientMap

func main() {
	DefaultClient = new(ClientMap)
	DefaultClient.Clients = make([]*server.DefaultUser, 0)
	DefaultClient.Log = logx.WithContext(context.Background())

	client1 := new(server.DefaultUser)
	client1.Log = logx.WithContext(context.Background())
	client1.InitUserInfo("2gDGQwDxsrX0UG8yRbophdHxHqD")
	client2 := new(server.DefaultUser)
	client2.Log = logx.WithContext(context.Background())
	client2.InitUserInfo("2gDGQugkyFF4MI10hK7WfT3W3Pe")
	client3 := new(server.DefaultUser)
	client3.Log = logx.WithContext(context.Background())
	client3.InitUserInfo("2kPybdu8GObZm5SVwHF1TLUrCE9")
	DefaultClient.Clients = append(DefaultClient.Clients, client1, client2, client3)
	DefaultClient.Run()
	select {}
}

func (c *ClientMap) Run() {
	var roomMap = tools.StoreMap
	// nginx 负载均衡代理地址
	var url = "ws://192.168.33.10:6990/ws"
	for _, client := range c.Clients {
		idx := uint32(0)
		for _, room := range roomMap {
			client.Operator(idx, url, room.StoreID, room.Name)
			idx++
		}
	}
}
