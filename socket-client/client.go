package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store-chat/socket-client/server"
	"store-chat/tools/tools"
	"store-chat/tools/yamls"
)

type ClientMap struct {
	Clients []*server.DefaultUser
	Log     logx.Logger
	Conf    *yamls.SocketClientCon
}

var DefaultClient *ClientMap

func main() {

	DefaultClient = new(ClientMap)
	DefaultClient.Clients = make([]*server.DefaultUser, 0)
	DefaultClient.Log = logx.WithContext(context.Background())
	DefaultClient.Conf = yamls.SocketClientConf

	client1 := new(server.DefaultUser)
	client1.Log = logx.WithContext(context.Background())
	client1.Conf = DefaultClient.Conf
	client1.InitUserInfo("2gDGQwDxsrX0UG8yRbophdHxHqD")

	client2 := new(server.DefaultUser)
	client2.Log = logx.WithContext(context.Background())
	client2.Conf = DefaultClient.Conf
	client2.InitUserInfo("2gDGQugkyFF4MI10hK7WfT3W3Pe")

	client3 := new(server.DefaultUser)
	client3.Log = logx.WithContext(context.Background())
	client3.Conf = DefaultClient.Conf
	client3.InitUserInfo("2kPybdu8GObZm5SVwHF1TLUrCE9")

	DefaultClient.Clients = append(DefaultClient.Clients, client1, client2, client3)
	DefaultClient.Run()
	select {}
}

func (c *ClientMap) Run() {
	var roomMap = tools.StoreMap
	for _, client := range c.Clients {
		idx := uint32(0)
		for _, room := range roomMap {
			client.Operator(idx, c.Conf.WebsocketUrl, room.StoreID, room.Name)
			idx++
		}
	}
}
