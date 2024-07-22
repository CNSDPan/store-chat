package socket

import (
	"fmt"
	"store-chat/dbs"
	"store-chat/model/mysqls"
	"store-chat/tools/consts"
	"store-chat/tools/tools"
	"store-chat/tools/types"
	"strconv"
	"testing"
	"time"
)

const (
	//socketUrl = "ws://192.168.33.10:6991/ws"
	//socketUrl = "ws://192.168.33.10:6992/ws"
	socketUrl = "ws://websocket.cn:6990/ws"

	//socketUrl = "ws://roha:8888/ws"
)

func init() {
	_, _ = dbs.NewRedisClient()
}

var user = new(DefaultUser)

var store = tools.StoreMap

func TestVirtualUser(t *testing.T) {
	var (
		has  = "2gDGQwDxsrX0UG8yRbophdHxHqD"
		not  = "11111"
		user mysqls.UserApi
		ok   bool
	)
	user = tools.UserMap[has]
	fmt.Printf("查找到的账号：%+v \n", user)
	if user, ok = tools.UserMap[not]; !ok {
		fmt.Printf("不存在的账号：%+v \n", user)
	}
}

func TestUser1Room1(t *testing.T) {
	user.InitUserInfo("2gDGQwDxsrX0UG8yRbophdHxHqD")
	user.InitSocket(socketUrl)
	user.Client.Auth(user.AuthToken, 1810940924055547904, user.UserId)
	go func() {
		Read()
	}()
	SendQA()
	select {
	case isClose := <-user.IsClose:
		if isClose == 1 {
			goto END
		}
	}
END:
	fmt.Printf("关闭连接\n")
	return
}

func TestUser2Room1(t *testing.T) {
	user.InitUserInfo("2gDGQugkyFF4MI10hK7WfT3W3Pe")
	user.InitSocket(socketUrl)
	user.Client.Auth(user.AuthToken, 1810940924055547904, user.UserId)
	go func() {
		Read()
	}()
	//SendQA()
	select {
	case isClose := <-user.IsClose:
		if isClose == 1 {
			goto END
		}
	}
END:
	fmt.Printf("关闭连接\n")
	return
}

func Read() {
	var (
		clientIdStr string
		roomIdStr   string
		userIdStr   string
		roomId      int64
		fromUserId  int64
	)
	for {
		select {
		case e := <-user.Client.RevMsgFail:
			_ = user.Client.Conn.Close()
			user.IsClose <- 1
			fmt.Printf("断开连接：%v\n", e)
			goto END
		case m := <-user.Client.RevMsgChan:
			if m.Method == consts.METHOD_ENTER_MSG {
				if data, ok := m.Event.Data.(map[string]interface{}); !ok {
					fmt.Printf("m.Event.Data typeOf types.DataByEnter not ok\n")
				} else {
					clientIdStr = data["clientId"].(string)
					user.Client.ClientId, _ = strconv.ParseInt(clientIdStr, 10, 64)
				}
			} else if m.Method == consts.METHOD_NORMAL_MSG {
				if data, ok := m.Event.Data.(map[string]interface{}); !ok {
					fmt.Printf("m.Event.Data typeOf types.DataByNormal not ok\n")
				} else {
					if m.Operate == consts.OPERATE_SINGLE_MSG {
						fmt.Printf(m.ResponseTime+":私聊消息：\n     %s\n", data["message"])
					} else if m.Operate == consts.OPERATE_GROUP_MSG {
						fmt.Printf(m.ResponseTime+":广播消息：\n     %s\n", data["message"])
					}
					roomIdStr, _ = data["roomId"].(string)
					userIdStr = data["fromUserId"].(string)
					if user.UserName == "蟑螂恶霸" {
						roomId, _ = strconv.ParseInt(roomIdStr, 10, 64)
						fromUserId, _ = strconv.ParseInt(userIdStr, 10, 64)
						user.Client.QAChan <- QA{
							roomId:       roomId,
							fromUserId:   fromUserId,
							fromUserName: data["fromUserName"].(string),
							message:      data["message"].(string),
						}
					}
				}
			}
		}
	}
END:
	return
}

func Send(operate int, roomId int64, toUserId int64, msg string, after time.Duration, sendNum int, autoToken string) {
	send := types.ReceiveMsg{
		Version:      1,
		Operate:      operate,
		Method:       consts.METHOD_NORMAL_MSG,
		AuthToken:    autoToken,
		RoomId:       roomId,
		FromUserId:   user.UserId,
		FromClientId: user.Client.ClientId,
		ToUserId:     toUserId,
		Event:        types.Event{},
	}
	sendIndex := 0
	go func() {
		for {
			// 限制发送次数
			if sendNum > 0 && sendIndex >= sendNum {
				return
			}
			select {
			case <-time.After(after):
				sendIndex++
				send.Event.Params = "自动回复:" + msg
				user.Client.SendMsgChan <- send
			}
		}
	}()
}

