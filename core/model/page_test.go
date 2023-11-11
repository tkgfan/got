// author gmfan
// date 2023/11/11
package model

import "testing"

func TestNewPageResult(t *testing.T) {
	res := NewPageResult(nil, 0)
	if res.Rows == nil {
		t.Error("rows 不能为 nil")
	}
}
