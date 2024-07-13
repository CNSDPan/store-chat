package broadcastlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"store-chat/dbs"
	"store-chat/rpc/socket/internal/svc"
	"store-chat/rpc/socket/pb/socket"
	"store-chat/tools/commons"
	"store-chat/tools/consts"
	"store-chat/tools/types"
	"time"
)

type BroadcastNormalLogic struct {
	module string
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBroadcastNormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastNormalLogic {
	return &BroadcastNormalLogic{
		module: svcCtx.Config.ServerName + "BroadcastNormal",
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BroadcastNormalLogic) BroadcastNormal(in *socket.ReqBroadcastMsg) (result *socket.Result, rpcErr error) {
	var (
		err      error
		writeMsg types.WriteMsg
		body     []byte
	)
	result = &socket.Result{
		Module: l.module,
		Code:   commons.RESPONSE_SUCCESS,
	}
	defer func() {
		result.Code, result.Msg = commons.GetCodeMessage(result.Code)
		if err != nil {
			result.ErrMsg = err.Error()
			l.Logger.Errorf("%s Broadcast ope:%d code:%s msg:%s fail:%s", result.Module, in.Operate, result.Code, result.Msg, err.Error())
		}
	}()
	writeMsg = types.WriteMsg{
		Version:    int(in.Version),
		Operate:    int(in.Operate),
		Method:     in.Method,
		RoomId:     in.RoomId,
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Extend:     in.Extend,
		Body: types.WriteMsgBody{
			Version:      int(in.Version),
			Operate:      int(in.Operate),
			Method:       in.Method,
			ResponseTime: time.Now().Format("2006-01-02 15:04:05"),
			Event: types.Event{
				Params: "",
				Data:   "",
			},
		},
	}
	switch in.Method {
	case consts.METHOD_ENTER_MSG:
		params := &socket.EventParamsLogin{}
		if err = in.Event.Params.UnmarshalTo(params); err != nil {
			result.Code = commons.SOCKET_BROADCAST_NORMAL_SINGLE
			goto resultHan
		}
		writeMsg.Body.Event.Data = types.DataByEnter{
			RoomId:   in.RoomId,
			ClientId: params.ClientId,
			UserId:   params.UserId,
			UserName: params.UserName,
		}
	case consts.METHOD_NORMAL_MSG:
		params := &socket.EventParamsNormal{}
		if err = in.Event.Params.UnmarshalTo(params); err != nil {
			result.Code = commons.SOCKET_BROADCAST_NORMAL_GROUP
			goto resultHan
		}
		writeMsg.Body.Event.Data = types.DataByNormal{
			RoomId:        in.RoomId,
			FromUserId:    in.FromUserId,
			FrommUserName: in.FromUserName,
			Message:       params.Message,
		}
	}
	if body, err = jsonx.Marshal(writeMsg); err != nil {
		result.Code = commons.SOCKET_BROADCAST_NORMAL
		goto resultHan
	}
	// 发布消息，将消息都分发给订阅了的消费者,群聊|私聊都是同一个发布者，这里不做区分
	if err = dbs.RedisClient.Publish(l.ctx, commons.PUB_SUB_SOCKET_MESSAGE_NORMAL_CHANNEL_KEY, string(body)).Err(); err != nil {
		result.Code = commons.SOCKET_BROADCAST_NORMAL_GROUP
		goto resultHan
	}
resultHan:
	return &socket.Result{}, nil
}
