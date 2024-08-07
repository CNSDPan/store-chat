package socket_client

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store-chat/test/socket"
	"store-chat/tools/tools"
	"sync"
)

type ClientMap struct {
	Clock   sync.RWMutex
	Clients []*socket.DefaultUser
	Log     logx.Logger
}

var DefaultClient *ClientMap

func main() {
	DefaultClient = new(ClientMap)
	DefaultClient.Clients = make([]*socket.DefaultUser, 0)
	DefaultClient.Log = logx.WithContext(context.Background())

	client1 := new(socket.DefaultUser)
	client1.Log = logx.WithContext(context.Background())
	client1.InitUserInfo("2gDGQwDxsrX0UG8yRbophdHxHqD")
	client2 := new(socket.DefaultUser)
	client2.Log = logx.WithContext(context.Background())
	client2.InitUserInfo("2gDGQugkyFF4MI10hK7WfT3W3Pe")

	DefaultClient.Clients = append(DefaultClient.Clients, client1, client2)
	DefaultClient.Run()
	select {}
}

func (c *ClientMap) Run() {
	var roomMap = tools.StoreMap
	// nginx 负载均衡代理地址
	var url = "ws://websocket.cn:6990/ws"
	for idx, client := range c.Clients {
		for _, room := range roomMap {
			client.Operator(uint32(idx), url, room.StoreID, room.Name)
		}
	}
}
