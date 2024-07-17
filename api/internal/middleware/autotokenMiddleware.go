package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"store-chat/api/internal/types"
	"store-chat/model/mysqls"
	"store-chat/tools/commons"
)

type AutoTokenMiddleware struct {
}

func NewAutoTokenMiddleware() *AutoTokenMiddleware {
	return &AutoTokenMiddleware{}
}

func (m *AutoTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commons.SetHeader(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		result := types.NewResponseJson()
		autoToken := r.Header.Get("autoToken")
		// API接口没有使用rpc做业务处理，有需要自行更改
		if autoToken == "" {
			result.Code, result.Message = commons.GetCodeMessage(commons.RESPONSE_UNAUTHORIZED)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}
		user, err := mysqls.NewUserMgr().GetUser(mysqls.Users{
			Token:  autoToken,
			Status: 1,
		})
		if err != nil {
			result.Code, result.Message = commons.GetCodeMessage(commons.RESPONSE_FAIL)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}
		if user.UserID == 0 {
			result.Code, result.Message = commons.GetCodeMessage(commons.USER_INFO_FAIL)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		next(w, r)
	}
}
