// author lby
// date 2023/2/15

package structs

import (
	"errors"
	"testing"
)

func TestIsNil(t *testing.T) {
	// 1
	if IsNil(nil) != true {
		t.Error("IsNil has error")
	}

	// 2
	if IsNil("") == true {
		t.Error("IsNil has error")
	}

	// 3
	var a *int
	if IsNil(a) != true {
		t.Error("IsNil has error")
	}

	// 4
	a = nil
	var b interface{} = a
	if IsNil(b) != true {
		t.Error("IsNil has error")
	}
}

func TestIsSerializable(t *testing.T) {
	tests := []struct {
		name string
		val  any
		want bool
	}{
		{
			name: "结构体用例",
			val: struct {
				Name string
				Age  int
			}{
				Name: "lby",
				Age:  18,
			},
			want: true,
		},
		{
			name: "数组用例",
			val:  []int{1, 2, 3},
			want: true,
		},
		{
			name: "map用例",
			val: map[string]int{
				"a": 1,
				"b": 2,
			},
			want: true,
		},
		{
			name: "chan用例",
			val:  make(chan int),
			want: false,
		},
		{
			name: "func用例",
			val:  func() {},
			want: false,
		},
		{
			name: "nil用例",
			val:  nil,
			want: true,
		},
		{
			name: "interface 用例",
			val:  errors.New("error 是接口不可序列化"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsSerializable(tt.val) != tt.want {
				t.Errorf("IsSerializable(%v) = %v, want %v", tt.val, !tt.want, tt.want)
			}
		})
	}
}
