// author gmfan
// date 2023/12/23
package slices

import (
	"reflect"
	"testing"
)

func TestDuplicationStrs(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		res  []string
	}{
		{
			name: "普通测试-1",
			strs: []string{"apple", "banana", "apple"},
			res:  []string{"apple", "banana"},
		},
		{
			name: "普通测试-2",
			strs: []string{"苹果", "香蕉", "苹果"},
			res:  []string{"苹果", "香蕉"},
		},
		{
			name: "边界测试",
			strs: nil,
			res:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DuplicationStrs(tt.strs)
			if !reflect.DeepEqual(got, tt.res) {
				t.Errorf("got: %+v, expect: %+v", got, tt.res)
			}
		})
	}
}
