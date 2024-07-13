package user

import (
	"context"
	"store-chat/model/mysqls"
	"store-chat/tools/commons"

	"store-chat/api/internal/svc"
	"store-chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ReqList) (resp *types.Response, err error) {
	resp = types.NewResponseJson()
	db := mysqls.NewUserMgr()
	userPage, e := db.SelectPage(
		mysqls.NewPage(int64(req.Limit), int64(req.Offset)),
		db.WithStatus(mysqls.USER_STATUS_1),
	)
	if e != nil {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.RESPONSE_FAIL)
		l.Logger.Errorf("%s 查询失败：%s", l.svcCtx.Config.ServerName, e.Error())
		return
	}
	resp.Data = types.DataPageList{
		Total:   userPage.GetTotal(),
		Page:    userPage.GetPages(),
		Limit:   userPage.GetSize(),
		Offset:  userPage.Offset(),
		Current: userPage.GetCurrent(),
		Rows:    userPage.GetRecords(),
	}
	return
}
