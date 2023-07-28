// author gmfan
// date 2023/6/27
package model

import (
	"github.com/tkgfan/got/core/strings"
	"time"
)

const (
	// LogStatusNormal 正常状态
	LogStatusNormal = 0
	// LogStatusErr 错误状态
	LogStatusErr = 1
)

type (
	// TraceLog 链路追踪日志模型
	TraceLog struct {
		ID   string `json:"id"`
		Logs []*Log `json:"logs"`
	}

	Log struct {
		// 时间戳为毫秒
		Start  int64 `json:"start"`
		// 资源，可以是 URL 路径
		Source string `json:"source"`
		Info   any   `json:"info"`
		Status int8  `json:"status"`
		// 花费时长
		Expensive    int64 `json:"expensive"`
	}
)

func NewTraceLog() *TraceLog {
	return &TraceLog{
		ID:   strings.Rand(12),
		Logs: make([]*Log, 0),
	}
}

func (t *TraceLog) AddLog(log *Log) {
	log.Expensive = time.Now().UnixMilli()-log.Start
	t.Logs = append(t.Logs, log)
}

func NewLog() *Log {
	return &Log{
		Start: time.Now().UnixMilli(),
	}
}
