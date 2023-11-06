// description: 日志级别由低到高依次为：InfoLevel、ErrorLevel、PanicLevel。
// 系统默认为 InfoLevel
// author lby
// date 2023/10/31

package logx

import (
	"fmt"
	"log"
)

// 日志级别
const (
	InfoLevel  = "info"
	ErrorLevel = "error"
	PanicLevel = "panic"
)

const (
	redColorStart = "\033[31m "
	redColorEnd   = "\033[0m"
)

func red(str string) string {
	return redColorStart + str + redColorEnd
}

var level = InfoLevel

func SetLevel(_level string) {
	if level != InfoLevel && level != ErrorLevel && level != PanicLevel {
		panic(fmt.Sprintf("日志级别不合法：%s", level))
	}
	level = _level
}

func Info(v ...any) {
	if level == ErrorLevel || level == PanicLevel {
		return
	}
	log.Println(v...)
}

func Infof(format string, v ...any) {
	if level == ErrorLevel || level == PanicLevel {
		return
	}
	log.Printf(format, v...)
}

func Error(v ...any) {
	if level == PanicLevel {
		return
	}
	log.Println(red(fmt.Sprint(v...)))
}

func Errorf(format string, v ...any) {
	if level == PanicLevel {
		return
	}
	log.Println(red(fmt.Sprintf(format, v...)))
}

func Panic(v ...any) {
	log.Panicln(red(fmt.Sprint(v...)))
}

func Panicf(format string, v ...any) {
	log.Panicf(red(format), v...)
}
