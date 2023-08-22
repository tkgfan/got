// author lby
// date 2023/6/28

package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/tkgfan/got/core/errors"
	"time"
)

// Nil reply returned by Redis when key does not exist.
var Nil = redis.Nil

// Set 设置键值对，expire 单位为秒。此方法会对 value 进行序列化
func Set(ctx context.Context, key string, value any, expire int64) (err error) {
	// 序列化
	bs, err := json.Marshal(value)
	if err != nil {
		return errors.Wrapf(err, string(bs))
	}

	err = Client().Set(ctx, key, bs, time.Duration(expire)*time.Second).Err()
	if err != nil {
		return errors.Wrapf(err, string(bs))
	}
	return
}

// Get 获取 key 对应的 value，此方法会对 Redis 中的 value 进行反序列化
// 并将值保存到 res 中，所以 res 必须为指针
func Get(ctx context.Context, key string, res any) (err error) {
	val, err := Client().Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), res)
	if err != nil {
		return errors.Wrapf(err, val)
	}
	return
}

// Del 删除 keys
func Del(ctx context.Context, keys ...string) (err error) {
	err = Client().Del(ctx, keys...).Err()
	if err != nil {
		return errors.Wrapf(err, "keys=%+v", keys)
	}
	return
}
