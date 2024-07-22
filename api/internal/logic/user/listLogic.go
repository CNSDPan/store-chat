package user

import (
	"context"
	"github.com/redis/go-redis/v9"

	//"github.com/redis/go-redis/v9/"
	"store-chat/dbs"
	"store-chat/model/mysqls"
	"store-chat/tools/commons"
	"strconv"

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
	userPage, err := db.SelectPage(
		mysqls.NewPage(int64(req.Limit), int64(req.Offset)),
		db.WithStatus(mysqls.USER_STATUS_1),
	)
	if err != nil {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.RESPONSE_FAIL)
		l.Logger.Errorf("%s 查询失败：%s", l.svcCtx.Config.ServerName, err.Error())
		return
	}
	rows := make([]mysqls.UserApi, 0)
	for _, user := range userPage.GetRecords().([]mysqls.UserApi) {
		if err = dbs.RedisClient.Get(l.ctx, commons.USER_AUTHORIZATION_KEY+strconv.FormatInt(user.UserID, 10)).Err(); err == nil {
			user.WsConn = mysqls.WS_CONN_ON_LINE
		} else if err == redis.Nil {
			user.WsConn = mysqls.WS_CONN_LEAVE
		} else {
			l.Logger.Errorf("%s redis 查询失败：%s", l.svcCtx.Config.ServerName, err.Error())
			return
		}
		rows = append(rows, user)
		err = nil
	}

	resp.Data = types.DataPageList{
		Total:   userPage.GetTotal(),
		Page:    userPage.GetPages(),
		Limit:   userPage.GetSize(),
		Offset:  userPage.Offset(),
		Current: userPage.GetCurrent(),
		Rows:    rows,
	}
	return
}
