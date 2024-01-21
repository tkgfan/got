// author gmfan
// date 2024/1/21
package testx

import "testing"

// NotPassErrCheck 检查 err 是否能符合 expectHasErr 条件，不符合返回 true，反之 false
func NotPassErrCheck(t *testing.T, expectHasErr bool, err error) bool {
	if expectHasErr && err == nil {
		t.Error("err 不能为空")
		return true
	}
	if !expectHasErr && err != nil {
		t.Error(err)
		return true
	}
	return false
}
