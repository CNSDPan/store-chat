package consts

const (
	//OPERATE_SINGLE_MSG 单人聊天操作
	OPERATE_SINGLE_MSG = 2
	// OPERATE_GROUP_MSG 群体聊天操作
	OPERATE_GROUP_MSG = 3
	// OPERATE_CONN_MSG 建立连接操作
	OPERATE_CONN_MSG = 10

	// METHOD_CONN_MSG 建立连接操作 的事件命名
	METHOD_ENTER_MSG = "Enter"
	// METHOD_OUT_MSG 关闭连接操作 的事件命名
	METHOD_OUT_MSG = "Out"
	// METHOD_GROUP_MSG 聊天 的普通消息事件命名
	METHOD_NORMAL_MSG = "Normal"
	// METHOD_SERVER_MSG 获取当前client所在的服务信息
	METHOD_SERVER_MSG = "Server"
	// METHOD_PM_MSG 压测消息推送
	METHOD_PM_MSG = "PM"

	// FIRST_CONN 最新连接
	FIRST_CONN = "firstCONN"
	// REPEAT_CONN 重复连接
	REPEAT_CONN = "repeatConn"
)
