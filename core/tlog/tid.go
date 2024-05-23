// author gmfan
// date 2024/5/23

package tlog

import (
	"context"
	"github.com/tkgfan/got/core/strs"
)

const CtxTidKey = "tid"

// NewTid 生成 tid
func NewTid() string {
	return strs.Rand(16)
}

func GetTid(ctx context.Context) string {
	val := ctx.Value(CtxTidKey)
	if val == nil {
		return ""
	}
	return val.(string)
}

// WithTid 上下文中加入 tid，若 tid 已存在则不做任何操作
func WithTid(ctx context.Context) context.Context {
	tid := GetTid(ctx)
	if tid != "" {
		return ctx
	}
	return context.WithValue(ctx, CtxTidKey, NewTid())
}
