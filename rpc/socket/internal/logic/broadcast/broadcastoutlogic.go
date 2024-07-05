package broadcastlogic

import (
	"context"
	"github.com/redis/go-redis/v9"
	"store-chat/dbs"
	"store-chat/rpc/socket/internal/svc"
	"store-chat/rpc/socket/pb/socket"
	"store-chat/tools/commons"

	"github.com/zeromicro/go-zero/core/logx"
)

type BroadcastOutLogic struct {
	module string
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBroadcastOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastOutLogic {
	return &BroadcastOutLogic{
		module: svcCtx.Config.ServerName + "BroadcastNormal",
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BroadcastOutLogic) BroadcastOut(in *socket.ReqBroadcastMsg) (result *socket.Result, rpcErr error) {
	var (
		err error
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
	if err = dbs.RedisClient.Del(l.ctx, commons.GetSocketClientsKey(in.RoomId, in.FromUserId)).Err(); err != nil && err != redis.Nil {
		l.Logger.Errorf("%s 移除redis:socket:client:key;ERR:%v", result.Module, err)
		result.Code = commons.SOCKET_BROADCAST_OUT
		return result, rpcErr
	}
	return &socket.Result{}, nil
}
