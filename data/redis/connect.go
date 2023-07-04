// Package redis
// author gmfan
// date 2023/3/1
package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/tkgfan/got/core/env"
)

var c Conf
var client *redis.Client

func Client() *redis.Client {
	return client
}

// InitRedis 初始化 Redis 客户端
func InitRedis(conf Conf) (err error) {
	c = conf
	client = redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Pwd,
		DB:       c.DB,
		PoolSize: conf.PoolSize,
	})
	_, err = client.Ping(context.Background()).Result()
	return
}

type Conf struct {
	Host     string
	Pwd      string
	DB       int
	PoolSize int
}

// LoadEnvConf 加载环境比那辆中的 Redis 配置，must 设置为是否必须
func LoadEnvConf(c *Conf, must bool) {
	env.LoadStr(&c.Host, "REDIS_HOST", must)
	env.LoadStr(&c.Pwd, "REDIS_PWD", must)
	env.LoadInt(&c.DB, "REDIS_DB", must)
	env.LoadInt(&c.PoolSize, "REDIS_POOL_SIZE", must)
}
