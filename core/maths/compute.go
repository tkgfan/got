// author gmfan
// date 2024/5/8

package maths

import "math"

// RatioToAB 将 ra 比 rb 转换为 a，b 值范围内 resA/resB 损失最小。
func RatioToAB(ra, rb, a, b int) (resA, resB int) {
	if ra == 0 || rb == 0 || a == 0 || b == 0 {
		return 0, 0
	}
	t := float64(ra) / float64(rb)
	c := math.Abs(t - 1.0)
	resA = 1
	resB = 1
	for i := 1; i < a; i++ {
		for j := 1; j < b; j++ {
			tt := float64(i) / float64(j)
			if math.Abs(tt-t) < c {
				c = math.Abs(tt - t)
				resA = i
				resB = j
			}
		}
	}
	return
}
