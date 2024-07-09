package server

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"store-chat/tools/commons"
	"store-chat/tools/consts"
	"store-chat/tools/tools"
	"store-chat/tools/types"
	"strconv"
	"time"
)

type Server struct {
	Node         *snowflake.Node
	Log          logx.Logger
	ServerName   string
	ServerId     string
	ServerIp     string
	ServerPort   string
	Buckets      []*Bucket
	LenBucket    uint32
	ClientManage ClientManage
}

var DefaultServer *Server

// NewServer 初始化服务
func NewServer(ServerId string, node *snowflake.Node, l logx.Logger, BucketNumber uint, serverIP string, serverName string) *Server {
	buckets := NewBucket(BucketNumber)
	DefaultServer = &Server{
		Buckets:      buckets,
		LenBucket:    uint32(len(buckets)),
		ServerId:     ServerId,
		Node:         node,
		Log:          l,
		ServerIp:     serverIP,
		ServerName:   serverName,
		ClientManage: new(DefaultClientManage),
	}
	return DefaultServer
}

// GetBucket 获取连接池
func (s *Server) GetBucket(userId int64) *Bucket {
	userIdStr := strconv.FormatInt(userId, 10)
	// 通过cityHash算法 % 池子数量进行取模,得出需要放入哪个连接池里
	idx := tools.CityHash32([]byte(userIdStr), uint32(len(userIdStr))) % s.LenBucket
	return s.Buckets[idx]
}

// writeChannel
// @Desc：写消息协程
// @param：client
func (s *Server) writeChannel(client *Client) {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		// 断开连接
		if client.RoomId == 0 || client.UserId == 0 || client.IsBreak == false {
			s.Log.Infof("%s readClient.close,RoomId || UserId is 0", s.ServerName)
			_ = client.WsConn.Close()
			return
		}
		// 移除业务,重复进入房间不移除原有缓存
		if _, _, err := s.ClientManage.DisConnect(int32(1), client); err != nil {
			s.Log.Errorf("%s 移除client处理业务;userId:%d;ERR:%s", s.ServerName, client.UserId, err.Error())
		}
		// 移除连接池
		s.GetBucket(client.UserId).UnBucket(client)
		_ = client.WsConn.Close()
		s.Log.Infof("%s readClient.close,user:%s,room:%d,roomClientLen:", s.ServerName, client.Name, client.RoomId, len(s.GetBucket(client.UserId).RoomMap[client.RoomId]))
	}()
	for {
		select {
		case message, ok := <-client.Broadcast:
			// 每次写之前，都需要设置超时时间，如果只设置一次就会出现总是超时
			_ = client.WsConn.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				s.Log.Errorf("%s writeChannel <- client.Broadcast not ok ", s.ServerName)
				_ = client.WsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := client.WsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				s.Log.Errorf("%s Conn.NextWriter fail:%s", s.ServerName, err.Error())
				return
			}
			body, err := jsonx.Marshal(message.Body)
			if err != nil {
				s.Log.Errorf("%s jsonx.Marshal() fail:%s", s.ServerName, err.Error())
				continue
			}
			_, _ = w.Write(body)
			if err = w.Close(); err != nil {
				s.Log.Errorf("%s w.Close() fail:%s", s.ServerName, err.Error())
				return
			}
		case <-ticker.C:
			// 每次写之前，都需要设置超时时间，如果只设置一次就会出现总是超时
			_ = client.WsConn.SetWriteDeadline(time.Now().Add(PingPeriod))
			// 心跳检测
			if err := client.WsConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				s.Log.Errorf("%s WriteMessage fail:%s", s.ServerName, err.Error())
				return
			}
		}
	}
}

