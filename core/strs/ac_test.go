// author lby
// date 2023/3/29

package strs

import (
	"reflect"
	"testing"
)

func TestAC_FindFirst(t *testing.T) {
	tests := []struct {
		name        string
		words       []string
		text        string
		expectWord  string
		expectIndex int
	}{
		{
			name:        "字母用例-1",
			words:       []string{"issip"},
			text:        "mississippi",
			expectWord:  "issip",
			expectIndex: 4,
		},
		{
			name:        "字母用例-2",
			words:       []string{"man"},
			text:        "hellomawnorld",
			expectWord:  "",
			expectIndex: -1,
		},
		{
			name:        "中文用例-1",
			words:       []string{"蜜蜂", "蜂蜜"},
			text:        "蜜蜜蜜蜜蜂与蜂蜜的关系",
			expectWord:  "蜜蜂",
			expectIndex: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := NewAC()
			ac.AddWords(tt.words...)
			word, index := ac.FindFirst(tt.text)
			if word != tt.expectWord || index != tt.expectIndex {
				t.Errorf("%s, expect[word=%s,index=%d],got[word=%s,index=%d]", tt.name, tt.expectWord, tt.expectIndex, word, index)
			}
		})
	}
}

func TestAC_BitMap(t *testing.T) {
	tests := []struct {
		name     string
		patterns []string
		text     string
		bitMap   []bool
	}{
		{
			name:     "普通测试",
			patterns: []string{"llo", "el"},
			text:     "hello",
			bitMap:   []bool{false, true, true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			ac := NewAC()
			ac.AddWords(tt.patterns...)
			bitMap, _ := ac.BitMap(tt.text)
			if !reflect.DeepEqual(bitMap, tt.bitMap) {
				t.Errorf("got: %+v, expect: %+v", bitMap, tt.bitMap)
			}
		})
	}
}

func TestAC_WrapByFn(t *testing.T) {
	tests := []struct {
		name     string
		patterns []string
		text     string
		fn       func(word string) string
		resText  string
		words    []string
	}{
		{
			name:     "普通测试",
			patterns: []string{"he", "rl", "ld"},
			text:     "hello world",
			fn: func(word string) string {
				return "*" + word + "*"
			},
			resText: "*he*llo wo*rld*",
			words:   []string{"he", "rl", "ld"},
		},
		{
			name:     "普通测试-2",
			patterns: []string{"蜜蜂", "蜂蜜"},
			text:     "蜜蜜蜜蜜蜂与蜂蜜的关系",
			fn: func(word string) string {
				return "*" + word + "*"
			},
			resText: "蜜蜜蜜*蜜蜂*与*蜂蜜*的关系",
			words:   []string{"蜜蜂", "蜂蜜"},
		},
		{
			name:     "普通测试-3",
			patterns: []string{"NIKE"},
			text:     "NIKE fdjksfjie",
			fn: func(word string) string {
				return "*" + word + "*"
			},
			resText: "*NIKE* fdjksfjie",
			words:   []string{"NIKE"},
		},
		{
			name:     "普通测试-4",
			patterns: []string{"NIKE", " NIKE123"},
			text:     " NIKE",
			fn: func(word string) string {
				return "*" + word + "*"
			},
			resText: " *NIKE*",
			words:   []string{"NIKE"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := NewAC()
			ac.AddWords(tt.patterns...)
			resText, words := ac.WrapByFn(tt.text, tt.fn)
			if resText != tt.resText {
				t.Errorf("got resText: %s , expect resText: %s", resText, tt.resText)
				return
			}
			if !reflect.DeepEqual(words, tt.words) {
				t.Errorf("got words: %+v, expect words: %+v", words, tt.words)
				return
			}
		})
	}
}
