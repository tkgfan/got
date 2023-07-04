// author gmfan
// date 2023/6/27
package model

import (
	"github.com/tkgfan/got/core/strings"
	"time"
)

type (
	// TraceLog 链路追踪日志模型
	TraceLog struct {
		ID   string `json:"id"`
		Logs []*Log `json:"logs"`
	}

	Log struct {
		Start int64 `json:"start"`
		Info  any   `json:"info"`
		IsErr bool  `json:"is_err"`
		End   int64 `json:"end"`
	}
)

func NewTraceLog() *TraceLog {
	return &TraceLog{
		ID:   strings.Rand(12),
		Logs: make([]*Log, 0),
	}
}

func (t *TraceLog) AddLog(log *Log) {
	log.End = time.Now().Unix()
	t.Logs = append(t.Logs, log)
}

func NewLog() *Log {
	return &Log{
		Start: time.Now().Unix(),
	}
}
