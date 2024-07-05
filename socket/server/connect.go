package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// MaxMessageSize 消息大小
	MaxMessageSize = 8192
	// PingPeriod 每次ping的间隔时长
	PingPeriod = 30 * time.Second
	// PongPeriod 每次pong的间隔时长，可以是PingPeriod的一倍|两倍
	PongPeriod = 60 * time.Second
	// WriteWait client的写入等待时长
	WriteWait = 5 * time.Second
	// ReadWait client的读取等待时长
	ReadWait = 60 * time.Second
)

type Connect struct{}

func NewConnect() *Connect {
	return &Connect{}
}

func (c *Connect) Run(w http.ResponseWriter, r *http.Request, socket *Server) {
	wsConn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		panic("http转换升级为websocket失败：")
	}

	wsConn.SetReadLimit(MaxMessageSize)
	wsConn.SetPongHandler(func(string) error {
		_ = wsConn.SetReadDeadline(time.Now().Add(PongPeriod))
		return nil
	})

	client := NewClient(wsConn)
	client.ClientId = socket.Node.Generate().Int64()

	go socket.writeChannel(client)
	go socket.readChannel(client)
}
