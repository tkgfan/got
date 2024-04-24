// author gmfan
// date 2024/3/30
package logx

import (
	"context"
	"errors"
	"github.com/tkgfan/got/core/errs"
	"testing"
)

func TestTraceInfo(t *testing.T) {
	TraceInfo(context.Background(), "test")
	people := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "tkg",
		Age:  18,
	}
	TraceInfo(context.TODO(), people)
}

func TestTraceError(t *testing.T) {
	TraceError(context.Background(), errors.New("cause error info"))
	TraceError(SetTraceCtx(context.Background(),
		"tid", "127.0.0.1", "source"),
		errs.New("cause error 1"),
		errs.New("cause error 2"))
}
