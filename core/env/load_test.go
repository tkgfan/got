// author gmfan
// date 2023/6/21

package env

import (
	"os"
	"testing"
)

type loadArg struct {
	must     bool
	key      string
	envKey   string
	envValue string
}

type loadTest struct {
	arg    loadArg
	expect any
	hasErr bool
}

func execute(t *testing.T, tt loadTest, fn executeFunc) {
	defer func() {
		os.Unsetenv(tt.arg.envKey)
		err := recover()
		if tt.hasErr && err == nil {
			t.Errorf("expect: err!=nil, got err=nil, arg:%+v", tt.arg)
		}
		if !tt.hasErr && err != nil {
			t.Errorf("expect: err=nil, got: err=%+v, arg:%+v", err, tt.arg)
		}
	}()

	os.Setenv(tt.arg.envKey, tt.arg.envValue)

	if res, ok := fn(); !ok {
		t.Errorf("expect: %+v, got: %+v, arg: %+v", tt.expect, res, tt.arg)
	}
}

type executeFunc func() (res any, ok bool)

func TestLoadStr(t *testing.T) {
	tests := []loadTest{
		{
			arg:    loadArg{must: true, key: "str"},
			expect: "",
			hasErr: true,
		},
		{
			arg:    loadArg{must: true, key: "str", envKey: "str", envValue: "str"},
			expect: "str",
			hasErr: false,
		},
		{
			arg:    loadArg{must: false, key: "str"},
			expect: "",
			hasErr: false,
		},
		{
			arg:    loadArg{must: false, key: "str", envKey: "str", envValue: "str"},
			expect: "str",
			hasErr: false,
		},
	}

	for _, tt := range tests {
		fn := func() (res any, ok bool) {
			var dst string
			LoadStr(&dst, tt.arg.key, tt.arg.must)
			if dst != tt.expect.(string) {
				return dst, false
			}
			return dst, true
		}

		execute(t, tt, fn)
	}
}

func TestLoadInt64(t *testing.T) {
	tests := []loadTest{
		{
			arg:    loadArg{must: true, key: "int64"},
			expect: int64(0),
			hasErr: true,
		},
		{
			arg:    loadArg{must: true, key: "int64", envKey: "int64", envValue: "47"},
			expect: int64(47),
			hasErr: false,
		},
		{
			arg:    loadArg{must: false, key: "int64"},
			expect: int64(0),
			hasErr: false,
		},
		{
			arg:    loadArg{must: false, key: "int64", envKey: "int64", envValue: "147"},
			expect: int64(147),
			hasErr: false,
		},
		{
			arg:    loadArg{must: false, key: "int64", envKey: "int64", envValue: "not a number"},
			expect: int64(0),
			hasErr: true,
		},
	}

	for _, tt := range tests {
		fn := func() (res any, ok bool) {
			var dst int64
			LoadInt64(&dst, tt.arg.key, tt.arg.must)
			if dst != tt.expect.(int64) {
				return dst, false
			}
			return dst, true
		}

		execute(t, tt, fn)
	}
}

func TestLoadInt(t *testing.T) {
	tests := []loadTest{
		{
			arg:    loadArg{must: true, key: "int"},
			expect: int(0),
			hasErr: true,
		},
		{
			arg:    loadArg{must: true, key: "int", envKey: "int", envValue: "47"},
			expect: int(47),
			hasErr: false,
		},
		{
			arg:    loadArg{must: false, key: "int"},
			expect: int(0),
			hasErr: false,
		},
		{
			arg:    loadArg{must: false, key: "int", envKey: "int", envValue: "147"},
			expect: int(147),
			hasErr: false,
		},
		{
			arg:    loadArg{must: false, key: "int", envKey: "int", envValue: "not a number"},
			expect: 0,
			hasErr: true,
		},
	}

	for _, tt := range tests {
		fn := func() (res any, ok bool) {
			var dst int
			LoadInt(&dst, tt.arg.key, tt.arg.must)
			if dst != tt.expect.(int) {
				return dst, false
			}
			return dst, true
		}

		execute(t, tt, fn)
	}
}
