// author gmfan
// date 2024/3/30
package logx

import (
	"context"
	"fmt"
	"github.com/tkgfan/got/core/structs"
)

func TraceInfo(ctx context.Context, v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	tracePrint(ctx, InfoLevel, v)
}

func TraceInfof(ctx context.Context, format string, v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	tracePrint(ctx, InfoLevel, fmt.Sprintf(format, v...))
}

func TraceWarn(ctx context.Context, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	tracePrint(ctx, WarnLevel, v)
}

func TraceWarnf(ctx context.Context, format string, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	tracePrint(ctx, WarnLevel, fmt.Sprintf(format, v...))
}

func TraceError(ctx context.Context, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	tracePrint(ctx, ErrorLevel, v)
}

func TraceErrorf(ctx context.Context, format string, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	tracePrint(ctx, ErrorLevel, fmt.Sprintf(format, v...))
}

func TracePanic(ctx context.Context, v ...any) {
	if levelNotPass(PanicLevel) {
		return
	}
	tracePrint(ctx, PanicLevel, v)
}

func TracePanicf(ctx context.Context, format string, v ...any) {
	if levelNotPass(PanicLevel) {
		return
	}
	tracePrint(ctx, PanicLevel, fmt.Sprintf(format, v...))
}

func tracePrint(ctx context.Context, level string, logInfo any) {
	// 提取上下文
	tc := GetTraceCtx(ctx)
	log := &TraceLog{
		Level: level,
		Info:  logInfo,
	}
	structs.CopyFields(log, tc)
	var outStr string
	if levelToInt(level) > levelToInt(InfoLevel) {
		// 需要染色
		outStr = red(TraceLogFormat(log))
	} else {
		// 不需要染色
		outStr = TraceLogFormat(log)
	}
	_, err := TraceOut.Write([]byte(outStr + "\n"))
	if err != nil {
		Error(err)
	}
}
