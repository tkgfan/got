// author gmfan
// date 2024/5/23

package tlog

import (
	"context"
	"fmt"
	"os"
	"time"
)

type PrintFunc func(tid, level string, logStr string)

// Print 打印函数，可自定义
var Print PrintFunc = func(tid, level string, logStr string) {
	// 格式化时间
	now := time.Now().Format("2006-01-02 15:04:05.000")
	str := fmt.Sprintf("[%s] tid=%s level=%s %s\n", now, tid, level, logStr)
	if level == ErrorLevel {
		str = "\033[31m" + str + "\033[0m"
	} else if level == WarnLevel {
		str = "\033[33m" + str + "\033[0m"
	}
	_, err := os.Stdout.WriteString(str)
	if err != nil {
		panic(err)
	}
}

func Info(tid string, v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	Print(tid, InfoLevel, fmt.Sprint(v...))
}

func Infof(tid string, format string, v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	Print(tid, InfoLevel, fmt.Sprintf(format, v...))
}

func Warn(tid string, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	Print(tid, WarnLevel, fmt.Sprint(v...))
}

func Warnf(tid string, format string, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	Print(tid, WarnLevel, fmt.Sprintf(format, v...))
}

func Error(tid string, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	Print(tid, ErrorLevel, fmt.Sprint(v...))
}

func Errorf(tid string, format string, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	Print(tid, ErrorLevel, fmt.Sprintf(format, v...))
}

func CtxInfo(ctx context.Context, v ...any) {
	Info(GetTid(ctx), v...)
}

func CtxInfof(ctx context.Context, format string, v ...any) {
	Infof(GetTid(ctx), format, v...)
}

func CtxWarn(ctx context.Context, v ...any) {
	Warn(GetTid(ctx), v...)
}

func CtxWarnf(ctx context.Context, format string, v ...any) {
	Warnf(GetTid(ctx), format, v...)
}

func CtxError(ctx context.Context, v ...any) {
	Error(GetTid(ctx), v...)
}

func CtxErrorf(ctx context.Context, format string, v ...any) {
	Errorf(GetTid(ctx), format, v...)
}
