// Package mongos
// author gmfan
// date 2023/2/28
package mongos

import (
	"context"
	"github.com/tkgfan/got/core/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var db *mongo.Database

func DB() *mongo.Database {
	return db
}

// InitDatabase 初始化 Mongo
func InitDatabase(ctx context.Context, c Conf) (err error) {
	client, err := mongo.Connect(ctx, options.
		Client().
		ApplyURI(c.URI).
		SetMaxConnecting(uint64(c.MaxPoolSize)).
		SetConnectTimeout(time.Duration(c.MaxTimeout)*time.Second))
	if err != nil {
		return
	}

	db = client.Database(c.DB)
	return nil
}

type Conf struct {
	URI         string
	DB          string
	MaxPoolSize int64
	// 单位秒
	MaxTimeout int64
}

// LoadEnvConf 加载环境变量中的 Mongo 配置，must 设置为环境变量中是否必须存在
func LoadEnvConf(c *Conf, must bool) {
	env.LoadStr(&c.URI, "MONGO_URI", must)
	env.LoadStr(&c.DB, "MONGO_DB", must)
	env.LoadInt64(&c.MaxPoolSize, "MONGO_MAX_POOL_SIZE", must)
	env.LoadInt64(&c.MaxTimeout, "MONGO_MAX_TIME_OUT", must)
}
