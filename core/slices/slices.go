// author gmfan
// date 2023/12/23
package slices

// DuplicationStrs 去重 strs 数组并且返回结果不改变原顺序。
func DuplicationStrs(strs []string) []string {
	set := make(map[string]struct{})
	for i := 0; i < len(strs); i++ {
		if _, ok := set[strs[i]]; ok {
			strs = append(strs[:i], strs[i+1:]...)
			i--
		}
		set[strs[i]] = struct{}{}
	}
	return strs
}
