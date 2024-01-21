// Package errors
// author gmfan
// date 2023/2/24
package errors

import (
	stderrors "errors"
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
		_, got := tt.err.(*stackError)
		if got != tt.want {
			t.Errorf("IsStackError(): name: %s, got: %v, want: %v", tt.name, got, tt.want)
		}
	}
}
