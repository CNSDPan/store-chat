package socket

import (
	"context"
	"fmt"
	"store-chat/rpc/socket/pb/socket"
	"store-chat/tools/commons"
	"store-chat/tools/consts"
	"testing"
)

func TestBroadcastLogin(t *testing.T) {
	//messageA := &socket.EventDataLogin{
	//	RoomId:   1,
	//	UserId:   1,
	//	UserName: "欸",
	//}
	//anyA, err := anypb.New(messageA)
	//fmt.Printf("aa:%+v err:%v \n", anyA, err)
	//messageB := &socket.EventDataLogin{}
	//err = anyA.UnmarshalTo(messageB)
	//fmt.Printf("messageB:%+v err:%v \n", messageB, err)

	rpcSocketBroadcast := New()
	result, err := rpcSocketBroadcast.BroadcastLogin(context.Background(), &socket.ReqBroadcastMsg{
		Version:   int32(1),
		Operate:   int32(10),
		Method:    consts.METHOD_ENTER_MSG,
		AutoToken: "2gDGQwDxsrX0UG8yRbophdHxHqD",
		RoomId:    1,
	})
	if err != nil {
		panic("rpc.err " + err.Error())
	}
	data := &socket.EventDataLogin{}
	if result.Code == commons.RESPONSE_SUCCESS {
		if err = result.Data.UnmarshalTo(data); err != nil {
			panic("result.Data.UnmarshalTo " + err.Error())
		}
	}
	fmt.Printf("code:%s msg:%s data:%+v \n", result.Code, result.Msg, data)
}
