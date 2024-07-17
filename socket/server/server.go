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
	userClient := new(UserClient)
	var ok bool
	defer func() {
		ticker.Stop()
		// 断开连接
		if client.UserId == 0 || client.RoomId == 0 || client == nil || client.IsRepeatConn == consts.REPEAT_CONN {
			s.Log.Infof("%s writeChannel.close;UserId and RoomId is 0,client is nil", s.ServerName)
			_ = client.WsConn.Close()
			client = nil
			return
		}
		//if _, _, err := s.ClientManage.DisConnect(int32(1), client.RoomId, userClient.UserId); err != nil {
		//	s.Log.Errorf("%s 移除client处理业务;userId:%d;ERR:%s", s.ServerName, client.UserId, err.Error())
		//}
		// 移除连接池
		s.GetBucket(client.UserId).UnBucket(client)
		_ = client.WsConn.Close()
		s.Log.Infof("%s writeChannel.close,room:%d,user:%s", s.ServerName, client.RoomId, userClient.UserName)
		client = nil
	}()
	for {
		if _, ok = s.GetBucket(client.UserId).UserClientMap[client.UserId]; ok {
			userClient = s.GetBucket(client.UserId).UserClientMap[client.UserId]
		}
		select {
		case handleClose := <-client.HandleClose:
			// websocket重复握手，关闭之前的,但保留最后一次
			if handleClose == "run" {
				b := types.WriteMsgBody{
					Version:      1,
					Operate:      consts.OPERATE_SINGLE_MSG,
					Method:       consts.METHOD_OUT_MSG,
					ResponseTime: "",
					Event:        types.Event{Params: "", Data: "您已在别处登录，该设备自动退出"},
				}
				body, err := jsonx.Marshal(b)
				if err != nil {
					s.Log.Errorf("%s 重复握手 handleClose jsonx.Marshal() fail:%s", s.ServerName, err.Error())
					return
				}
				_ = client.WsConn.SetWriteDeadline(time.Now().Add(WriteWait))
				// 强制下线
				if err = client.WsConn.WriteMessage(websocket.TextMessage, body); err != nil {
					s.Log.Errorf("%s 重复握手 handleClose WriteMessage fail:%s", s.ServerName, err.Error())
					return
				}
				return
			}
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
		code         string
		msg          string
		receiveMsg   types.ReceiveMsg
		userClient   = new(UserClient)
		roomId       int64
		ok           bool
		sendMsg      string
		toUserClient = new(UserClient)
		toClient     *Client
	)
	defer func() {
		// 断开连接
		if client.RoomId == 0 || client.UserId == 0 || client == nil || client.IsRepeatConn == consts.REPEAT_CONN {
			s.Log.Infof("%s readChannel.close,RoomId || UserId is 0", s.ServerName)
			_ = client.WsConn.Close()
			client = nil
			return
		}
		// 移除业务,重复进入房间不移除原有缓存
		//if _, _, err := s.ClientManage.DisConnect(int32(receiveMsg.Version), roomId, userClient.UserId); err != nil {
		//	s.Log.Errorf("%s 移除client处理业务;userId:%d;ERR:%s", s.ServerName, client.UserId, err.Error())
		//}
		if _, _, err := s.ClientManage.PushBroadcast(receiveMsg, userClient.UserId, userClient.UserName, userClient.UserName+" 离开了"); err != nil {
			s.Log.Errorf("%s 移除client推送离开信息;userId:%d;ERR:%s", s.ServerName, client.UserId, err.Error())
		}
		// 移除连接池
		s.GetBucket(client.UserId).UnBucket(client)
		_ = client.WsConn.Close()
		s.Log.Infof("%s readChannel.close,room:%d,user:%s", s.ServerName, client.RoomId, userClient.UserName)
		client = nil
	}()
	for {
		messageType, message, err := client.WsConn.ReadMessage()
		if err != nil || (message == nil && messageType == -1) {
			s.Log.Errorf("客户端【%s】断连; messageType:%d; fail:%v", userClient.UserName, messageType, err)
			return
		}
		if err = jsonx.Unmarshal(message, &receiveMsg); err != nil {
			s.Log.Errorf("客户端【%s】断连; message 转换 types.ReceiveMsg fail:%s", userClient.UserName, err.Error())
			continue
		}
		// 每次需设置读超时时间，否则接收不到
		client.WsConn.SetReadLimit(MaxMessageSize)
		_ = client.WsConn.SetReadDeadline(time.Now().Add(ReadWait))
		client.WsConn.SetPongHandler(func(string) error {
			_ = client.WsConn.SetReadDeadline(time.Now().Add(PongPeriod))
			return nil
		})
		if receiveMsg.Version == 0 || receiveMsg.Operate == 0 || receiveMsg.RoomId == 0 || receiveMsg.AuthToken == "" || receiveMsg.FromUserId == 0 {
			s.Log.Errorf("%s 消息缺失必要条件 msg:%+v", s.ServerName, receiveMsg)
			return
		}
		roomId = receiveMsg.RoomId
		receiveMsg.FromClientId = client.ClientId
		receiveMsg.FromUserName = userClient.UserName
		if userClient.UserId > 0 {
			receiveMsg.FromUserId = userClient.UserId
		}
		switch receiveMsg.Operate {
		case consts.OPERATE_SINGLE_MSG:
			if sendMsg, ok = receiveMsg.Event.Params.(string); !ok {
				continue
			}
			bucket := s.GetBucket(receiveMsg.ToUserId)
			if toUserClient, ok = bucket.UserClientMap[receiveMsg.ToUserId]; ok {
				if toClient, ok = toUserClient.RoomClients[roomId]; ok {
					receiveMsg.ToClientId = toClient.ClientId
					receiveMsg.ToUserId = toUserClient.UserId
					receiveMsg.ToUserName = toUserClient.UserName
					if code, msg, err = s.ClientManage.PushBroadcast(receiveMsg, toUserClient.UserId, toUserClient.UserName, sendMsg); err != nil {
						s.Log.Errorf("%s %s 私聊 %s：code:%s msg:%s fail:%s", s.ServerName, userClient.UserName, toUserClient.UserName, code, msg, err.Error())
					}
				}
			}
		case consts.OPERATE_GROUP_MSG:
			if sendMsg, ok = receiveMsg.Event.Params.(string); ok {
				if code, msg, err = s.ClientManage.PushBroadcast(receiveMsg, 0, "", sendMsg); err != nil {
					s.Log.Errorf("%s %s 进群聊消息发布：code:%s msg:%s fail:%s", s.ServerName, userClient.UserName, code, msg, err.Error())
				}
			}
		case consts.OPERATE_CONN_MSG:
			if code, msg, err, client.UserId, userClient.UserName = s.ClientManage.InitConnect(receiveMsg); err != nil {
				s.Log.Errorf("%s 校验client用户的有效性 fail:%s", s.ServerName, err.Error())
				return
			} else if code == commons.RESPONSE_SUCCESS && client.UserId > 0 {
				client.RoomId = roomId
				bucket := s.GetBucket(client.UserId)
				userClient = bucket.GetUserClient(client.UserId, userClient.UserName)
				userClient.SystemId = s.ServerIp
				userClient.AuthToken = receiveMsg.AuthToken
				bucket.AddBucket(roomId, client, userClient)
				s.Log.Infof("池子的用户的连接数：%d", len(userClient.RoomClients))
				// 给当前连接者client信息
				receiveMsg.Operate = consts.OPERATE_SINGLE_MSG
				receiveMsg.Method = consts.METHOD_ENTER_MSG
				receiveMsg.FromUserId = userClient.UserId
				receiveMsg.FromUserName = userClient.UserName
				if code, msg, err = s.ClientManage.PushBroadcast(receiveMsg, userClient.UserId, userClient.UserName, ""); err != nil {
					s.Log.Errorf("%s %s 返回私人消息发布：code:%s msg:%s fail:%s", s.ServerName, userClient.UserName, code, msg, err.Error())
				}
				// 广播群聊通知有人进来了
				receiveMsg.Operate = consts.OPERATE_GROUP_MSG
				receiveMsg.Method = consts.METHOD_NORMAL_MSG
				if code, msg, err = s.ClientManage.PushBroadcast(receiveMsg, userClient.UserId, userClient.UserName, fmt.Sprintf("%s 进来了", userClient.UserName)); err != nil {
					s.Log.Errorf("%s %s 进群聊消息发布：code:%s msg:%s fail:%s", s.ServerName, userClient.UserName, code, msg, err.Error())
				}
			} else {
				return
			}
		}
	}
}
