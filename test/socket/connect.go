package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"log"
	"store-chat/tools/consts"
	"store-chat/tools/types"
)

type TestClient struct {
	Conn        *websocket.Conn
	Timeout     int
	ClientId    int64
	UserId      int64
	UserName    string
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

func New(url string) (tClient *TestClient, err error) {
	//var d *websocket.Dialer
	//d.HandshakeTimeout = 30 * time.Second
	//conn, res, err := d.Dial(url, nil)
	fmt.Println(url)
	conn, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("拨号失败:%v fail:%v", res, err)

		return nil, err
	}
	return &TestClient{
		Conn:        conn,
		SendMsgChan: make(chan types.ReceiveMsg, 100),
		RevMsgChan:  make(chan types.WriteMsgBody, 100),
		RevMsgFail:  make(chan string, 100),
		QAChan:      make(chan QA, 100),
		Timeout:     30,
	}, nil
}

func (t *TestClient) Auth(authToken string) error {
	msg := types.ReceiveMsg{
		Version:   1,
		Operate:   10,
		Method:    consts.METHOD_ENTER_MSG,
		AutoToken: authToken,
		RoomId:    1,
		Event:     types.Event{},
	}
	b, err := jsonx.Marshal(msg)
	if err != nil {
		fmt.Println("jsonx.Marshal fail:", err.Error())
		return err
	}
	err = t.Conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		fmt.Println("t.Conn.WriteMessage fail:", err.Error())
		return err
	}
	return nil
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
