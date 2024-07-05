package dbs

import (
	"context"
	"github.com/redis/go-redis/v9"
	"store-chat/tools/yamls"
	"sync"
)

var syncLock sync.Mutex

var RedisClient *redis.Client
var AloneRedisClientMap = map[string]*redis.Client{}

func NewRedisClient() (redis *redis.Client, err error) {
	switch yamls.RedisConf.RedisClientType {
	case 1:
		redis, err = NewAloneRedis()
	}
	RedisClient = redis
	return
}

// NewAloneRedis
// @Desc：单机
// @return：*redis.Client
// @return：error
func NewAloneRedis() (*redis.Client, error) {
	var (
		redisClient *redis.Client
		ok          bool
		err         error
	)
	syncLock.Lock()
	defer func() {
		_, err = redisClient.Ping(context.Background()).Result()
	}()
	defer syncLock.Unlock()
	if redisClient, ok = AloneRedisClientMap[yamls.RedisConf.AloneRedisConf.Addr]; ok {
		return redisClient, err
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     yamls.RedisConf.AloneRedisConf.Addr,
		Password: yamls.RedisConf.AloneRedisConf.Password,
		DB:       yamls.RedisConf.AloneRedisConf.DB,
	})
	AloneRedisClientMap[yamls.RedisConf.AloneRedisConf.Addr] = redisClient
	return AloneRedisClientMap[yamls.RedisConf.AloneRedisConf.Addr], err
}
