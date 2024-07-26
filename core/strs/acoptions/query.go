// author gmfan
// date 2024/2/28

package acoptions

type QueryOptions struct {
	// 分割符
	Separator *string
	// 字符替换。匹配时 key 会被替换为 value
	CharacterReplace map[byte]byte
	// 句子分割符
	SentenceSeparators []rune
}

func NewQueryOptions() *QueryOptions {
	separator := ""
	return &QueryOptions{
		Separator:        &separator,
		CharacterReplace: make(map[byte]byte),
	}
}

func MergeQueryOptions(opts ...*QueryOptions) *QueryOptions {
	res := NewQueryOptions()
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.Separator != nil {
			res.Separator = opt.Separator
		}
		if opt.CharacterReplace != nil {
			for k, v := range opt.CharacterReplace {
				res.CharacterReplace[k] = v
			}
		}
		if opt.SentenceSeparators != nil {
			res.SentenceSeparators = append(res.SentenceSeparators, opt.SentenceSeparators...)
		}
	}

	if len(res.SentenceSeparators) > 0 {
		// 句子分割符去重
		set := make(map[rune]struct{})
		for _, v := range res.SentenceSeparators {
			set[v] = struct{}{}
		}
		res.SentenceSeparators = make([]rune, 0, len(set))
		for k := range set {
			res.SentenceSeparators = append(res.SentenceSeparators, k)
		}
	}
	return res
}
