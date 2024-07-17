package commons

const (
	// PUB_SUB_SOCKET_MESSAGE_LOGIN_CHANNEL_KEY  发布|订阅频道KEY  socket连接进入消息
	PUB_SUB_SOCKET_MESSAGE_LOGIN_CHANNEL_KEY = "socket:message:login"
	// PUB_SUB_SOCKET_MESSAGE_NORMAL_CHANNEL_KEY 发布|订阅频道KEY socket普通消息
	PUB_SUB_SOCKET_MESSAGE_NORMAL_CHANNEL_KEY = "socket:message:normal"
	// PUB_SUB_SOCKET_MESSAGE_OUT_CHANNEL_KEY 发布|订阅频道KEY socket退出消息
	PUB_SUB_SOCKET_MESSAGE_OUT_CHANNEL_KEY = "socket:message:out"
)

const (
	// client的用户key，分布式下用于校验是否已经连接了
	SOCKET_CHAT_KEY = "socket:chat:"
	// client的用户key的有效时长，务必比pong长
	SOCKET_CHAT_KEY_EXPIRE_SECOND = 60
)
