// author gmfan
// date 2023/3/27

package strs

import (
	"github.com/tkgfan/got/core/slices"
	"github.com/tkgfan/got/core/strs/acoptions"
	"strings"
	"unicode/utf8"
)

type (
	// AC 使用 UTF-8 编码
	AC struct {
		root         *trieNode
		buildOptions *acoptions.BuildOptions
	}

	trieNode struct {
		children map[string]*trieNode
		// 字符串长度，当字符串不存在时为 0 ，存在则为字符串长度
		len int
		// 失配指针，指向最长后缀节点
		fail *trieNode
	}
)

func NewAC(opts ...*acoptions.BuildOptions) *AC {
	return &AC{
		root: &trieNode{
			children: make(map[string]*trieNode),
		},
		buildOptions: acoptions.MergeBuildOptions(opts...),
	}
}

// addWord 字典树添加单词 world，单词按照 string 来分割字符
func (t *trieNode) addWord(world []string) {
	// p 作为根节点
	p := t
	// 遍历所有字符，依次加入字典树中
	for _, c := range world {
		if v, ok := p.children[c]; ok {
			// 子节点已存在
			p = v
		} else {
			// 节点不存在，创建新子节点
			p.children[c] = &trieNode{
				children: make(map[string]*trieNode),
			}
			p = p.children[c]
		}
	}
	// 设置节点单词长度
	p.len = len(world)
}

// AddWords 将单词添加到 AC 自动机中
func (a *AC) AddWords(words ...string) {
	// 将单词添加到字典树中
	for _, word := range words {
		if *a.buildOptions.IgnoreCase {
			// 忽略大小写时，将单词转小写
			word = strings.ToLower(word)
		}
		// 将单词分割
		arr := strings.Split(word, *a.buildOptions.Separator)
		a.root.addWord(arr)
	}

	// 重新构建失配指针
	a.buildFail()
}

// 获取存在子节点 c 的失配指针
func (a *AC) getFailHasChild(fail *trieNode, c string) *trieNode {
	for fail != nil {
		if _, ok := fail.children[c]; ok {
			return fail
		}
		fail = fail.fail
	}
	return nil
}

// buildFail 构建失配指针
func (a *AC) buildFail() {
	var que []*trieNode
	// 将 root 的所有子节点加入到队列中
	for _, v := range a.root.children {
		v.fail = a.root
		que = append(que, v)
	}

	// 构建失配指针
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		for r, node := range p.children {
			// 获取子节点存在 r 的失配指针
			f := a.getFailHasChild(p.fail, r)
			if f == nil {
				node.fail = a.root
			} else {
				node.fail = f.children[r]
			}

			que = append(que, node)
		}
	}
}

// FindFirst 获取本文中匹配的第一个单词，如果文本中不存在匹配单词则返回空。index 为第一个
// 匹配的下标，不存在则为 -1。
func (a *AC) FindFirst(text string, opts ...*acoptions.QueryOptions) (word string, index int) {
	qOpt := acoptions.MergeQueryOptions(opts...)
	originRs := strings.Split(text, *qOpt.Separator)
	var rs []string
	if *a.buildOptions.IgnoreCase {
		// 忽略大小写
		for i := 0; i < len(originRs); i++ {
			rs = append(rs, strings.ToLower(originRs[i]))
		}
	} else {
		rs = originRs
	}
	cur := a.root
	for i, r := range rs {
		f := cur.fail
		cur = cur.children[r]
		// 当前字符不匹配使用失配指针跳转到最近一个可匹配位置
		if cur == nil {
			f = a.getFailHasChild(f, r)
			if f != nil {
				cur = f.children[r]
			}
		}

		// 从头开始
		if cur == nil {
			cur = a.root
			// 匹配到单词
		} else if cur.len > 0 {
			// 匹配成功
			left := i - cur.len + 1
			// 计算 index
			index = a.computeOriginTextIdx(rs, left, len(*qOpt.Separator))
			word = strings.Join(originRs[left:i+1], *qOpt.Separator)
			return
		}
	}
	return word, -1
}

// 计算原文本中下标地址，strs 是原文本根据分割符分割的数组，separatorLen 是
// 分割符长度。
func (a *AC) computeOriginTextIdx(strs []string, i, separatorLen int) int {
	res := 0
	for j := 0; j < i; j++ {
		res += len(strs[j]) + separatorLen
	}
	return res
}

