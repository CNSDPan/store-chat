package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"store-chat/api/internal/logic/user"
	"store-chat/api/internal/svc"
	"store-chat/api/internal/types"
)

func OutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqOut
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewOutLogic(r.Context(), svcCtx)
		resp, err := l.Out(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
