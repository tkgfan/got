// author gmfan
// date 2023/6/27
package logx

import (
	"time"
)

const (
	// LogCtxKey 上下文 key
	LogCtxKey = "log_ctx"
)

type (
	// LogCtx 日志上下文
	LogCtx struct {
		Tid   string    `json:"tid"`
		Start time.Time `json:"start"`
		// 请求资源，可以是 URL 路径
		Source string `json:"source"`
	}
)