func SendQA() {
	var (
		now        time.Time
		weekday    time.Weekday
		sendMsg    = ""
		week       string
		weekdayStr = [...]string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	)
	go func() {
		for {
			select {
			case msg := <-user.Client.QAChan:
				switch msg.message {
				case "我是谁":
					sendMsg = "你是 " + msg.fromUserName
					Send(consts.OPERATE_GROUP_MSG, msg.roomId, 0, sendMsg, 0*time.Second, 1, user.AuthToken)
					//time.Sleep(500 * time.Millisecond)
					//Send(consts.OPERATE_SINGLE_MSG, msg.roomId, msg.fromUserId, "再偷偷私信你~你叫 "+msg.fromUserName, 0*time.Second, 1, user.AuthToken)
				case "当前时间":
					now = time.Now()
					weekday = now.Weekday()
					week = weekdayStr[weekday]
					sendMsg = fmt.Sprintf("私信---今天是%s %s", now.Format("2006-01-02 15:04:05"), week)
					Send(consts.OPERATE_GROUP_MSG, msg.roomId, msg.fromUserId, sendMsg, 0*time.Second, 1, user.AuthToken)
				}
			}
		}
	}()
}

//
//// 模拟的账号可再tools/tools/dbTest UserMap得到
//func TestUser1Room1(t *testing.T) {
//	autoToken := "2gDGQwDxsrX0UG8yRbophdHxHqD"
//	if err := NewConnect(autoToken, 1); err != nil {
//		panic("TestUser1Room1 连接失败:" + err.Error())
//	}
//	// 群聊
//	Send(consts.OPERATE_GROUP_MSG, 1, 0, "说:爸爸的爸爸叫什么", 5*time.Second, 0, autoToken)
//	Read()
//
//}
//func TestUser1Room2(t *testing.T) {
//	autoToken := "2gDGQwDxsrX0UG8yRbophdHxHqD"
//	if err := NewConnect(autoToken, 2); err != nil {
//		panic("TestUser1Room2 连接失败:" + err.Error())
//	}
//	// 群聊
//	Send(consts.OPERATE_GROUP_MSG, 2, 0, "妈妈的妈妈叫什么", 4*time.Second, 3, autoToken)
//	Read()
//}
//func TestUser1Single1(t *testing.T) {
//	autoToken := "2gDGQwDxsrX0UG8yRbophdHxHqD"
//	if err := NewConnect(autoToken, 1); err != nil {
//		panic("TestUser1Room2 连接失败:" + err.Error())
//	}
//	// 私信
//	Send(consts.OPERATE_SINGLE_MSG, 1, 0, "我在跟你私聊，群里的人不知道", 4*time.Second, 0, autoToken)
//	Read()
//}
//func TestUser2Room1(t *testing.T) {
//	autoToken := "2gDGQugkyFF4MI10hK7WfT3W3Pe"
//	if err := NewConnect(autoToken, 1); err != nil {
//		panic("TestUser2Room1 连接失败:" + err.Error())
//	}
//	// 群聊
//	go func() {
//		for {
//			select {
//			case msg := <-TClient.QAChan:
//				if strings.Contains(msg.message, "爸爸的爸爸叫什么") {
//					Send(consts.OPERATE_GROUP_MSG, msg.roomId, 0, "叫爷爷---群聊", 0*time.Second, 1, autoToken)
//				}
//			}
//		}
//	}()
//	Read()
//}
//func TestUser2Room2(t *testing.T) {
//	autoToken := "2gDGQugkyFF4MI10hK7WfT3W3Pe"
//	if err := NewConnect(autoToken, 2); err != nil {
//		panic("TestUser2Room1 连接失败:" + err.Error())
//	}
//	// 群聊
//	go func() {
//		for {
//			select {
//			case msg := <-TClient.QAChan:
//				if strings.Contains(msg.message, "爸爸的爸爸叫什么") {
//					Send(consts.OPERATE_GROUP_MSG, msg.roomId, 0, "叫爷爷---群聊", 0*time.Second, 1, autoToken)
//				}
//			}
//		}
//	}()
//	Read()
//}
//func TestUser3Room1(t *testing.T) {
//	autoToken := "2gDGQvEugR6Y5riFp2kVLdc7J0O"
//	if err := NewConnect(autoToken, 1); err != nil {
//		panic("TestUser3Room1 连接失败:" + err.Error())
//	}
//	// 群里私信
//	go func() {
//		for {
//			select {
//			case msg := <-TClient.QAChan:
//				if strings.Contains(msg.message, "爸爸的爸爸叫什么") {
//					Send(consts.OPERATE_SINGLE_MSG, msg.roomId, msg.fromUserId, "应该叫做爷爷---私信", 0*time.Second, 1, autoToken)
//				}
//			}
//		}
//	}()
//	Read()
//}
//func TestUser4Room1(t *testing.T) {
//	autoToken := "2gDGQwhqJQczjkCikEvg3StOKSR"
//	if err := NewConnect(autoToken, 1); err != nil {
//		panic("TestUser4Room1 连接失败:" + err.Error())
//	}
//	// 群聊
//	go func() {
//		for {
//			select {
//			case msg := <-TClient.QAChan:
//				if strings.Contains(msg.message, "爸爸的爸爸叫什么") {
//					Send(consts.OPERATE_GROUP_MSG, msg.roomId, msg.fromUserId, "不知道叫什么---群聊", 0*time.Second, 1, autoToken)
//				}
//			}
//		}
//	}()
//	Read()
//}
//func TestUser5Single1(t *testing.T) {
//	autoToken := "2gDGQvpg5xTE3Qn0SIzbyDXpdma"
//	if err := NewConnect(autoToken, 1); err != nil {
//		panic("TestUser4Room1 连接失败:" + err.Error())
//	}
//	// 私信
//
//	Send(consts.OPERATE_SINGLE_MSG, 1, 0, "大哥原来是你私信我啊", 5*time.Second, 0, autoToken)
//	Read()
//}
