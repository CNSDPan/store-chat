package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
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

func (u *DefaultUser) InitSocket(url string) {
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
			fmt.Printf("msg:%v\n", msg)
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
