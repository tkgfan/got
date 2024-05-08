// author gmfan
// date 2024/5/8
package dsl

import (
	"testing"
)

func TestDSL_PushToArray(t *testing.T) {
	d := New()
	d.PushToArray("hello", "a")
	if d.String() != `{"a":["hello"]}` {
		t.Error("push to array error")
	}

	d = New()
	d.PushToArray("hello", "a", "b")
	d.PushToArray("world", "a", "b")
	if d.String() != `{"a":{"b":["hello","world"]}}` {
		t.Error("push to array error")
	}
}
