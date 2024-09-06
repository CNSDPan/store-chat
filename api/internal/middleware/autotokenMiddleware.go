package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"store-chat/api/internal/types"
	"store-chat/model/mysqls"
	"store-chat/tools/commons"
	"strconv"
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
		statusStr := r.Header.Get("status")
		if statusStr == "" {
			statusStr = "1"
		}
		status, _ := strconv.Atoi(statusStr)
		// API接口没有使用rpc做业务处理，有需要自行更改gu
		if autoToken == "" {
			result.Code, result.Message = commons.GetCodeMessage(commons.RESPONSE_UNAUTHORIZED)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		user, err := mysqls.NewUserMgr().GetUser(mysqls.Users{
			Token:  autoToken,
			Status: int8(status),
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
