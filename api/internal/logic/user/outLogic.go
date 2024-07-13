package user

import (
	"context"

	"store-chat/api/internal/svc"
	"store-chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OutLogic {
	return &OutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OutLogic) Out(req *types.ReqOut) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
