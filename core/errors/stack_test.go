// Package errors
// author gmfan
// date 2023/2/24
package errors

import (
	"encoding/json"
	stderrors "errors"
	"github.com/tkgfan/got/core/testx"
	"testing"
)

func TestIsStackError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "Nil test",
			err:  nil,
			want: false,
		},
		{
			name: "普通 error",
			err:  stderrors.New("test"),
			want: false,
		},
		{
			name: "stackError",
			err:  New("stackError"),
			want: true,
		},
	}

	for _, tt := range tests {
		var stackError *stackError
		got := As(tt.err, &stackError)
		if got != tt.want {
			t.Errorf("IsStackError(): name: %s, got: %v, want: %v", tt.name, got, tt.want)
		}
	}
}

func TestStackError_MarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		arg    *stackError
		expect string
		hasErr bool
	}{
		{
			name: "普通测试",
			arg: &stackError{
				Cause:  stderrors.New("normal error"),
				Stacks: []*stack{},
			},
			expect: `{"cause":"normal error","stacks":[]}`,
			hasErr: false,
		},
		{
			name: "空指针测试",
			arg: &stackError{
				Stacks: []*stack{},
			},
			expect: `{"cause":"","stacks":[]}`,
			hasErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs, err := json.Marshal(tt.arg)
			if testx.NotPassErrCheck(t, tt.hasErr, err) {
				return
			}
			if string(bs) != tt.expect {
				t.Errorf("got: %s,expect: %s", bs, tt.expect)
			}
		})
	}

}
