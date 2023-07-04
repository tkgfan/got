// Package env
// author gmfan
// date 2023/2/27
package env

import (
	"fmt"
	"os"
	"strconv"
)

func handleMust(key string) {
	panic(fmt.Sprint("环境变量: ", key, " 不存在!"))
}

// LoadStr 加载环境变量，must 设置是否为必须，must=true 时
// 如果环境变量不存在则会 panic
func LoadStr(dst *string, key string, must bool) {
	if val := os.Getenv(key); val != "" {
		*dst = val
		return
	}
	if must {
		handleMust(key)
	}
}

// LoadInt64 加载环境变量，must 设置是否为必须，must=true 时
// // 如果环境变量不存在则会 panic。环境变量的值如果不为整形则会 panic
func LoadInt64(dst *int64, key string, must bool) {
	if val := os.Getenv(key); val != "" {
		var err error
		*dst, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic(fmt.Sprint("解析环境变量: ", key, " 失败，", err))
		}
		return
	}
	if must {
		handleMust(key)
	}
}

// LoadInt 加载环境变量，must 设置是否为必须
func LoadInt(dst *int, key string, must bool) {
	if val := os.Getenv(key); val != "" {
		var err error
		*dst, err = strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprint("解析环境变量: ", key, " 失败，", err))
		}
		return
	}
	if must {
		handleMust(key)
	}
}
