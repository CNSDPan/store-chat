package broadcastlogic

import (
	"context"
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
	return &socket.Result{}, nil
}
