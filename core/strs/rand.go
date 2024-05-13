// author gmfan
// date 2023/2/25
package strs

import (
	"crypto/rand"
	"fmt"
	"io"
	math_rand "math/rand"
)

const CharacterSequence = "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
const N = 62

// Rand 生成随机 token 字符串，len 为 token 长度。
// 随机字符从 CharacterSequence 选取。CharacterSequence 含有 62 个字符
func Rand(len int) string {
	// 获取随机序列
	ridx := make([]byte, len)
	_, err := io.ReadFull(rand.Reader, ridx)
	if err != nil {
		panic(err)
	}

	// 随机序列转换成 byte 切片
	bs := make([]byte, len)
	for i := 0; i < len; i++ {
		bs[i] = CharacterSequence[ridx[i]%N]
	}
	return string(bs)
}

// RandomUA 生成随机 UA 字符串
func RandomUA() string {
	ua := fmt.Sprintf("AppleWebKit/%d.%d  (KHTML, like Gecko) Chrome/69.0.%d.%d",
		math_rand.Intn(100)+500, math_rand.Intn(100)+1, math_rand.Intn(500)+3000, math_rand.Intn(100)+1)
	ua = fmt.Sprintf(" Safari/%d.%d %s",
		math_rand.Intn(100)+500, math_rand.Intn(100)+1, ua)
	ua = fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_1%d_%d) %s",
		math_rand.Intn(3)+3, math_rand.Intn(7)+1, ua)
	return ua
}
