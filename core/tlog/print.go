// author gmfan
// date 2024/5/23

package tlog

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type CallFrame struct {
	Pc   uintptr
	File string
	Line int
}

func caller(skip int) (frame CallFrame) {
	frame.Pc, frame.File, frame.Line, _ = runtime.Caller(skip + 1)
	return frame
}

type PrintFunc func(tid, level string, logStr string, frame CallFrame)

// Print 打印函数，可自定义
var Print PrintFunc = func(tid, level string, logStr string, frame CallFrame) {
	// 格式化时间
	now := time.Now().Format("2006-01-02 15:04:05.000")
	str := fmt.Sprintf("[%s] tid=%s level=%s %s caller=%s\n",
		now, tid, level, logStr, frame.File+" "+strconv.Itoa(frame.Line))
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
	Print(tid, InfoLevel, fmt.Sprint(v...), caller(1))
}

func Infof(tid string, format string, v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	Print(tid, InfoLevel, fmt.Sprintf(format, v...), caller(1))
}

func Warn(tid string, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	Print(tid, WarnLevel, fmt.Sprint(v...), caller(1))
}

func Warnf(tid string, format string, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	Print(tid, WarnLevel, fmt.Sprintf(format, v...), caller(1))
}

func Error(tid string, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	Print(tid, ErrorLevel, fmt.Sprint(v...), caller(1))
}

func Errorf(tid string, format string, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	Print(tid, ErrorLevel, fmt.Sprintf(format, v...), caller(1))
}

func CtxInfo(ctx context.Context, v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	Print(GetTid(ctx), InfoLevel, fmt.Sprint(v...), caller(1))
}

func CtxInfof(ctx context.Context, format string, v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	Print(GetTid(ctx), InfoLevel, fmt.Sprintf(format, v...), caller(1))
}

func CtxWarn(ctx context.Context, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	Print(GetTid(ctx), WarnLevel, fmt.Sprint(v...), caller(1))
}

func CtxWarnf(ctx context.Context, format string, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	Print(GetTid(ctx), WarnLevel, fmt.Sprintf(format, v...), caller(1))
}

func CtxError(ctx context.Context, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	Print(GetTid(ctx), ErrorLevel, fmt.Sprint(v...), caller(1))
}

func CtxErrorf(ctx context.Context, format string, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	Print(GetTid(ctx), ErrorLevel, fmt.Sprintf(format, v...), caller(1))
}
