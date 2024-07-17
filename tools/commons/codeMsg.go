package commons

import (
	"fmt"
)

var codeMessage = map[string]string{
	RESPONSE_SUCCESS:           "success",
	RESPONSE_FAIL:              "服务器内部错误",
	RESPONSE_NOT_FOUND:         "请求资源不存在",
	RESPONSE_UNAUTHORIZED:      "缺少身份认证",
	RESPONSE_NOT_CODE:          "无定义code码",
	RESPONSE_REQUEST_TIME_FAIL: "缺少请求日期",
	RESPONSE_TOKEN_FAIL:        "无效token",
	RESPONSE_APPID_FAIL:        "无效APPID",
	RESPONSE_SECRET_FAIL:       "无效secret",
	RESPONSE_SIGN_FAIL:         "无效sign",

	/***************************************/
	USER_INFO_FAIL:    "用户信息不存在",
	USER_ID_FAIL:      "用户ID不存在|错误",
	USER_TOKEN_FAIL:   "用户Token不存在|错误",
	USER_TOKEN_CREATE: "用户Token创建错误",
	USER_TOKEN_GET:    "用户Token获取错误",
	USER_LOGINED:      "用户已登录",

	/***************************************/
	SOCKET_BROADCAST_LOGINED:       "socket已连接",
	SOCKET_BROADCAST_LOGIN:         "socket连接错误",
	SOCKET_BROADCAST_OUT:           "socket关闭错误",
	SOCKET_BROADCAST_NORMAL:        "socket广播消息错误",
	SOCKET_BROADCAST_NORMAL_SINGLE: "socket广播消息错误：私聊消息",
	SOCKET_BROADCAST_NORMAL_GROUP:  "socket广播消息错误：群聊消息",
}

// ReturnOverCodeMessage
// @Desc：返回所有codeMsg
// @return：map[string]string
func ReturnOverCodeMessage() map[string]string {
	return codeMessage
}

// GetCodeMessage
// @Desc：获取code码对应message内容
// @param：code
// @return：string,string
func GetCodeMessage(code string) (string, string) {
	var (
		codeMsg map[string]string
		message string
		ok      bool
	)
	codeMsg = ReturnOverCodeMessage()
	if message, ok = codeMsg[code]; !ok {
		message = fmt.Sprintf("code: %s ,%s", code, codeMessage[RESPONSE_NOT_CODE])
		code = RESPONSE_NOT_CODE
	}
	return code, message
}
