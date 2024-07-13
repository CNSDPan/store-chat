package mq

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"store-chat/dbs"
	"store-chat/socket/server"
	"store-chat/tools/commons"
	"store-chat/tools/consts"
	"store-chat/tools/types"
	"time"
)

type Subscribe struct {
	Log    logx.Logger
	PubSub struct {
		ctx context.Context
		*redis.PubSub
	}
}

func NewSubscribe() (*Subscribe, error) {
	ctx := context.Background()
	pubSub := dbs.RedisClient.Subscribe(ctx, commons.PUB_SUB_SOCKET_MESSAGE_NORMAL_CHANNEL_KEY)
	if _, err := pubSub.ReceiveTimeout(ctx, 100*time.Millisecond); err != nil {
		fmt.Printf("订阅 %s 接收消息异常，尝试 ping...", commons.PUB_SUB_SOCKET_MESSAGE_NORMAL_CHANNEL_KEY)
		if err = pubSub.Ping(ctx, ""); err != nil {
			return &Subscribe{}, err
		}
	}
	sub := &Subscribe{}
	sub.PubSub.ctx = ctx
	sub.PubSub.PubSub = pubSub
	return sub, nil
}

// SubReceive
// @Desc：订阅者接收发布消息并传递到每个连接池管道里
func (sub *Subscribe) SubReceive() {
	go func() {
		var (
			err error
			msg interface{}
			m   *redis.Message
		)
		defer sub.PubSub.Close()
		for {
			if msg, err = sub.PubSub.ReceiveTimeout(sub.PubSub.ctx, 100*time.Millisecond); err != nil {
				if err = sub.PubSub.Ping(sub.PubSub.ctx, ""); err != nil {
					sub.Log.Info("订阅消息服务 PubSub.Ping timeout channel.name:%s", commons.PUB_SUB_SOCKET_MESSAGE_NORMAL_CHANNEL_KEY)
					break
				}
				continue
			}
			switch msg.(type) {
			case *redis.Message:
				m = msg.(*redis.Message)
				var writeMsg types.WriteMsg
				b := []byte(m.Payload)
				if err = jsonx.Unmarshal(b, &writeMsg); err != nil {
					sub.Log.Errorf("订阅消息服务 Receive Channel:%s json.Unmarshal  fail:%s", m.Channel, err.Error())
				} else {
					if writeMsg.Operate == consts.OPERATE_SINGLE_MSG {
						// 私人消息
						bucket := server.DefaultServer.GetBucket(writeMsg.ToUserId)
						bucket.Routines <- writeMsg
					} else if writeMsg.Operate == consts.OPERATE_GROUP_MSG {
						// 群消息

						for _, bucket := range server.DefaultServer.Buckets {
							bucket.Routines <- writeMsg
						}
					}

				}
			}
		}
		return
	}()
}
