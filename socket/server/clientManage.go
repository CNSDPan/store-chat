package server

import (
	"context"
	"google.golang.org/protobuf/types/known/anypb"
	"store-chat/rpc/socket/pb/socket"
	"store-chat/socket/rpc"
	"store-chat/tools/commons"
	"store-chat/tools/consts"
	"store-chat/tools/types"
)

type DefaultClientManage struct {
}

type ClientManage interface {
	// InitConnect 交给业务层校验autoToken和处理业务
	InitConnect(receiveMsg types.ReceiveMsg) (code string, msg string, err error, userId int64, userName string)
	// DisConnect 断连交给业务层处理其他业务
	DisConnect(version int32, roomId int64, userId int64) (code string, msg string, err error)
	// PushSingle 私聊发布交给业务层处理
	PushBroadcast(receiveMsg types.ReceiveMsg, toUserId int64, toUserName string, sendMsg string) (code string, msg string, err error)
}

// InitConnect
// @Desc：连接处理业务逻辑
// @param：autoToken
// @return：userId
// @return：userName
func (cManage *DefaultClientManage) InitConnect(receiveMsg types.ReceiveMsg) (code string, msg string, err error, userId int64, userName string) {
	var (
		result = &socket.Result{}
		data   = &socket.EventDataLogin{}
	)
	defer func() {
		code = result.Code
		msg = result.Msg
	}()
	result, err = rpc.GrpcSocket.Broadcast.BroadcastLogin(context.Background(), &socket.ReqBroadcastMsg{
		Version:   int32(receiveMsg.Version),
		Operate:   int32(consts.OPERATE_CONN_MSG),
		Method:    receiveMsg.Method,
		AutoToken: receiveMsg.AutoToken,
		RoomId:    receiveMsg.RoomId,
	})
	if err != nil {
		result.Code, result.Msg = commons.GetCodeMessage(commons.RESPONSE_FAIL)
		return
	}
	if result.Code == commons.RESPONSE_SUCCESS {
		err = result.Data.UnmarshalTo(data)
		if err != nil {
			result.Code, result.Msg = commons.GetCodeMessage(commons.RESPONSE_FAIL)
		}
		userId = data.UserId
		userName = data.UserName
	}
	return
}

// DisConnect
// @Desc：断连处理业务逻辑
// @param：userId
func (cManage *DefaultClientManage) DisConnect(version int32, roomId int64, userId int64) (code string, msg string, err error) {
	var (
		result = &socket.Result{}
	)
	defer func() {
		code = result.Code
		msg = result.Msg
	}()
	result, err = rpc.GrpcSocket.Broadcast.BroadcastOut(context.Background(), &socket.ReqBroadcastMsg{
		Version:    version,
		RoomId:     roomId,
		FromUserId: userId,
	})
	if err != nil {
		result.Code, result.Msg = commons.GetCodeMessage(commons.RESPONSE_FAIL)
		return
	}
	return
}

// PushBroadcast
// @Desc：私聊|群聊广播消息发布
// @param：receiveMsg
// @param：toUserId
// @param：toUserName
// @param：sendMsg
// @return：code
// @return：msg
// @return：err
func (cManage *DefaultClientManage) PushBroadcast(receiveMsg types.ReceiveMsg, toUserId int64, toUserName string, sendMsg string) (code string, msg string, err error) {
	var (
		result = &socket.Result{}
		params *anypb.Any
	)
	switch receiveMsg.Method {
	case consts.METHOD_ENTER_MSG:
		params, err = anypb.New(&socket.EventParamsLogin{
			RoomId:   receiveMsg.RoomId,
			ClientId: receiveMsg.FromClientId,
			UserId:   receiveMsg.FromUserId,
			UserName: receiveMsg.FromUserName,
		})
	case consts.METHOD_NORMAL_MSG:
		params, err = anypb.New(&socket.EventParamsNormal{
			Message: sendMsg,
		})
	default:
		return
	}
	if err != nil {
		code, msg = commons.GetCodeMessage(commons.SOCKET_BROADCAST_NORMAL_SINGLE)
		return
	}
	result, err = rpc.GrpcSocket.Broadcast.BroadcastNormal(context.Background(), &socket.ReqBroadcastMsg{
		Version:      int32(receiveMsg.Version),
		Operate:      int32(receiveMsg.Operate),
		Method:       receiveMsg.Method,
		AutoToken:    receiveMsg.AutoToken,
		RoomId:       receiveMsg.RoomId,
		FromUserId:   receiveMsg.FromUserId,
		FromUserName: receiveMsg.FromUserName,
		ToUserId:     toUserId,
		ToUserName:   toUserName,
		Event: &socket.BodyEvent{
			Params: params,
			Data:   nil,
		},
	})
	code = result.Code
	msg = result.Msg
	return
}
