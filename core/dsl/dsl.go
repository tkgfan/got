// author lby
// date 2024/2/19

package dsl

import "encoding/json"

type DSL map[string]any

func New() DSL {
	return map[string]any{}
}

func WrapDSL(dsl map[string]any) DSL {
	return dsl
}

func (t DSL) String() string {
	bs, _ := json.Marshal(t)
	return string(bs)
}

// Push 在指定位置添加 value，positions 为位置参数
func (t DSL) Push(value any, positions ...string) DSL {
	m := t
	for i := 0; i < len(positions); i++ {
		p := positions[i]
		if i == len(positions)-1 {
			m[p] = value
			break
		}
		if _, ok := m[p]; !ok {
			m[p] = make(map[string]any)
		}
		if v, ok := m[p].(map[string]any); ok {
			m = v
		} else {
			d := m[p].(DSL)
			m = d
		}
	}
	return t
}

// GetDSL 获取指定位置的 DSL ，如果位置不存在则会创建 DSL
func (t DSL) GetDSL(positions ...string) DSL {
	res := t
	for i := 0; i < len(positions)-1; i++ {
		p := positions[i]
		if _, ok := res[p]; !ok {
			res[p] = New()
		}
		res = res[p].(DSL)
	}
	return res
}

// PushToArray 添加一个元素至数组中，指定位置不存在则会自动创建 any 类型数组
func (t DSL) PushToArray(element any, positions ...string) DSL {
	m := t.GetDSL(positions...)
	end := positions[len(positions)-1]
	if _, ok := m[end]; !ok {
		m[end] = make([]any, 0)
	}
	m[end] = append(m[end].([]any), element)

	return t
}
