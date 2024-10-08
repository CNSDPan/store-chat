package broadcastlogic

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/anypb"
	"store-chat/dbs"
	"store-chat/model/mysqls"
	"store-chat/rpc/socket/internal/svc"
	"store-chat/rpc/socket/pb/socket"
	"store-chat/tools/commons"
	"store-chat/tools/tools"
	"strconv"
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
		err     error
		chatKey string
		user    mysqls.UserApi
		//ok         bool
		resultData = &socket.EventDataLogin{}
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
			l.Logger.Errorf("%s op:%d fail:%s", result.Module, in.Operate, err.Error())
		}
	}()
	if user, err = mysqls.NewUserMgr().GetUser(mysqls.Users{UserID: in.FromUserId}); err != nil {
		l.Logger.Errorf("%s 查询用户 fail:%s", result.Module, err.Error())
		result.Code = commons.USER_TOKEN_FAIL
		return result, rpcErr
	}
	if user.UserID == 0 {
		l.Logger.Errorf("%s 无用户 fail:%v", result.Module, err)
		result.Code = commons.USER_INFO_FAIL
		return result, rpcErr
	}

	// 模拟用户数据
	//if user, ok = tools.UserMap[in.FromUserId]; !ok {
	//	l.Logger.Errorf("%s 无用户 not ok", result.Module)
	//	result.Code = commons.USER_INFO_FAIL
	//	goto resultHan
	//}

	chatKey, err = dbs.RedisClient.Get(l.ctx, commons.USER_AUTHORIZATION_KEY+strconv.FormatInt(user.UserID, 10)).Result()
	if (err == nil && chatKey != in.AuthToken) || err == redis.Nil {
		l.Logger.Errorf("%s 【%s】用户token不匹配【%s】【%s】【%s】", result.Module, user.Name, in.AuthToken, chatKey, tools.StoreMap[in.RoomId].Name)
		result.Code = commons.USER_TOKEN_FAIL
		return result, rpcErr
	} else if err != nil {
		l.Logger.Errorf("%s 用户token获取 fail:%v", result.Module, err)
		result.Code = commons.USER_TOKEN_GET
		return result, rpcErr
	}
	// any类型反射
	resultData.RoomId = in.RoomId
	resultData.UserId = user.UserID
	resultData.UserName = user.Name
	if result.Data, err = anypb.New(resultData); err != nil {
		l.Logger.Errorf("%s Broadcast 发布消息 body json.Marshal fail:%s", result.Module, err.Error())
		result.Code = commons.SOCKET_BROADCAST_LOGIN
		return result, rpcErr
	}
	return result, rpcErr
}
