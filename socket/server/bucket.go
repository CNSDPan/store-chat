package server

import (
	"store-chat/tools/consts"
	"store-chat/tools/types"
	"sync"
)

type Bucket struct {
	Clock         sync.RWMutex
	UserClientMap map[int64]*UserClient
	RoomMap       map[int64][]*Client
	Routines      chan types.WriteMsg
	Idx           uint32
}

// NewBucket 初始化池子
func NewBucket(cpu uint) []*Bucket {
	buckets := make([]*Bucket, cpu)
	for i := uint(0); i < cpu; i++ {
		buckets[i] = &Bucket{
			Clock:         sync.RWMutex{},
			UserClientMap: make(map[int64]*UserClient),
			RoomMap:       make(map[int64][]*Client),
			Routines:      make(chan types.WriteMsg, 1000),
			Idx:           uint32(i),
		}
		go buckets[i].RoutineWriteMsg()
	}
	return buckets
}

// GetUserClient 获取客户端连接池
func (b *Bucket) GetUserClient(userId int64, userName string) *UserClient {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	userClient, ok := b.UserClientMap[userId]
	if !ok {
		userClient = NewUserClient()
		userClient.UserId = userId
		userClient.Name = userName
		userClient.BucketId = b.Idx
	}
	return userClient
}

// AddBucket 将客户端加入连接池
func (b *Bucket) AddBucket(roomId int64, client *Client, userClient *UserClient) {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	b.UserClientMap[userClient.UserId] = userClient
	b.RoomMap[roomId] = append(b.RoomMap[roomId], client)
}

// UnBucket 将客户端移除连接池
func (b *Bucket) UnBucket(client *Client) {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	if userClient, ok := b.UserClientMap[client.UserId]; ok {
		userClient.UnClientMap(client.RoomId)
		if len(userClient.RoomClients) < 1 {
			delete(b.UserClientMap, client.UserId)
		}

	}
	var newRooms = b.RoomMap[client.RoomId][:0]
	if clients, ok := b.RoomMap[client.RoomId]; ok {
		for _, cl := range clients {
			if cl.ClientId == client.ClientId {
				continue
			}
			newRooms = append(newRooms, cl)
		}
	}
	b.RoomMap[client.RoomId] = newRooms
}

// RoutineWriteMsg
// @Desc：每个池子有单独接收订阅者传输过来的数据并处理的协程
func (b *Bucket) RoutineWriteMsg() {
	for {
		select {
		case writeMsg := <-b.Routines:
			switch writeMsg.Operate {
			case consts.OPERATE_SINGLE_MSG:
				if userClient := b.GetUserClient(writeMsg.ToUserId, ""); userClient != nil {
					if client := userClient.GetClient(writeMsg.RoomId); client != nil {
						client.Broadcast <- writeMsg
					}
				}
			case consts.OPERATE_GROUP_MSG:
				if clients, ok := b.RoomMap[writeMsg.RoomId]; ok {
					for _, client := range clients {
						client.Broadcast <- writeMsg
					}
				}
			}
		}
	}
}
