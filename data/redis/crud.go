// author lby
// date 2023/6/28

package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/tkgfan/got/core/errs"
	"time"
)

// Nil reply returned by Redis when key does not exist.
var Nil = redis.Nil

// Set 设置键值对，expire 单位为秒。此方法会对 value 进行序列化
func Set(ctx context.Context, key string, value any, expire int64) (err error) {
	// 序列化
	bs, err := json.Marshal(value)
	if err != nil {
		return errs.Wrapf(err, string(bs))
	}

	err = Client().Set(ctx, key, bs, time.Duration(expire)*time.Second).Err()
	if err != nil {
		return errs.Wrapf(err, string(bs))
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
		return errs.Wrapf(err, val)
	}
	return
}

// Del 删除 keys
func Del(ctx context.Context, keys ...string) (err error) {
	err = Client().Del(ctx, keys...).Err()
	if err != nil {
		return errs.Wrapf(err, "keys=%+v", keys)
	}
	return
}

// CountLimit 有效期内限制 key 的数量，并返回剩余数量。每次调用此函数都会增加 key 计数值。
// maxCount 为计数最大值，duration 为计数周期单位时间秒，remain 是返回剩余次数小于 0 则
// 表示已经达到上限。
func CountLimit(ctx context.Context, key string, maxCount int64, duration int64) (remain int64, err error) {
	// 尝试设置 key 起始值为 1
	res := Client().SetNX(ctx, key, 1, time.Duration(duration)*time.Second)
	if err = res.Err(); err != nil {
		return
	}
	if res.Val() {
		return maxCount - 1, nil
	}

	// 加一
	count, err := Client().Incr(ctx, key).Result()
	if err != nil {
		return
	}
	// key 超时过期
	if count == 1 {
		// 重新设置 key
		err = Client().Set(ctx, key, 1, time.Duration(duration)*time.Second).Err()
		return maxCount - 1, err
	}

	return maxCount - count, nil
}
