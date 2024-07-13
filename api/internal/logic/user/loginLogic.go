package user

import (
	"context"
	"store-chat/api/internal/svc"
	"store-chat/api/internal/types"
	"store-chat/model/mysqls"
	"store-chat/tools/commons"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.ReqLogin) (resp *types.Response, err error) {
	resp = types.NewResponseJson()
	user, e := mysqls.NewUserMgr().GetUser(mysqls.Users{Token: req.AutoToken})
	if e != nil {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.RESPONSE_FAIL)
		l.Logger.Errorf("%s 查询失败：%s", l.svcCtx.Config.ServerName, e.Error())
		return
	}
	if user.UserID == 0 {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.USER_INFO_FAIL)
		return
	}
	resp.Data = user
	return
}
