package commons

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

const (
	// PUB_SUB_SOCKET_MESSAGE_LOGIN_CHANNEL_KEY  发布|订阅频道KEY  socket连接进入消息
	PUB_SUB_SOCKET_MESSAGE_LOGIN_CHANNEL_KEY = "socket:message:login"
	// PUB_SUB_SOCKET_MESSAGE_NORMAL_CHANNEL_KEY 发布|订阅频道KEY socket普通消息
	PUB_SUB_SOCKET_MESSAGE_NORMAL_CHANNEL_KEY = "socket:message:normal"
	// PUB_SUB_SOCKET_MESSAGE_OUT_CHANNEL_KEY 发布|订阅频道KEY socket退出消息
	PUB_SUB_SOCKET_MESSAGE_OUT_CHANNEL_KEY = "socket:message:out"

	// SOCKET_CLIENTS_KEY 每次新建连接token校验通过后,生成一个标识,用于判断当前房间的用户是否已经连上聊天了
	SOCKET_CLIENTS_KEY = "socket:clients:"
)

// GetSocketClientsKey
// @Desc：获取每次新建socket连接的标识,用于判断当前房间的用户是否已经连上聊天了
// @param：roodId
// @param：userId
// @return：key
func GetSocketClientsKey(roodId int64, userId int64) (key string) {
	m := md5.New()
	m.Write([]byte(strconv.FormatInt(roodId, 10) + "_" + strconv.FormatInt(userId, 10)))
	key = SOCKET_CLIENTS_KEY + hex.EncodeToString(m.Sum(nil))
	return key
}
