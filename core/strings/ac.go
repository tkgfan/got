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
		c     rune
		child map[rune]*trieNode
		// 字符串长度，当字符串不存在时为 0 ，存在则为字符串长度
		len  int
		fail *trieNode
	}
)

// 是否存在子节点
func (t *trieNode) hasChild(k rune) bool {
	_, ok := t.child[k]
	return ok
}

func NewAC() *AC {
	return &AC{
		root: &trieNode{
			child: make(map[rune]*trieNode),
		},
	}
}

// AddWords 将单词添加到 AC 自动机中
func (a *AC) AddWords(words ...string) {
	for _, word := range words {
		cs := []rune(word)
		p := a.root
		for _, c := range cs {
			if v, ok := p.child[c]; ok {
				p = v
			} else {
				p.child[c] = &trieNode{
					c:     c,
					child: make(map[rune]*trieNode),
				}
				p = p.child[c]
			}
		}
		p.len = len(cs)
	}

	// 重新构建失配指针
	a.buildFail()
}

// buildFail 构建失配指针
func (a *AC) buildFail() {
	var que []*trieNode
	// 将 root 的所有子节点加入到队列中
	for _, v := range a.root.child {
		v.fail = a.root
		que = append(que, v)
	}

	// 构建失配指针
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		for r, node := range p.child {
			f := a.hasRuneFail(p.fail, r)
			if f == nil {
				node.fail = a.root
			} else {
				node.fail = f.child[r]
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
		cur = cur.child[r]
		// 当前字符不匹配使用失配指针跳转到最近一个可匹配位置
		if cur == nil {
			f = a.hasRuneFail(f, r)
			if f != nil {
				cur = f.child[r]
			}
		}

		// 从头开始
		if cur == nil {
			cur = a.root
			// 匹配到单词
		} else if cur.len > 0 {
			index = i - cur.len + 1
			return string(rs[index : i+1]), index
		}
	}
	return word, -1
}

// 获取第一个 child 中存在 r 的失配指针，失配指针不存在则返回 nil
func (a *AC) hasRuneFail(fail *trieNode, r rune) *trieNode {
	for fail != nil {
		if fail.hasChild(r) {
			return fail
		}
		fail = fail.fail
	}
	return nil
}
