package socket

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"store-chat/dbs"
	"store-chat/model/mysqls"
	"store-chat/tools/tools"
	"testing"
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
	user.InitSocket(socketUrl, "")
	user.Client.Auth(user.AuthToken, 1810940924055547904, user.UserId)
	user.Read()
	user.SendQA()
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
	user.InitSocket(socketUrl, "")
	user.Client.Auth(user.AuthToken, 1810940924055547904, user.UserId)
	go func() {
		user.Read()
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

func TestGoUser1(t *testing.T) {
	user.Log = logx.WithContext(context.Background())
	user.InitUserInfo("2gDGQwDxsrX0UG8yRbophdHxHqD")

	var roomMap = tools.StoreMap
	for _, room := range roomMap {
		//fmt.Printf(room.Name)
		user.Operator(uint32(0), socketUrl, room.StoreID, room.Name)
	}
	select {}
}
