package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"store-chat/tools/commons"
	"store-chat/tools/consts"
	"store-chat/tools/types"
	"store-chat/tools/yamls"
	"strconv"
	"strings"
	"sync"
	"time"
)

var apiUrl = ""
var loginUrl = "/api/user/login"

type UserApi interface {
	InitUserInfo(autoToken string)
	InitSocket(url string)
}

type DefaultUser struct {
	Clock     sync.RWMutex
	UserId    int64
	UserName  string
	AuthToken string
	Clients   []*TestClient
	Log       logx.Logger
	Conf      *yamls.SocketClientCon
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
	Idx         uint32
	IsClose     chan int
	R           *rand.Rand
}

type QA struct {
	roomId       int64
	fromUserId   int64
	fromUserName string
	message      string
	Extra        string
}

func (u *DefaultUser) InitUserInfo(autoToken string) {
	apiUrl = u.Conf.HttpApiUrl
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
	u.Clients = make([]*TestClient, 0)
}

func (u *DefaultUser) InitSocket(url string, roomName string, idx uint32) {
	if u.AuthToken == "" {
		panic("token为空，不进行InitSocket\n")
	}
	conn, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic(fmt.Sprintf("拨号失败:%v fail:%s", res, err.Error()))
	}
	u.Clients = append(u.Clients, &TestClient{
		Conn:        conn,
		SendMsgChan: make(chan types.ReceiveMsg, 100),
		RevMsgChan:  make(chan types.WriteMsgBody, 100),
		RevMsgFail:  make(chan string, 100),
		QAChan:      make(chan QA, 100),
		Timeout:     30,
		RoomName:    roomName,
		Idx:         idx,
		IsClose:     make(chan int),
		R:           rand.New(rand.NewSource(time.Now().UnixNano())),
	})
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
			//fmt.Printf("msg:%v %s\n", msg, t.RoomName)
			//fmt.Printf("%s 管道未读条数：%d\n", time.Now().Format("2006-01-02 15:04:05"), len(t.RevMsgChan))
			t.ReadPush(msg)
		}
	}()
}
func (t *TestClient) ReadPush(msg types.WriteMsgBody) {
	select {
	case t.RevMsgChan <- msg:
	}
	return
}
func (t *TestClient) Send() {
	go func() {
		for {
			select {
			case r := <-t.SendMsgChan:
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
		}
	}()
}
func (t *TestClient) IsCloseRun(wg *sync.WaitGroup) {
	select {
	case <-t.IsClose:
		wg.Done()
	}
}

// Operator
// @Desc：
// @param：url
func (u *DefaultUser) Operator(idx uint32, url string, roomId int64, roomName string) {
	u.InitSocket(url, roomName, idx)
	go func() {
		waiter := &sync.WaitGroup{}
		waiter.Add(1)
		u.Clients[idx].Auth(u.AuthToken, roomId, u.UserId)
		u.Clients[idx].ReadMsg(u)
		u.Clients[idx].SendQA(u)
		go func() {
			if u.UserName == "和平星(管理员)" {
				for {
					select {
					case <-time.After(time.Minute):
						u.Clients[idx].QAChan <- QA{
							roomId:       roomId,
							fromUserId:   u.UserId,
							fromUserName: u.UserName,
							message:      "红包推文",
						}
					}
				}
			} else if u.UserName == "压测官" {
				for {
					select {
					case <-time.After(time.Hour):
						iMax := u.Clients[idx].R.Intn(9000) + 1000
						u.Clients[idx].QAChan <- QA{
							roomId:       roomId,
							fromUserId:   u.UserId,
							fromUserName: u.UserName,
							message:      "run-pm",
							Extra:        strconv.Itoa(iMax),
						}
						for i := 0; i < iMax; i++ {
							u.Clients[idx].QAChan <- QA{
								roomId:       roomId,
								fromUserId:   u.UserId,
								fromUserName: u.UserName,
								message:      "PM",
								Extra:        strconv.Itoa(i),
							}
						}
					}
				}
			}
		}()
		u.Clients[idx].IsCloseRun(waiter)
		waiter.Wait()
		//fmt.Printf("【%s】离开了【%s】\n", u.UserName, roomName)
		return
	}()
}

func (t *TestClient) ReadMsg(user *DefaultUser) {
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
			case <-t.RevMsgFail:
				_ = t.Conn.Close()
				t.IsClose <- 1

			case m := <-t.RevMsgChan:
				if m.Method == consts.METHOD_ENTER_MSG {
					if data, ok := m.Event.Data.(map[string]interface{}); !ok {
						//fmt.Printf("m.Event.Data typeOf types.DataByEnter not ok\n")
					} else {
						clientIdStr = data["clientId"].(string)
						t.ClientId, _ = strconv.ParseInt(clientIdStr, 10, 64)
					}
				} else if m.Method == consts.METHOD_NORMAL_MSG {
					if data, ok := m.Event.Data.(map[string]interface{}); !ok {
						//fmt.Printf("m.Event.Data typeOf types.DataByNormal not ok\n")
					} else {
						//if m.Operate == consts.OPERATE_SINGLE_MSG {
						//	fmt.Printf("[%s]%s:接收私聊消息：\n     %s\n", user.UserName, m.ResponseTime, data["message"])
						//} else if m.Operate == consts.OPERATE_GROUP_MSG {
						//	fmt.Printf("[%s]%s:接收广播消息：\n     %s\n", user.UserName, m.ResponseTime, data["message"])
						//}
						roomIdStr, _ = data["roomId"].(string)
						userIdStr = data["fromUserId"].(string)
						if user.UserName == "蜻蜓队长(管理员)" || user.UserName == "和平星(管理员)" || user.UserName == "压测官" {
							roomId, _ = strconv.ParseInt(roomIdStr, 10, 64)
							fromUserId, _ = strconv.ParseInt(userIdStr, 10, 64)
							qA := QA{
								roomId:       roomId,
								fromUserId:   fromUserId,
								fromUserName: data["fromUserName"].(string),
								message:      data["message"].(string),
							}
							select {
							case t.QAChan <- qA:
							}
						}
					}
				}
			}
		}
	}()
}

