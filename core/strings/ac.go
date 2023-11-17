// author gmfan
// date 2023/3/27

package strings

type (
	// AC 使用 UTF-8 编码
	AC struct {
		root *trieNode
	}

	trieNode struct {
		// 字符
		c        rune
		children map[rune]*trieNode
		// 字符串长度，当字符串不存在时为 0 ，存在则为字符串长度
		len int
		// 失配指针，指向最长后缀节点
		fail *trieNode
	}
)

func NewAC() *AC {
	return &AC{
		root: &trieNode{
			children: make(map[rune]*trieNode),
		},
	}
}

// addWord 字典树添加单词 world
func (t *trieNode) addWord(world []rune) {
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
				c:        c,
				children: make(map[rune]*trieNode),
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
		a.root.addWord([]rune(word))
	}

	// 重新构建失配指针
	a.buildFail()
}

// 是否存在子节点
func (t *trieNode) hasChild(k rune) bool {
	_, ok := t.children[k]
	return ok
}

// 获取存在子节点 c 的失配指针
func (a *AC) getFailHasChild(fail *trieNode, c rune) *trieNode {
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
// 匹配的下标，不存在则为 -1，值得注意的是这里的 index 下标返回的是 []rune(text) 的下标。
func (a *AC) FindFirst(text string) (word string, index int) {
	rs := []rune(text)
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
			index = i - cur.len + 1
			return string(rs[index : i+1]), index
		}
	}
	return word, -1
}

// WrapByFn 将文本中所有匹配的单词使用 fn 函数包裹起来后返回 text 处理结果
func (a *AC) WrapByFn(text string, fn func(word string) string) (resText string, words []string) {
	bitMap, words := a.BitMap(text)
	if len(words) == 0 {
		// 没有模式串匹配到
		return text, words
	}

	rs := []rune(text)
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
				resText += string(rs[idx:l])
				idx = r
			}
			resText += fn(string(rs[l:r]))
			idx = r
		}
		l = r
	}
	resText += string(rs[idx:])
	return
}

// BitMap 获取 text 的位图。
// 示例：
// text="hello"
// 模式串="llo"
// 结果：
// bitMap=[false,false,true,true,true]
// words=["llo"]
func (a *AC) BitMap(text string) (bitMap []bool, words []string) {
	rs := []rune(text)
	// 位图
	bitMap = make([]bool, len(rs))
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
			fillBitMap(bitMap, left, i, true)
			words = append(words, string(rs[left:i+1]))
		}

		// 失配指针可能还存在匹配模式串
		fail := cur.fail
		for fail != nil {
			if fail.len > 0 {
				// 匹配成功
				left := i - fail.len + 1
				fillBitMap(bitMap, left, i, true)
				words = append(words, string(rs[left:i+1]))
			}
			fail = fail.fail
		}
	}
	return
}

func fillBitMap(bitMap []bool, left, right int, flag bool) {
	for i := left; i <= right; i++ {
		bitMap[i] = flag
	}
}
