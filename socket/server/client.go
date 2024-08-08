package server

import (
	"github.com/gorilla/websocket"
	"store-chat/tools/consts"
	"store-chat/tools/types"
	"sync"
	"time"
)

type UserClient struct {
	CLock       sync.RWMutex
	SystemId    string
	AuthToken   string
	UserId      int64
	UserName    string
	BucketId    uint32
	RoomClients map[int64]*Client
}

type Client struct {
	ConnectTime  uint64
	WsConn       *websocket.Conn
	ClientId     int64
	RoomId       int64
	UserId       int64
	IsRepeatConn string
	Extend       string
	Broadcast    chan types.WriteMsg
	HandleClose  chan string
}

// NewClient 创建客户端
func NewClient(wsConn *websocket.Conn) *Client {
	return &Client{
		ConnectTime:  uint64(time.Now().Unix()),
		WsConn:       wsConn,
		ClientId:     0,
		RoomId:       0,
		UserId:       0,
		IsRepeatConn: "",
		Extend:       "",
		Broadcast:    make(chan types.WriteMsg, 10000),
		HandleClose:  make(chan string, 100),
	}
}

// NewUserClient 创建用户客户端组
func NewUserClient() *UserClient {
	return &UserClient{
		SystemId:    "",
		AuthToken:   "",
		UserId:      0,
		UserName:    "",
		BucketId:    0,
		RoomClients: make(map[int64]*Client),
	}
}

// AddClientMap 添加群聊|私聊
func (uClient *UserClient) AddClientMap(roomId int64, client *Client) {
	defer uClient.CLock.Unlock()
	uClient.CLock.Lock()
	// 判断是否已经加入过房间
	if lastClient, ok := uClient.RoomClients[roomId]; ok && lastClient.ClientId != client.ClientId {
		lastClient.IsRepeatConn = consts.REPEAT_CONN
		lastClient.HandleDown()
	}
	client.IsRepeatConn = consts.FIRST_CONN
	uClient.RoomClients[roomId] = client
	return
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
			ClientId:     0,
			WsConn:       nil,
			ConnectTime:  0,
			IsRepeatConn: "",
			Extend:       "",
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

func (client *Client) Push(writeMsg types.WriteMsg) (err error) {
	select {
	case client.Broadcast <- writeMsg:
	default:
	}
	return
}

func (client *Client) HandleDown() {
	select {
	case client.HandleClose <- "run":
	default:
	}
	return
}