func (t *TestClient) SendQA(user *DefaultUser) {
	var (
		now        time.Time
		weekday    time.Weekday
		sendMsg    = ""
		week       string
		weekdayStr = [...]string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
		moneySlice = make([]float64, 10)
	)
	moneySlice[0] = 0
	for i := 1; i < 10; i++ {
		select {
		case <-time.Tick(time.Millisecond):
			randNum := t.R.Float64()*(100-float64(i*10)) + float64(i*10)
			moneySlice[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", randNum), 64)
		}
	}
	sort.Float64s(moneySlice)
	go func() {
		for {
			select {
			case msg := <-t.QAChan:
				pushMsg := types.ReceiveMsg{
					Version:      1,
					Operate:      0,
					Method:       consts.METHOD_NORMAL_MSG,
					AuthToken:    user.AuthToken,
					RoomId:       msg.roomId,
					FromUserId:   user.UserId,
					FromClientId: t.ClientId,
					ToUserId:     0,
					Event:        types.Event{},
				}
				if user.UserName == "蜻蜓队长(管理员)" {
					switch msg.message {
					case "我是谁":
						sendMsg = "你是 " + msg.fromUserName
						pushMsg.Operate = consts.OPERATE_GROUP_MSG
						t.PushMsg(pushMsg, sendMsg, user)
					case "当前时间":
						now = time.Now()
						weekday = now.Weekday()
						week = weekdayStr[weekday]
						sendMsg = fmt.Sprintf("私信---今天是%s %s", now.Format("2006-01-02 15:04:05"), week)

						pushMsg.Operate = consts.OPERATE_SINGLE_MSG
						pushMsg.ToUserId = msg.fromUserId
						t.PushMsg(pushMsg, sendMsg, user)
					}
				}

				if user.UserName == "和平星(管理员)" {
					switch msg.message {
					case "许愿和平星":
						money := moneySlice[t.R.Intn(len(moneySlice))]
						if money == 0 {
							sendMsg = fmt.Sprintf("和平星与您插肩而过~下次再许愿吧")
						} else {
							sendMsg = fmt.Sprintf("恭喜您被和平星砸中了 获得$%v（纯文字）", money)
						}
						pushMsg.Operate = consts.OPERATE_SINGLE_MSG
						pushMsg.ToUserId = msg.fromUserId
						t.PushMsg(pushMsg, sendMsg, user)
					case "红包推文":
						sendMsg = "输入`许愿和平星`即可随机获得红包奖励哦~"
						pushMsg.Operate = consts.OPERATE_GROUP_MSG
						t.PushMsg(pushMsg, sendMsg, user)
					}
				}

				if user.UserName == "压测官" {
					switch msg.message {
					case "PM":
						sendMsg = fmt.Sprintf("%s:%s", msg.message, msg.Extra)
						pushMsg.Operate = consts.OPERATE_GROUP_MSG
						pushMsg.Method = consts.METHOD_PM_MSG
						t.PushMsg(pushMsg, sendMsg, user)
					case "run-pm":
						sendMsg = fmt.Sprintf("每小时随机进行 %s 消息推送压力测试", msg.Extra)
						pushMsg.Operate = consts.OPERATE_GROUP_MSG
						t.PushMsg(pushMsg, sendMsg, user)
					}
				}
			}
		}
	}()
}

func (t *TestClient) PushMsg(send types.ReceiveMsg, msg string, user *DefaultUser) {
	send.Event.Params = fmt.Sprintf("【%s:%s】BOX:%s", t.RoomName, user.UserName, msg)
	select {
	case t.SendMsgChan <- send:
	}
	return
}
