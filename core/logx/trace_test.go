// author gmfan
// date 2023/10/15
package logx

import (
	"context"
	"testing"
)

func TestSetTraceLog(t *testing.T) {
	type (
		arg struct {
			ctx         context.Context
			traceLogStr string
			source      string
		}
		expect struct {
			traceLog *TraceLog
			log      *Log
			hasErr   bool
		}
	)

	tests := []struct {
		name   string
		arg    arg
		expect expect
	}{
		{
			name: "普通测试",
			arg: arg{
				traceLogStr: "",
				source:      "/normal/test",
				ctx:         context.TODO(),
			},
			expect: expect{
				traceLog: &TraceLog{
					Source: "/normal/test",
				},
				log: &Log{
					Source: "/normal/test",
				},
				hasErr: false,
			},
		},
		{
			name: "测试 traceLogStr",
			arg: arg{
				traceLogStr: `{"id":"1","start":1,"source":"/t","expensive":10,"logs":[]}`,
				source:      "/t",
				ctx:         context.TODO(),
			},
			expect: expect{
				traceLog: &TraceLog{
					ID:        "1",
					Start:     1,
					Source:    "/t",
					Expensive: 10,
					Logs:      make([]*Log, 0),
				},
				log: &Log{
					Source: "/t",
				},
				hasErr: false,
			},
		},
		{
			name: "测试反序列化错误",
			arg: arg{
				traceLogStr: "{{",
				ctx:         context.TODO(),
			},
			expect: expect{
				hasErr: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.arg
			e := tt.expect
			got, err := SetTraceLog(a.ctx, a.traceLogStr, a.source)
			if (e.hasErr && err == nil) || (!e.hasErr && err != nil) {
				t.Error(e.hasErr, err)
				return
			}
			if err != nil {
				return
			}

			gotTraceLog := got.Value(TraceLogKey).(*TraceLog)
			if gotTraceLog.Source != e.traceLog.Source ||
				len(gotTraceLog.Logs) != len(e.traceLog.Logs) {
				t.Errorf("expect: %+v, got: %+v", e.traceLog, gotTraceLog)
				return
			}

			gotLog := got.Value(LogKey).(*Log)
			if gotLog.Source != e.log.Source {
				t.Errorf("expect: %+v, got: %+v", e.log, gotLog)
				return
			}
		})
	}
}

func TestUpdateTraceLog(t *testing.T) {
	// TODO
}
