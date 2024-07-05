package yamls

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

type RedisConfer struct {
	IsInit          int8       `default:"0"`
	RedisClientType int8       `json:",optional"`
	AloneRedisConf  AloneRedis `json:",optional,inherit"`
}

// AloneRedis 单机
type AloneRedis struct {
	Name     string `json:",optional"`
	Addr     string `json:",optional"`
	Password string `json:",optional"`
	DB       int    `json:",optional"`
}

var RedisConf *RedisConfer

func init() {
	// 获取配置文件的路径
	realPath := getCurrentDir()
	redisFilePath := realPath + "/file-redis.yaml"
	redisFile := flag.String("redis-f", redisFilePath, "the redis config file")
	var c RedisConfer
	conf.MustLoad(*redisFile, &c)
	c.IsInit = 1
	RedisConf = &c
}
