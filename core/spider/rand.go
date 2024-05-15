// author lby
// date 2024/5/9

package spider

import (
	"fmt"
	"math/rand"
)

// RandomUA 生成随机 UA 字符串
func RandomUA() string {
	ua := fmt.Sprintf("AppleWebKit/%d.%d  (KHTML, like Gecko) Chrome/69.0.%d.%d",
		rand.Intn(100)+500, rand.Intn(100)+1, rand.Intn(500)+3000, rand.Intn(100)+1)
	ua = fmt.Sprintf(" Safari/%d.%d %s",
		rand.Intn(100)+500, rand.Intn(100)+1, ua)
	ua = fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_1%d_%d) %s",
		rand.Intn(3)+3, rand.Intn(7)+1, ua)
	return ua
}
