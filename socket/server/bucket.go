package server

import (
	"fmt"
	"store-chat/tools/consts"
	"store-chat/tools/tools"
	"store-chat/tools/types"
	"sync"
)

type Bucket struct {
	Clock         sync.RWMutex
	UserClientMap map[int64]*UserClient
	RoomMap       map[int64]*Room
	Routines      chan types.WriteMsg
	Idx           uint32
}

type Room struct {
	RoomId   int64
	RoomName string
	Clients  []*Client
}

// NewBucket 初始化池子
func NewBucket(cpu uint) []*Bucket {
	buckets := make([]*Bucket, cpu)
	for i := uint(0); i < cpu; i++ {
		roomMap := make(map[int64]*Room, 1)
		buckets[i] = &Bucket{
			Clock:         sync.RWMutex{},
			UserClientMap: make(map[int64]*UserClient),
			RoomMap:       roomMap,
			Routines:      make(chan types.WriteMsg, 1000),
			Idx:           uint32(i),
		}
		buckets[i].RoutineWriteMsg()
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
		userClient.UserName = userName
		userClient.BucketId = b.Idx
	}
	return userClient
}

// AddBucket 将客户端加入连接池
func (b *Bucket) AddBucket(roomId int64, client *Client, userClient *UserClient) {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	userClient.AddClientMap(roomId, client)
	b.UserClientMap[userClient.UserId] = userClient
	if _, ok := b.RoomMap[roomId]; !ok {
		room := &Room{
			RoomId:   roomId,
			RoomName: "",
			Clients:  make([]*Client, 0),
		}
		room.Clients = append(room.Clients, client)
		b.RoomMap[roomId] = room
	} else {
		b.RoomMap[roomId].RoomId = roomId
		b.RoomMap[roomId].Clients = append(b.RoomMap[roomId].Clients, client)
	}
}

// UnBucket 将客户端移除连接池
func (b *Bucket) UnBucket(client *Client) {
	defer b.Clock.Unlock()
	b.Clock.Lock()
	if userClient, ok := b.UserClientMap[client.UserId]; ok {
		fmt.Printf("移除连接池:rid:%d u:%s", client.RoomId, userClient.UserName)
		userClient.UnClientMap(client.RoomId)
		if len(userClient.RoomClients) == 0 {
			delete(b.UserClientMap, client.UserId)
		}
	}
	if room, ok := b.RoomMap[client.RoomId]; ok {
		var newRooms = b.RoomMap[client.RoomId].Clients[:0]
		for _, cl := range room.Clients {
			if cl.ClientId == client.ClientId {
				continue
			}
			newRooms = append(newRooms, cl)
		}
		b.RoomMap[client.RoomId].Clients = newRooms
	}
}

func (b *Bucket) BroadcastRoom(writeMsg types.WriteMsg) (err error) {
	select {
	case b.Routines <- writeMsg:
	default:
	}
	return
}

// RoutineWriteMsg
// @Desc：每个池子有单独接收订阅者传输过来的数据并处理的协程
func (b *Bucket) RoutineWriteMsg() {
	go func() {
		for {
			select {
			case writeMsg := <-b.Routines:
				if writeMsg.Operate == consts.OPERATE_SINGLE_MSG {
					if userClient := b.GetUserClient(writeMsg.ToUserId, ""); userClient != nil {
						if client := userClient.GetClient(writeMsg.RoomId); client != nil {
							_ = client.Push(writeMsg)
						}
					}
				} else if writeMsg.Operate == consts.OPERATE_GROUP_MSG {
					if room, ok := b.RoomMap[writeMsg.RoomId]; ok {
						for _, client := range room.Clients {
							_ = client.Push(writeMsg)
						}
					}
				}
				DefaultServer.Log.Errorf("idx【%d】循环【%d:%s】chan【%d】", b.Idx, writeMsg.RoomId, tools.StoreMap[writeMsg.RoomId].Name, len(b.Routines))
			}
		}
	}()
}
