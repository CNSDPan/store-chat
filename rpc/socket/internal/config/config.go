package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	ServerName string `json:",optional"`
	zrpc.RpcServerConf
}
