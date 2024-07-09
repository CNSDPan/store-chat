package server

import (
	"github.com/gorilla/websocket"
	"store-chat/tools/commons"
	"store-chat/tools/types"
	"sync"
	"time"
)

type UserClient struct {
	CLock       sync.RWMutex
	SystemId    string
	AutoToken   string
	UserId      int64
	Name        string
	BucketId    uint32
	RoomClients map[int64]*Client
}
type Client struct {
	ConnectTime uint64
	WsConn      *websocket.Conn
	ClientId    int64
	RoomId      int64
	UserId      int64
	Name        string
	IsBreak     bool
	Extend      string
	Broadcast   chan types.WriteMsg
}

// NewClient 创建客户端
func NewClient(wsConn *websocket.Conn) *Client {
	return &Client{
		ConnectTime: uint64(time.Now().Unix()),
		WsConn:      wsConn,
		ClientId:    0,
		RoomId:      0,
		UserId:      0,
		IsBreak:     true,
		Extend:      "",
		Broadcast:   make(chan types.WriteMsg, 10000),
	}
}

// NewUserClient 创建用户客户端组
func NewUserClient() *UserClient {
	return &UserClient{
		SystemId:    "",
		AutoToken:   "",
		UserId:      0,
		Name:        "",
		BucketId:    0,
		RoomClients: make(map[int64]*Client),
	}
}

// AddClientMap 添加群聊|私聊
func (uClient *UserClient) AddClientMap(client *Client) string {
	defer uClient.CLock.Unlock()
	uClient.CLock.Lock()
	// 判断是否已经加入过房间
	if _, ok := uClient.RoomClients[client.RoomId]; ok {
		return commons.SOCKET_BROADCAST_LOGINED
	}
	uClient.RoomClients[client.RoomId] = client
	return commons.RESPONSE_SUCCESS
}

// UnClientMap 移除群聊|私聊
func (uClient *UserClient) UnClientMap(roomId int64) {
	defer uClient.CLock.Unlock()
	uClient.CLock.Lock()
	delete(uClient.RoomClients, roomId)
}

// GetClient 获取一个客户端
func (uClient *UserClient) GetClient(roomId int64) *Client {
	defer uClient.CLock.Unlock()
	uClient.CLock.Lock()
	if client, ok := uClient.RoomClients[roomId]; ok {
		return client
	} else {
		return &Client{
			ClientId:    0,
			WsConn:      nil,
			ConnectTime: 0,
			RoomId:      0,
			UserId:      0,
			Name:        "",
			IsBreak:     false,
			Extend:      "",
		}
	}
}

// CheckSystemId 检测是否当前系统服务
func (uClient *UserClient) CheckSystemId(systemId string) bool {
	defer uClient.CLock.Unlock()
	uClient.CLock.Lock()
	if uClient.SystemId == systemId {
		return true
	} else {
		return false
	}
}
