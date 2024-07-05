package broadcastlogic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/anypb"
	"store-chat/dbs"
	"store-chat/model/mysqls"
	"store-chat/rpc/socket/internal/svc"
	"store-chat/rpc/socket/pb/socket"
	"store-chat/tools/commons"
)

type BroadcastLoginLogic struct {
	module string
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBroadcastLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastLoginLogic {
	return &BroadcastLoginLogic{
		module: svcCtx.Config.ServerName + "BroadcastLogin",
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BroadcastLoginLogic) BroadcastLogin(in *socket.ReqBroadcastMsg) (result *socket.Result, rpcErr error) {
	var (
		err  error
		user mysqls.UserApi
		//ok         bool
		resultData = &socket.EventDataLogin{}
		hasKey     int64
	)
	result = &socket.Result{
		Module: l.module,
		ErrMsg: "",
		Code:   commons.RESPONSE_SUCCESS,
		Msg:    "",
		Data:   nil,
	}
	defer func() {
		result.Code, result.Msg = commons.GetCodeMessage(result.Code)
		if err != nil && err != redis.Nil {
			result.ErrMsg = err.Error()
			l.Logger.Errorf("%s Broadcast ope:%d fail:%s", result.Module, in.Operate, err.Error())
		}
	}()
	if user, err = mysqls.NewUserMgr().GetUser(mysqls.Users{Token: in.AutoToken}); err != nil {
		l.Logger.Errorf("%s 查询用户 fail:%s", result.Module, err.Error())
		result.Code = commons.USER_TOKEN_FAIL
		return result, rpcErr
	}
	if user.UserID == 0 {
		l.Logger.Errorf("%s 无用户 fail:%v", result.Module, err)
		result.Code = commons.USER_INFO_FAIL
		return result, rpcErr
	}
	//if user, ok = tools.UserMap[in.AutoToken]; !ok {
	//	l.Logger.Errorf("%s 无用户 not ok", result.Module)
	//	result.Code = commons.USER_INFO_FAIL
	//	goto resultHan
	//}
	// 判断是否连接过当前房间
	if hasKey, err = dbs.RedisClient.Exists(l.ctx, commons.GetSocketClientsKey(in.RoomId, user.UserID)).Result(); err != nil && err != redis.Nil {
		l.Logger.Errorf("%s 校验redis:socket连接key是否存在;ERR:%v", result.Module, err)
		result.Code = commons.SOCKET_BROADCAST_LOGIN
		return result, rpcErr
	}
	if hasKey > 0 {
		result.Code = commons.SOCKET_BROADCAST_LOGIN
		return result, rpcErr
	}
	if err = dbs.RedisClient.Set(l.ctx, commons.GetSocketClientsKey(in.RoomId, user.UserID), fmt.Sprintf("%d_%d", in.RoomId, user.UserID), -1).Err(); err != nil {
		l.Logger.Errorf("%s 存储redis:socket:client:key;ERR:%v", result.Module, err)
		result.Code = commons.SOCKET_BROADCAST_LOGIN
		return result, rpcErr
	}
	// any类型反射
	resultData.RoomId = in.RoomId
	resultData.UserId = user.UserID
	resultData.UserName = user.Name
	if result.Data, err = anypb.New(resultData); err != nil {
		l.Logger.Errorf("%s Broadcast 发布消息 body json.Marshal fail:%s", result.Module, err.Error())
		result.Code = commons.SOCKET_BROADCAST_LOGIN
	}
	return result, rpcErr
}
