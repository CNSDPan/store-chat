package socket

import (
	"fmt"
	"store-chat/model/mysqls"
	"store-chat/tools/consts"
	"store-chat/tools/tools"
	"store-chat/tools/types"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	socketUrl = "ws://192.168.33.10:6991/ws"
	//socketUrl = "ws://roha:8888/ws"
)

var TClient *TestClient

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

func NewConnect(authToken string, roomId int64) error {
	tClient, err := New(socketUrl)
	if err != nil {
		return err
	}
	err = tClient.Auth(authToken, roomId)
	if err != nil {
		return err
	}
	TClient = tClient
	TClient.Send()
	TClient.Read()
	return nil
}

func Read() {
	var (
		roomIdStr   string
		clientIdStr string
		userIdStr   string
		qa          = QA{
			roomId:       0,
			fromUserId:   0,
			fromUserName: "",
		}
	)
	for {
		select {
		case e := <-TClient.RevMsgFail:
			fmt.Println(e)
		case m := <-TClient.RevMsgChan:
			switch m.Method {
			case consts.METHOD_ENTER_MSG:
				if data, ok := m.Event.Data.(map[string]interface{}); !ok {
					fmt.Println("m.Event.Data typeOf types.DataByEnter not ok\n ")
				} else {
					//fmt.Printf("enter:%+v\n", data)
					clientIdStr = data["clientId"].(string)
					userIdStr = data["userId"].(string)
					TClient.ClientId, _ = strconv.ParseInt(clientIdStr, 10, 64)
					TClient.UserId, _ = strconv.ParseInt(userIdStr, 10, 64)
					TClient.UserName = data["userName"].(string)
				}
			case consts.METHOD_NORMAL_MSG:
				if data, ok := m.Event.Data.(map[string]interface{}); !ok {
					fmt.Println("m.Event.Data typeOf types.DataByNormal not ok\n ")
				} else {
					if m.Operate == consts.OPERATE_SINGLE_MSG {
						fmt.Printf(m.ResponseTime+":私聊消息：\n     %s\n", data["message"])
					} else if m.Operate == consts.OPERATE_GROUP_MSG {
						fmt.Printf(m.ResponseTime+":广播消息：\n     %s\n", data["message"])
					}
					roomIdStr, _ = data["roomId"].(string)
					userIdStr = data["fromUserId"].(string)
					//userIdStr = "1"
					qa.roomId, _ = strconv.ParseInt(roomIdStr, 10, 64)
					qa.fromUserId, _ = strconv.ParseInt(userIdStr, 10, 64)
					qa.fromUserName = data["fromUserName"].(string)
					qa.message = data["message"].(string)
					TClient.QAChan <- qa
				}
			}

		}
	}
}

func Send(operate int, roomId int64, toUserId int64, msg string, after time.Duration, sendNum int, autoToken string) {
	send := types.ReceiveMsg{
		Version:      1,
		Operate:      operate,
		Method:       consts.METHOD_NORMAL_MSG,
		AutoToken:    autoToken,
		RoomId:       roomId,
		FromClientId: TClient.ClientId,
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
				send.Event.Params = TClient.UserName + ":" + msg
				TClient.SendMsgChan <- send
			}
		}
	}()
}

