package yamls

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

type SocketClientCon struct {
	Name         string `json:",optional"`
	ServerName   string `json:",optional"`
	WebsocketUrl string `json:",optional"`
	HttpApiUrl   string `json:",optional"`
}

var SocketClientConf *SocketClientCon

func init() {
	// 获取配置文件的路径
	realPath := getCurrentDir()
	websocketFilePath := realPath + "/socket-client.yaml"
	websocketFile := flag.String("socket-client-f", websocketFilePath, "the socket client config file")
	var c SocketClientCon
	conf.MustLoad(*websocketFile, &c)
	SocketClientConf = &c
}
