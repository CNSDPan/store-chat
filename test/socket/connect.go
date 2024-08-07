package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"log"
	"net/http"
	"store-chat/tools/commons"
	"store-chat/tools/consts"
	"store-chat/tools/types"
	"strconv"
	"strings"
	"time"
)

var apiUrl = "http://192.168.33.10:7000"
var loginUrl = "/api/user/login"

type UserApi interface {
	InitUserInfo(autoToken string)
	InitSocket(url string)
}

type DefaultUser struct {
	UserId    int64
	UserName  string
	AuthToken string
	Client    *TestClient
	IsClose   chan int
	Log       logx.Logger
}

type ApiResponse struct {
	Modult       string      `json:"modult"`
	Code         string      `json:"code"`
	Message      string      `json:"msg"`
	ResponseTime string      `json:"responseTime"`
	Data         interface{} `json:"data"`
}

type TestClient struct {
	Conn        *websocket.Conn
	Timeout     int
	ClientId    int64
	SendMsgChan chan types.ReceiveMsg
	RevMsgChan  chan types.WriteMsgBody
	RevMsgFail  chan string
	QAChan      chan QA
	RoomName    string
}

type QA struct {
	roomId       int64
	fromUserId   int64
	fromUserName string
	message      string
}

func (u *DefaultUser) InitUserInfo(autoToken string) {
	client := http.Client{}
	req := map[string]interface{}{
		"version":     "1",
		"requestTime": time.Now().UnixMilli(),
		"source":      "goTest",
	}
	b, _ := jsonx.Marshal(req)
	request, err := http.NewRequest(http.MethodPost, apiUrl+loginUrl, strings.NewReader(string(b)))
	if err != nil {
		panic("http错误：" + err.Error())
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Autotoken", autoToken)
	request.Header.Add("Authorization", "")
	request.Header.Add("Status", "2")

	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		panic("请求登录:" + err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Fatalf("status code:%d error: %s body: %s", resp.StatusCode, resp.Status, string(body))
	}
	result := ApiResponse{}
	if err = jsonx.Unmarshal(body, &result); err != nil {
		panic("result结构:" + err.Error())
	}
	if result.Code != commons.RESPONSE_SUCCESS {
		log.Fatalf("登录失败：" + result.Message)
	}
	userApi := result.Data.(map[string]interface{})
	u.UserId, _ = strconv.ParseInt(userApi["userId"].(string), 10, 64)
	u.UserName = userApi["name"].(string)
	u.AuthToken = userApi["authorization"].(string)
	u.IsClose = make(chan int)
}

func (u *DefaultUser) InitSocket(url string, roomName string) {
	if u.AuthToken == "" {
		panic("token为空，不进行InitSocket\n")
	}
	conn, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic(fmt.Sprintf("拨号失败:%v fail:%s", res, err.Error()))
	}
	u.Client = &TestClient{
		Conn:        conn,
		SendMsgChan: make(chan types.ReceiveMsg, 100),
		RevMsgChan:  make(chan types.WriteMsgBody, 100),
		RevMsgFail:  make(chan string, 100),
		QAChan:      make(chan QA, 100),
		Timeout:     30,
		RoomName:    "",
	}
}

func (t *TestClient) Auth(authToken string, roomId int64, userId int64) {
	msg := types.ReceiveMsg{
		Version:    1,
		Operate:    10,
		Method:     consts.METHOD_ENTER_MSG,
		AuthToken:  authToken,
		RoomId:     roomId,
		FromUserId: userId,
		Event:      types.Event{},
	}
	b, err := jsonx.Marshal(msg)
	if err != nil {
		panic(fmt.Sprintf("连接房间：jsonx.Marshal fail %s", err.Error()))
	}
	err = t.Conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		panic(fmt.Sprintf("连接房间：WriteMessage fail %s", err.Error()))
	}
	t.Read()
	t.Send()
}
func (t *TestClient) Read() {
	go func() {
		var (
			err error
			b   []byte
		)
		for {
			_, b, err = t.Conn.ReadMessage()
			if err != nil {
				t.RevMsgFail <- fmt.Sprintf("读取失败：%s\n " + err.Error())
				continue
			}
			msg := types.WriteMsgBody{}
			if err = jsonx.Unmarshal(b, &msg); err != nil {
				t.RevMsgFail <- fmt.Sprintf("读取失败：jsonx.Unmarshal %s\n " + err.Error())
				continue
			}
			fmt.Printf("msg:%v %s\n", msg, t.RoomName)
			//fmt.Printf("%s 管道未读条数：%d\n", time.Now().Format("2006-01-02 15:04:05"), len(t.RevMsgChan))
			t.RevMsgChan <- msg
		}
	}()
}
func (t *TestClient) Send() {
	go func() {
		for {
			r := <-t.SendMsgChan
			b, err := jsonx.Marshal(r)
			if err != nil {
				log.Println("jsonx.Marshal fail:", err.Error())
				continue
			}
			err = t.Conn.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("send t.Conn.WriteMessage fail:", err.Error())
				continue
			}
		}
	}()
}

