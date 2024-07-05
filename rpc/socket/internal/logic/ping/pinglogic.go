package pinglogic

import (
	"context"

	"store-chat/rpc/socket/internal/svc"
	"store-chat/rpc/socket/pb/socket"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *socket.ReqPing) (*socket.ResPing, error) {
	// todo: add your logic here and delete this line

	return &socket.ResPing{}, nil
}
