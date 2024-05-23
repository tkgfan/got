// author gmfan
// date 2024/5/23

package tlog

import (
	"fmt"
	"io"
	"os"
)

// 日志级别
const (
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

// TraceOut 链路日志输出，可自定义
var TraceOut io.Writer = os.Stdout

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

func SetLevel(_level string) {
	toInt := levelToInt(_level)
	if toInt == -1 {
		panic(fmt.Sprintf("日志级别不合法：%s", _level))
	}
	curLevel = toInt
}
