package types

import (
	"store-chat/tools/commons"
	"time"
)

func NewResponseJson() *Response {
	code, msg := commons.GetCodeMessage(commons.RESPONSE_SUCCESS)
	return &Response{
		Modult:       "",
		Code:         code,
		Message:      msg,
		ResponseTime: time.Now().Format("2006-01-02 15:04:05"),
		Data:         make(map[string]interface{}),
	}
}