// readChannel
// @Desc：读消息协程
// @param：client
func (s *Server) readChannel(client *Client) {
	var (
		code       string
		msg        string
		receiveMsg types.ReceiveMsg
		userClient *UserClient
		ok         bool
		sendMsg    string
		sendClient *Client
	)
	defer func() {
		// 断开连接
		if client.RoomId == 0 || client.UserId == 0 || client.IsBreak == false {
			s.Log.Infof("%s readClient.close,RoomId || UserId is 0", s.ServerName)
			_ = client.WsConn.Close()
			return
		}
		// 移除业务,重复进入房间不移除原有缓存
		if _, _, err := s.ClientManage.DisConnect(int32(receiveMsg.Version), client); err != nil {
			s.Log.Errorf("%s 移除client处理业务;userId:%d;ERR:%s", s.ServerName, client.UserId, err.Error())
		}
		if _, _, err := s.ClientManage.PushBroadcast(receiveMsg, client.UserId, client.Name, client.Name+" 离开了"); err != nil {
			s.Log.Errorf("%s 移除client推送离开信息;userId:%d;ERR:%s", s.ServerName, client.UserId, err.Error())
		}
		// 移除连接池
		s.GetBucket(client.UserId).UnBucket(client)
		_ = client.WsConn.Close()
		s.Log.Infof("%s readClient.close,user:%s,room:%d,roomClientLen:", s.ServerName, client.Name, client.RoomId, len(s.GetBucket(client.UserId).RoomMap[client.RoomId]))
	}()
	for {
		messageType, message, err := client.WsConn.ReadMessage()
		if err != nil || (message == nil && messageType == -1) {
			s.Log.Errorf("客户端断连：%d; messageType:%d; fail:%v", messageType, client.UserId, err)
			return
		}
		if err = jsonx.Unmarshal(message, &receiveMsg); err != nil {
			s.Log.Errorf("客户端：%d; message 转换 types.ReceiveMsg fail:%s", client.UserId, err.Error())
			continue
		}
		// 每次需设置读超时时间，否则接收不到
		client.WsConn.SetReadLimit(MaxMessageSize)
		_ = client.WsConn.SetReadDeadline(time.Now().Add(ReadWait))
		client.WsConn.SetPongHandler(func(string) error {
			_ = client.WsConn.SetReadDeadline(time.Now().Add(PongPeriod))
			return nil
		})
		if receiveMsg.Version == 0 || receiveMsg.Operate == 0 || receiveMsg.AutoToken == "" || receiveMsg.RoomId == 0 {
			s.Log.Errorf("%s 消息缺失必要条件 msg:%+v", s.ServerName, receiveMsg)
			continue
		}
		if client.UserId == 0 {
			userClient = &UserClient{}
		}
		receiveMsg.FromClientId = client.ClientId
		receiveMsg.FromUserId = userClient.UserId
		receiveMsg.FromUserName = userClient.Name
		switch receiveMsg.Operate {
		case consts.OPERATE_SINGLE_MSG:
			// 这里麻烦了点，还有另外一个思路，给当前client结构体增加nextClient和lastClient；
			// 1、每次新的client校验有效性后，将上一个client赋值到当前client的lastClient，而nextClient这时候肯定是空的
			// 2、第一步操作的同时，要把上一个client的nextClient进行赋值当前client，
			// 3、client断开前，lastClient和nextClient都要进行维护并重新赋值上下client
			// 这样就可以获取当前client时，直接取上一个或下一个client进行发送，无效在进行查找
			s.Log.Infof("客户端：%s;私聊:%d send:%s", userClient.Name, client.RoomId, receiveMsg.Event.Params)
			if sendMsg, ok = receiveMsg.Event.Params.(string); !ok {
				continue
			}
			bucket := s.GetBucket(receiveMsg.ToUserId)
			if userClient, ok = bucket.UserClientMap[receiveMsg.ToUserId]; ok {
				if sendClient, ok = userClient.RoomClients[receiveMsg.RoomId]; ok {
					receiveMsg.ToClientId = sendClient.ClientId
					receiveMsg.ToUserId = sendClient.UserId
					receiveMsg.ToUserName = sendClient.Name
					if code, msg, err = s.ClientManage.PushBroadcast(receiveMsg, sendClient.UserId, sendClient.Name, sendMsg); err != nil {
						s.Log.Errorf("%s %s 私聊 %s：code:%s msg:%s fail:%s", s.ServerName, client.Name, sendClient.Name, code, msg, err.Error())
					}
				}
			}
		case consts.OPERATE_GROUP_MSG:
			if sendMsg, ok = receiveMsg.Event.Params.(string); ok {
				if code, msg, err = s.ClientManage.PushBroadcast(receiveMsg, client.UserId, client.Name, sendMsg); err != nil {
					s.Log.Errorf("%s %s 进群聊消息发布：code:%s msg:%s fail:%s", s.ServerName, client.Name, code, msg, err.Error())
				}
			}
		case consts.OPERATE_CONN_MSG:
			if code, msg, err, client.RoomId, client.UserId, client.Name = s.ClientManage.InitConnect(receiveMsg); err != nil {
				s.Log.Errorf("%s 校验client用户的有效性 fail:%s", s.ServerName, err.Error())
				return
			} else if code == commons.RESPONSE_SUCCESS && client.UserId > 0 {
				bucket := s.GetBucket(client.UserId)
				userClient = bucket.GetUserClient(client.UserId, client.Name)
				userClient.SystemId = s.ServerIp
				userClient.AutoToken = receiveMsg.AutoToken
				if userClient.AddClientMap(client) == commons.SOCKET_BROADCAST_LOGINED {
					client.IsBreak = false
					s.Log.Errorf("%s roomId:%d user:%s 已进入当前房间，请勿重复进入", s.ServerName, receiveMsg.RoomId, client.Name)
					return
				}

				s.Log.Infof("池子的用户的连接数：%d", len(userClient.RoomClients))
				bucket.AddBucket(client.RoomId, client, userClient)
				// 给当前连接者client信息
				receiveMsg.Operate = consts.OPERATE_SINGLE_MSG
				receiveMsg.Method = consts.METHOD_ENTER_MSG
				receiveMsg.FromUserId = userClient.UserId
				receiveMsg.FromUserName = userClient.Name
				if code, msg, err = s.ClientManage.PushBroadcast(receiveMsg, userClient.UserId, userClient.Name, ""); err != nil {
					s.Log.Errorf("%s %s 返回私人消息发布：code:%s msg:%s fail:%s", s.ServerName, userClient.Name, code, msg, err.Error())
				}
				// 广播群聊通知有人进来了
				receiveMsg.Operate = consts.OPERATE_GROUP_MSG
				receiveMsg.Method = consts.METHOD_NORMAL_MSG
				if code, msg, err = s.ClientManage.PushBroadcast(receiveMsg, userClient.UserId, userClient.Name, fmt.Sprintf("%s 进来了", client.Name)); err != nil {
					s.Log.Errorf("%s %s 进群聊消息发布：code:%s msg:%s fail:%s", s.ServerName, userClient.Name, code, msg, err.Error())
				}
			}
		}
	}
}
