package types

// ReceiveMsg 接收socket消息结构体
type ReceiveMsg struct {
	Version      int    `json:"version"`             // 用于区分业务版本号
	Operate      int    `json:"operate"`             // 操作
	Method       string `json:"method"`              // 事件
	AutoToken    string `json:"autoToken"`           // token
	RoomId       int64  `json:"roomId,string"`       // 房间
	FromClientId int64  `json:"fromClientId,string"` // 消息发送人client
	FromUserId   int64  `json:"fromUserId"`          // 消息发送人ID
	FromUserName string `json:"fromUserName"`        // 发送人client
	ToClientId   int64  `json:"toClientId,string"`   // 消息发送指定人
	ToUserId     int64  `json:"toUserId"`            // 消息接收人
	ToUserName   string `json:"toUserName"`          // 指定人
	Event        Event  `json:"event"`               // 请求&响应参数
}

// WriteMsg 广播消息结构体
type WriteMsg struct {
	Version    int          `json:"version"`
	Operate    int          `json:"operate"`
	Method     string       `json:"method"`
	AutoToken  string       `json:"autoToken"`
	RoomId     int64        `json:"RoomId,string"`
	FromUserId int64        `json:"fromClientId,string"`
	ToUserId   int64        `json:"toUserId,string"`
	Extend     string       `json:"extend"`
	Body       WriteMsgBody `json:"body"`
}
type WriteMsgBody struct {
	Version      int    `json:"version"`
	Operate      int    `json:"operate"`
	Method       string `json:"method"`
	ResponseTime string `json:"responseTime"`
	Event        Event  `json:"event"`
}

/******************Event 请求&响应结构*********************/

// Event 请求&响应结构
type Event struct {
	Params interface{} `json:"params"` // 请求参数
	Data   interface{} `json:"data"`   // 响应参数
}

// DataByEnter 进入房间响应Data结构
type DataByEnter struct {
	RoomId   int64  `json:"roomId,string"`   // 房间
	ClientId int64  `json:"clientId,string"` // clientId
	UserId   int64  `json:"userId,string"`   // 用户id
	UserName string `json:"userName"`        // 发送人
}

// DataByNormal 进入普通消息响应Data结构
type DataByNormal struct {
	RoomId        int64  `json:"roomId,string"`     // 房间
	FromUserId    int64  `json:"fromUserId,string"` // 消息来源用户
	FrommUserName string `json:"fromUserName"`      // 消息来源用户
	Message       string `json:"message"`           // 消息内容
}

/******************Event 请求&响应结构*********************/
