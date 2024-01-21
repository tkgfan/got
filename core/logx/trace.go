// author gmfan
// date 2023/10/15
package logx

import (
	"context"
	"encoding/json"
	"github.com/tkgfan/got/core/structs"
)

const (
	TraceLogKey = "Trace-Log"
	LogKey      = "Log"
)

// SetTraceLog 在上下文中设置链路日志与当前日志，traceLogStr 不为空则上下文链路日志直接为
// traceLogStr
func SetTraceLog(ctx context.Context, traceLogStr, source string) (res context.Context, err error) {
	// 设置当前日志
	curLog := NewLog()
	curLog.Source = source
	res = context.WithValue(ctx, LogKey, curLog)

	tl := NewTraceLog()
	if traceLogStr != "" {
		// traceLogStr 不为空则上下文中的链路日志直接使用 traceLogStr
		err = json.Unmarshal([]byte(traceLogStr), tl)
		if err != nil {
			return
		}
	}

	// 设置链路日志
	tl.Source = source
	res = context.WithValue(res, TraceLogKey, tl)

	return res, nil
}

// UpdateTraceLog 更新链路日志，将链路日志设置为 traceLogStr，若 traceLogStr 为空
// 且上下文中不存在链路日志则创建一个链路日志。上下文中不存在当前日志则会自动创建一个当前日志。
func UpdateTraceLog(ctx context.Context, traceLogStr, source string) (res context.Context, err error) {
	var tl *TraceLog
	// 尝试获取上下文中的链路日志
	if t := ctx.Value(TraceLogKey); t != nil {
		tl = t.(*TraceLog)
	} else {
		// 上下文中不存在 Trace Log 则创建一个链路日志并保存到上下文中
		tl = NewTraceLog()
		tl.Source = source
		// 设置链路日志
		ctx = context.WithValue(ctx, TraceLogKey, tl)
	}

	if traceLogStr != "" {
		// traceLogStr 不为空则上下文中的链路日志直接使用 traceLogStr
		err = json.Unmarshal([]byte(traceLogStr), tl)
		if err != nil {
			return
		}
	}

	// 若上下文不存在当前日志则设置当前日志
	if ctx.Value(LogKey) == nil {
		curLog := NewLog()
		curLog.Source = source
		ctx = context.WithValue(ctx, LogKey, curLog)
	}
	return ctx, nil
}

// GetTraceLog 获取链路日志
func GetTraceLog(ctx context.Context) (res *TraceLog) {
	return ctx.Value(TraceLogKey).(*TraceLog)
}

// GetTraceLogStr 获取序列化后的链路日志
func GetTraceLogStr(ctx context.Context) string {
	res := GetTraceLog(ctx)
	bs, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// TraceLogAddLog 将日志添加到链路日志中
func TraceLogAddLog(ctx context.Context, log *Log) {
	tl := ctx.Value(TraceLogKey).(*TraceLog)
	tl.AddLog(log)
}

// TraceLogMergeLog Trace Log 合并 Log
func TraceLogMergeLog(ctx context.Context, err error) {
	tl := GetTraceLog(ctx)
	curLog := ctx.Value(LogKey).(*Log)
	if !structs.IsNil(err) {
		curLog.Info = err
		curLog.Status = StatusErr
		tl.HasErr = StatusErr
	}
	tl.AddLog(curLog)
}
