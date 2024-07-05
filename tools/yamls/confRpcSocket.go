package yamls

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

type RpcSocketClientCon struct {
	zrpc.RpcClientConf
	ServiceId string `json:",optional"`
}

var RpcSocketClientConf *RpcSocketClientCon

func init() {
	// 获取配置文件的路径
	realPath := getCurrentDir()
	websocketFilePath := realPath + "/file-rpc-socket.yaml"
	websocketFile := flag.String("socket-f", websocketFilePath, "the socket config file")
	var c RpcSocketClientCon
	conf.MustLoad(*websocketFile, &c)
	RpcSocketClientConf = &c
}
