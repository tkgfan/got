// author gmfan
// date 2024/3/30
package logx

import (
	"context"
	"github.com/tkgfan/got/core/errs"
	"testing"
)

func TestTraceInfo(t *testing.T) {
	TraceInfo(context.Background(), "test")
}

func TestTraceError(t *testing.T) {
	TraceError(context.Background(), errs.New("cause error info"))
	TraceError(SetTraceCtx(context.Background(),
		"tid", "127.0.0.1", "source"),
		errs.New("cause error 1"),
		errs.New("cause error 2"))
}
