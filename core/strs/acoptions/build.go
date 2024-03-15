// author lby
// date 2024/2/28

package acoptions

type BuildOptions struct {
	// 忽略大小写
	IgnoreCase *bool
	// 分割符
	Separator *string
}

func NewBuildOptions() *BuildOptions {
	ignoreCase := false
	separator := ""
	return &BuildOptions{
		IgnoreCase: &ignoreCase,
		Separator:  &separator,
	}
}

// MergeBuildOptions 合并参数
func MergeBuildOptions(opts ...*BuildOptions) *BuildOptions {
	res := NewBuildOptions()
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.IgnoreCase != nil {
			res.IgnoreCase = opt.IgnoreCase
		}
		if opt.Separator != nil {
			res.Separator = opt.Separator
		}
	}
	return res
}
