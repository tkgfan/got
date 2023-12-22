// author gmfan
// date 2023/6/27
package logx

import (
	"time"
)

const (
	// StatusNormal 正常状态
	StatusNormal = 0
	// StatusErr 错误状态
	StatusErr = 1
)

type (
	// TraceLog 链路追踪日志模型
	TraceLog struct {
		ID string `json:"id"`
		// 时间戳为毫秒
		Start int64 `json:"start"`
		// 资源，可以是 URL 路径
		Source string `json:"source"`
		// 花费时长，此值会自动更新
		Expensive int64  `json:"expensive"`
		Logs      []*Log `json:"logs"`
	}

	Log struct {
		// 时间戳为毫秒
		Start int64 `json:"start"`
		// 资源，可以是 URL 路径
		Source string `json:"source"`
		Info   any    `json:"info"`
		Status int8   `json:"status"`
		// 花费时长，此值会自动更新
		Expensive int64 `json:"expensive"`
	}
)

func NewTraceLog() *TraceLog {
	return &TraceLog{
		ID:    strs.Rand(16),
		Start: time.Now().UnixMilli(),
		Logs:  make([]*Log, 0),
	}
}

func (t *TraceLog) AddLog(log *Log) {
	// 计算 Expensive
	curMilli := time.Now().UnixMilli()
	log.Expensive = curMilli - log.Start
	t.Expensive = curMilli - t.Start

	t.Logs = append(t.Logs, log)
}

func NewLog() *Log {
	return &Log{
		Start: time.Now().UnixMilli(),
	}
}
