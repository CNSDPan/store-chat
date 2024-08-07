package user

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"store-chat/api/internal/svc"
	"store-chat/api/internal/types"
	"store-chat/dbs"
	"store-chat/model/mysqls"
	"store-chat/tools/commons"
	"store-chat/tools/tools"
	"strconv"
	"time"

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
	l.Logger.Errorf("%#v req.AutoToken:%s", user, req.AutoToken)
	if user.UserID == 0 {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.USER_INFO_FAIL)
		return
	}
	if user.Authorization, err = dbs.RedisClient.Get(l.ctx, commons.USER_AUTHORIZATION_KEY+strconv.FormatInt(user.UserID, 10)).Result(); err != nil && err != redis.Nil {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.USER_TOKEN_GET)
		return
	} else if user.Authorization != "" && req.Source != "goTest" {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.USER_LOGINED)
		return
	}

	if user.Authorization != "" {
		goto END
	}
	if user.Authorization, err = tools.JWTCreateAuthorizationBy32(jwt.MapClaims{
		"iss": l.svcCtx.Config.Name,
		"sub": user.UserID,
		"aud": user.Name,
		"exp": time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.AccessExpire)).UnixMilli(),
	}); err != nil {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.USER_TOKEN_CREATE)
		return
	}
	if err = dbs.RedisClient.SetNX(l.ctx, commons.USER_AUTHORIZATION_KEY+strconv.FormatInt(user.UserID, 10), user.Authorization, time.Duration(l.svcCtx.Config.AccessExpire)*time.Second).Err(); err != nil {
		resp.Code, resp.Message = commons.GetCodeMessage(commons.USER_TOKEN_CREATE)
		return
	}
END:
	resp.Data = user
	return
}