// Operator
// @Desc：
// @param：url
func (u *DefaultUser) Operator(idx uint32, url string, roomId int64, roomName string) {
	go func() {
		u.InitSocket(url, roomName)
		u.Client.Auth(u.AuthToken, roomId, u.UserId)
		u.Read()
		u.SendQA()
		select {
		case isClose := <-u.IsClose:
			if isClose == 1 {
				goto END
			}
		}
	END:
		u.Log.Errorf("【%s】房【%s】断开连接", roomName)
		return
	}()
}

func (u *DefaultUser) Read() {
	go func() {
		var (
			clientIdStr string
			roomIdStr   string
			userIdStr   string
			roomId      int64
			fromUserId  int64
		)
		for {
			select {
			case e := <-u.Client.RevMsgFail:
				_ = u.Client.Conn.Close()
				u.IsClose <- 1
				fmt.Printf("断开连接：%v\n", e)
				goto END
			case m := <-u.Client.RevMsgChan:
				if m.Method == consts.METHOD_ENTER_MSG {
					if data, ok := m.Event.Data.(map[string]interface{}); !ok {
						fmt.Printf("m.Event.Data typeOf types.DataByEnter not ok\n")
					} else {
						clientIdStr = data["clientId"].(string)
						u.Client.ClientId, _ = strconv.ParseInt(clientIdStr, 10, 64)
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
						if u.UserName == "蜻蜓队长(管理员)" {
							roomId, _ = strconv.ParseInt(roomIdStr, 10, 64)
							fromUserId, _ = strconv.ParseInt(userIdStr, 10, 64)
							u.Client.QAChan <- QA{
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
	}()
}

func (u *DefaultUser) Send(operate int, roomId int64, toUserId int64, msg string, after time.Duration, sendNum int, autoToken string) {
	send := types.ReceiveMsg{
		Version:      1,
		Operate:      operate,
		Method:       consts.METHOD_NORMAL_MSG,
		AuthToken:    autoToken,
		RoomId:       roomId,
		FromUserId:   u.UserId,
		FromClientId: u.Client.ClientId,
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
				u.Client.SendMsgChan <- send
			}
		}
	}()
}

func (u *DefaultUser) SendQA() {
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
			case msg := <-u.Client.QAChan:
				switch msg.message {
				case "我是谁":
					sendMsg = "你是 " + msg.fromUserName
					u.Send(consts.OPERATE_GROUP_MSG, msg.roomId, 0, sendMsg, 0*time.Second, 1, u.AuthToken)
					//time.Sleep(500 * time.Millisecond)
					//Send(consts.OPERATE_SINGLE_MSG, msg.roomId, msg.fromUserId, "再偷偷私信你~你叫 "+msg.fromUserName, 0*time.Second, 1, user.AuthToken)
				case "当前时间":
					now = time.Now()
					weekday = now.Weekday()
					week = weekdayStr[weekday]
					sendMsg = fmt.Sprintf("私信---今天是%s %s", now.Format("2006-01-02 15:04:05"), week)
					u.Send(consts.OPERATE_GROUP_MSG, msg.roomId, msg.fromUserId, sendMsg, 0*time.Second, 1, u.AuthToken)
				}
			}
		}
	}()
}
