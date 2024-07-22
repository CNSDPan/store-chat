package broadcastlogic

import (
	"context"
	"github.com/redis/go-redis/v9"
	"store-chat/dbs"
	"store-chat/rpc/socket/internal/svc"
	"store-chat/rpc/socket/pb/socket"
	"store-chat/tools/commons"
	"strconv"

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
		chatKey string
		err     error
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
	// 判断当前缓存authToken是否一致;一致时清除
	chatKey, err = dbs.RedisClient.Get(l.ctx, commons.USER_AUTHORIZATION_KEY+strconv.FormatInt(in.FromUserId, 10)).Result()
	if err != nil && err != redis.Nil {
		l.Logger.Errorf("%s 用户token获取 fail:%v", result.Module, err)
		result.Code = commons.USER_TOKEN_GET
		return result, rpcErr
	} else if chatKey == in.AuthToken {
		if _, err = dbs.RedisClient.Del(l.ctx, commons.USER_AUTHORIZATION_KEY+strconv.FormatInt(in.FromUserId, 10)).Result(); err != nil {
			l.Logger.Errorf("%s 用户token获取 fail:%v", result.Module, err)
			result.Code = commons.USER_TOKEN_GET
			return result, rpcErr
		}
	}

	return &socket.Result{}, nil
}
