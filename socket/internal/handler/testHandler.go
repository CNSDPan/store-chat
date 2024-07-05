package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"store-chat/socket/internal/logic"
	"store-chat/socket/internal/svc"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTestLogic(r.Context(), svcCtx)
		err := l.Test()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
