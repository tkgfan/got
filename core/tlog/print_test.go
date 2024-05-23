// author gmfan
// date 2024/5/23

package tlog

import (
	"context"
	"testing"
)

func TestInfo(t *testing.T) {
	Info(NewTid(), "test")
	ctx := WithTid(context.TODO())
	CtxError(ctx, "hello world")
}

func TestSetLevel(t *testing.T) {
	SetLevel(WarnLevel)
	Info(NewTid(), "test")
	ctx := WithTid(context.TODO())
	CtxWarnf(ctx, "arg: %s", "hello world")
	CtxError(ctx, "hello world")
}
