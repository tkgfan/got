// author gmfan
// date 2023/6/27
package logx

import (
	"context"
	"encoding/json"
	"github.com/tkgfan/got/core/strs"
	"io"
	"os"
	"time"
)

// TraceOut 链路日志输出，可自定义
var TraceOut io.Writer = os.Stdout

// 日志级别
const (
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"
)

func levelToInt(level string) int {
	switch level {
	case InfoLevel:
		return 0
	case WarnLevel:
		return 1
	case ErrorLevel:
		return 2
	case PanicLevel:
		return 3
	default:
		return -1
	}
}

func levelNotPass(level string) bool {
	levelInt := levelToInt(level)
	return levelInt < curLevel
}

// 当前日志级别
var curLevel = 0

const (
	// TraceCtxKey 链路日志上下文 key
	TraceCtxKey = "trace_log"
)

// TraceLogFormat 格式化链路日志输出，可自定义
var TraceLogFormat TraceLogFormatFunc = func(log *TraceLog) string {
	if log == nil {
		return ""
	}
	entry := make(map[string]any)
	entry["tid"] = log.Tid
	entry["start"] = log.Start.UnixMilli()
	entry["source"] = log.Source
	entry["level"] = log.Level
	entry["duration"] = log.Start.UnixMilli() - time.Now().UnixMilli()
	entry["info"] = log.Info
	bs, err := json.Marshal(entry)
	if err != nil {
		Error(err)
		return ""
	}
	return string(bs)
}

type (
	// TraceCtx 链路日志上下文
	TraceCtx struct {
		Tid   string    `json:"tid"`
		Start time.Time `json:"start"`
		// 请求资源，可以是 URL 路径
		Source string `json:"source"`
	}

	TraceLog struct {
		Tid   string    `json:"tid"`
		Start time.Time `json:"start"`
		// 请求资源，可以是 URL 路径
		Source   string `json:"source"`
		Level    string `json:"level"`
		Duration int64  `json:"duration"`
		Info     any    `json:"info"`
	}

	// TraceLogFormatFunc 格式化链路日志
	TraceLogFormatFunc func(log *TraceLog) string
)

func NewTraceCtx(source string) *TraceCtx {
	return &TraceCtx{
		Tid:    strs.Rand(16),
		Start:  time.Now(),
		Source: source,
	}
}

// SetTraceCtx 设置链路日志上下文
func SetTraceCtx(ctx context.Context, source string) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, TraceCtxKey, NewTraceCtx(source))
}

// GetTraceCtx 获取链路日志上下文，不存在则会自动创建一个
func GetTraceCtx(ctx context.Context) *TraceCtx {
	if ctx == nil {
		return nil
	}
	if v, ok := ctx.Value(TraceCtxKey).(*TraceCtx); ok {
		return v
	}
	return NewTraceCtx("")
}
