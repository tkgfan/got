// author gmfan
// date 2023/6/27
package logx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tkgfan/got/core/strs"
	"github.com/tkgfan/got/core/structs"
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
)

func levelToInt(level string) int {
	switch level {
	case InfoLevel:
		return 0
	case WarnLevel:
		return 1
	case ErrorLevel:
		return 2
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
	entry["ip"] = log.Ip
	entry["source"] = log.Source
	entry["level"] = log.Level
	entry["duration"] = time.Now().UnixMilli() - log.Start.UnixMilli()

	// 处理 info 数组
	switch log.Info.(type) {
	case []interface{}:
		arr := log.Info.([]interface{})
		var infos []any
		for i := 0; i < len(arr); i++ {
			info := arr[i]
			if structs.IsSerializable(info) {
				infos = append(infos, info)
			} else {
				infos = append(infos, fmt.Sprint(info))
			}
		}
		entry["info"] = infos
	default:
		entry["info"] = log.Info
	}

	// 执行序列化
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
		Tid string `json:"tid"`
		// 用户 IP
		Ip    string    `json:"ip"`
		Start time.Time `json:"start"`
		// 请求资源，可以是 URL 路径
		Source string `json:"source"`
	}

	TraceLog struct {
		Tid   string    `json:"tid"`
		Ip    string    `json:"ip"`
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

func NewTraceCtx(tid, ip, source string) *TraceCtx {
	if tid == "" {
		tid = strs.Rand(16)
	}
	return &TraceCtx{
		Tid:    tid,
		Ip:     ip,
		Start:  time.Now(),
		Source: source,
	}
}

// SetTraceCtx 设置链路日志上下文
func SetTraceCtx(ctx context.Context, tid, ip, source string) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, TraceCtxKey, NewTraceCtx(tid, ip, source))
}

// GetTraceCtx 获取链路日志上下文，不存在则会自动创建一个
func GetTraceCtx(ctx context.Context) *TraceCtx {
	if ctx == nil {
		return nil
	}
	if v, ok := ctx.Value(TraceCtxKey).(*TraceCtx); ok {
		return v
	}
	return NewTraceCtx("", "", "")
}
