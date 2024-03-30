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
	TraceError(context.Background(), errs.New("cause error 1"), errs.New("cause error 2"))
}
