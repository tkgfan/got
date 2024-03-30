// description: 日志级别由低到高依次为：InfoLevel、ErrorLevel、PanicLevel。
// 系统默认为 InfoLevel
// author lby
// date 2023/10/31

package logx

import (
	"fmt"
	"log"
)

const (
	redColorStart = "\033[31m "
	redColorEnd   = "\033[0m"
)

func red(str string) string {
	return redColorStart + str + redColorEnd
}

func SetLevel(_level string) {
	toInt := levelToInt(_level)
	if toInt == -1 {
		panic(fmt.Sprintf("日志级别不合法：%s", _level))
	}
	curLevel = toInt
}

func Info(v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	colorPrint(InfoLevel, fmt.Sprint(v...))
}

func Infof(format string, v ...any) {
	if levelNotPass(InfoLevel) {
		return
	}
	colorPrint(InfoLevel, fmt.Sprintf(format, v...))
}

func Warn(v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	colorPrint(WarnLevel, fmt.Sprint(v...))
}

func Warnf(format string, v ...any) {
	if levelNotPass(WarnLevel) {
		return
	}
	colorPrint(WarnLevel, fmt.Sprintf(format, v...))
}

func Error(v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	colorPrint(ErrorLevel, fmt.Sprint(v...))
}

func Errorf(format string, v ...any) {
	if levelNotPass(ErrorLevel) {
		return
	}
	colorPrint(ErrorLevel, fmt.Sprintf(format, v...))
}

func Panic(v ...any) {
	if levelNotPass(PanicLevel) {
		return
	}
	colorPrint(PanicLevel, fmt.Sprint(v...))
}

func Panicf(format string, v ...any) {
	if levelNotPass(PanicLevel) {
		return
	}
	colorPrint(PanicLevel, fmt.Sprintf(format, v...))
}

func colorPrint(level, logStr string) {
	if level == PanicLevel {
		log.Panic(red(logStr))
	}
	if levelToInt(level) < levelToInt(WarnLevel) {
		// 普通打印
		log.Println(logStr)
		return
	}
	log.Println(red(logStr))
}