// 模拟的账号可再tools/tools/dbTest UserMap得到
func TestUser1Room1(t *testing.T) {
	autoToken := "2gDGQwDxsrX0UG8yRbophdHxHqD"
	if err := NewConnect(autoToken, 1); err != nil {
		panic("TestUser1Room1 连接失败:" + err.Error())
	}
	// 群聊
	Send(consts.OPERATE_GROUP_MSG, 1, 0, "说:爸爸的爸爸叫什么", 5*time.Second, 0, autoToken)
	Read()

}
func TestUser1Room2(t *testing.T) {
	autoToken := "2gDGQwDxsrX0UG8yRbophdHxHqD"
	if err := NewConnect(autoToken, 2); err != nil {
		panic("TestUser1Room2 连接失败:" + err.Error())
	}
	// 群聊
	Send(consts.OPERATE_GROUP_MSG, 2, 0, "妈妈的妈妈叫什么", 4*time.Second, 3, autoToken)
	Read()
}
func TestUser1Single1(t *testing.T) {
	autoToken := "2gDGQwDxsrX0UG8yRbophdHxHqD"
	if err := NewConnect(autoToken, 1); err != nil {
		panic("TestUser1Room2 连接失败:" + err.Error())
	}
	// 私信
	Send(consts.OPERATE_SINGLE_MSG, 1, 0, "我在跟你私聊，群里的人不知道", 4*time.Second, 0, autoToken)
	Read()
}
func TestUser2Room1(t *testing.T) {
	autoToken := "2gDGQugkyFF4MI10hK7WfT3W3Pe"
	if err := NewConnect(autoToken, 1); err != nil {
		panic("TestUser2Room1 连接失败:" + err.Error())
	}
	// 群聊
	go func() {
		for {
			select {
			case msg := <-TClient.QAChan:
				if strings.Contains(msg.message, "爸爸的爸爸叫什么") {
					Send(consts.OPERATE_GROUP_MSG, msg.roomId, 0, "叫爷爷---群聊", 0*time.Second, 1, autoToken)
				}
			}
		}
	}()
	Read()
}
func TestUser2Room2(t *testing.T) {
	autoToken := "2gDGQugkyFF4MI10hK7WfT3W3Pe"
	if err := NewConnect(autoToken, 2); err != nil {
		panic("TestUser2Room1 连接失败:" + err.Error())
	}
	// 群聊
	go func() {
		for {
			select {
			case msg := <-TClient.QAChan:
				if strings.Contains(msg.message, "爸爸的爸爸叫什么") {
					Send(consts.OPERATE_GROUP_MSG, msg.roomId, 0, "叫爷爷---群聊", 0*time.Second, 1, autoToken)
				}
			}
		}
	}()
	Read()
}
func TestUser3Room1(t *testing.T) {
	autoToken := "2gDGQvEugR6Y5riFp2kVLdc7J0O"
	if err := NewConnect(autoToken, 1); err != nil {
		panic("TestUser3Room1 连接失败:" + err.Error())
	}
	// 群里私信
	go func() {
		for {
			select {
			case msg := <-TClient.QAChan:
				if strings.Contains(msg.message, "爸爸的爸爸叫什么") {
					Send(consts.OPERATE_SINGLE_MSG, msg.roomId, msg.fromUserId, "应该叫做爷爷---私信", 0*time.Second, 1, autoToken)
				}
			}
		}
	}()
	Read()
}
func TestUser4Room1(t *testing.T) {
	autoToken := "2gDGQwhqJQczjkCikEvg3StOKSR"
	if err := NewConnect(autoToken, 1); err != nil {
		panic("TestUser4Room1 连接失败:" + err.Error())
	}
	// 群聊
	go func() {
		for {
			select {
			case msg := <-TClient.QAChan:
				if strings.Contains(msg.message, "爸爸的爸爸叫什么") {
					Send(consts.OPERATE_GROUP_MSG, msg.roomId, msg.fromUserId, "不知道叫什么---群聊", 0*time.Second, 1, autoToken)
				}
			}
		}
	}()
	Read()
}
func TestUser5Single1(t *testing.T) {
	autoToken := "2gDGQvpg5xTE3Qn0SIzbyDXpdma"
	if err := NewConnect(autoToken, 1); err != nil {
		panic("TestUser4Room1 连接失败:" + err.Error())
	}
	// 私信

	Send(consts.OPERATE_SINGLE_MSG, 1, 0, "大哥原来是你私信我啊", 5*time.Second, 0, autoToken)
	Read()
}