// WrapByFn 将文本中所有匹配的单词使用 fn 函数包裹起来后返回 text 处理结果
func (a *AC) WrapByFn(text string, fn func(word string) string, opts ...*acoptions.QueryOptions) (resText string, words []string) {
	bitMap, words := a.BitMap(text, acoptions.MergeQueryOptions(opts...))
	if len(words) == 0 {
		// 没有模式串匹配到
		return text, words
	}

	rs := text
	idx, l, r := 0, 0, 0
	for l < len(bitMap) && r < len(bitMap) {
		for l < len(bitMap) && !bitMap[l] {
			l++
		}
		r = l
		for r < len(bitMap) && bitMap[r] {
			r++
		}
		if (r-l) > 1 && r <= len(bitMap) {
			if idx < l {
				resText += rs[idx:l]
				idx = r
			}
			resText += fn(rs[l:r])
			idx = r
		}
		l = r
	}
	resText += rs[idx:]
	return
}

// BitMap 获取 text 的位图。
// 示例：
// text="hello"
// 模式串="llo"
// 结果：
// bitMap=[false,false,true,true,true]
// words=["llo"]
func (a *AC) BitMap(text string, opts ...*acoptions.QueryOptions) (bitMap []bool, words []string) {
	opt := acoptions.MergeQueryOptions(opts...)
	if len(opt.CharacterReplace) > 0 {
		// 字符替换
		bs := []byte(text)
		for i := 0; i < len(bs); i++ {
			if v, ok := opt.CharacterReplace[bs[i]]; ok {
				bs[i] = v
			}
		}
		text = string(bs)
	}

	if len(opt.SentenceSeparators) > 0 {
		// 句子分割
		set := make(map[rune]bool)
		for _, r := range opt.SentenceSeparators {
			set[r] = true
		}
		rs := []rune(text)
		l, r := 0, 0
		for r < len(rs) {
			if !set[rs[r]] {
				// 不是分割符
				r++
				continue
			}

			if l < r-1 {
				// 分割符分割的句子
				sBitMap, sWords := a.sentenceBitMap(string(rs[l:r]), opt)
				bitMap = append(bitMap, sBitMap...)
				words = append(words, sWords...)
			}
			bitMap = append(bitMap, make([]bool, utf8.RuneLen(rs[r]))...)
			r++
			l = r
		}

		if r > l {
			sBitMap, sWords := a.sentenceBitMap(string(rs[l:r]), opt)
			bitMap = append(bitMap, sBitMap...)
			words = append(words, sWords...)
		}
		return bitMap, slices.DuplicationStrs(words)
	}
	// 不需要进行句子分割，直接处理文本即可
	return a.sentenceBitMap(text, opt)
}

// 生成句子位图，并返回匹配的单词
func (a *AC) sentenceBitMap(sentence string, opt *acoptions.QueryOptions) (bitMap []bool, words []string) {
	sepLen := len(*opt.Separator)
	originRs := strings.Split(sentence, *opt.Separator)
	// rs 用于匹配
	var rs []string
	if *a.buildOptions.IgnoreCase {
		// 忽略大小写
		for i := 0; i < len(originRs); i++ {
			rs = append(rs, strings.ToLower(originRs[i]))
		}
	} else {
		rs = originRs
	}
	// 位图
	bitMap = make([]bool, len(sentence))
	// 开始匹配
	cur := a.root
	for i, r := range rs {
		for cur.children[r] == nil && cur != a.root {
			cur = cur.fail
		}
		if cur.children[r] != nil {
			cur = cur.children[r]
		}
		if cur.len > 0 {
			// 匹配成功
			left := i - cur.len + 1
			words = append(words, strings.Join(originRs[left:i+1], *opt.Separator))
			bitLeft := a.computeOriginTextIdx(originRs, left, sepLen)
			bitRight := a.computeOriginTextIdx(originRs, i, sepLen)
			// 处理右边界时还要加上单词本身的长度，需要注意本身长度为一时就不需要了
			if len(originRs[i]) > 1 {
				bitRight += len(originRs[i]) - 1
			}
			fillBitMap(bitMap, bitLeft, bitRight, true)
		}

		// 失配指针可能还存在匹配模式串
		fail := cur.fail
		for fail != nil {
			if fail.len > 0 {
				// 匹配成功
				left := i - fail.len + 1
				words = append(words, strings.Join(originRs[left:i+1], *opt.Separator))
				bitLeft := a.computeOriginTextIdx(originRs, left, sepLen)
				bitRight := a.computeOriginTextIdx(originRs, i, sepLen)
				// 处理右边界时还要加上单词本身的长度，需要注意本身长度为一时就不需要了
				if len(originRs[i]) > 1 {
					bitRight += len(originRs[i]) - 1
				}
				fillBitMap(bitMap, bitLeft, bitRight, true)
			}
			fail = fail.fail
		}
	}
	// 结果去重
	words = slices.DuplicationStrs(words)
	return
}

func fillBitMap(bitMap []bool, left, right int, flag bool) {
	for i := left; i <= right; i++ {
		bitMap[i] = flag
	}
}
