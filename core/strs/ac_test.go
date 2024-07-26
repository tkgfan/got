// author gmfan
// date 2023/3/29

package strs

import (
	"github.com/tkgfan/got/core/strs/acoptions"
	"reflect"
	"testing"
)

func TestAC_FindFirst(t *testing.T) {
	type arg struct {
		ignoreCase     bool
		buildSeparator string
		querySeparator string
		words          []string
		text           string
	}
	tests := []struct {
		name        string
		arg         arg
		expectWord  string
		expectIndex int
	}{
		{
			name: "字母用例-1",
			arg: arg{
				words: []string{"issip"},
				text:  "mississippi",
			},
			expectWord:  "issip",
			expectIndex: 4,
		},
		{
			name: "字母用例-2",
			arg: arg{
				words: []string{"man"},
				text:  "hellomawnorld",
			},
			expectWord:  "",
			expectIndex: -1,
		},
		{
			name: "中文用例-1",
			arg: arg{
				words: []string{"蜜蜂", "蜂蜜"},
				text:  "蜜蜜蜜蜜蜂与蜂蜜的关系",
			},
			expectWord:  "蜜蜂",
			expectIndex: 9,
		},
		{
			name: "特殊字符-1",
			arg: arg{
				words: []string{"OS-A & AKR"},
				text:  "OS-A & AKR",
			},
			expectWord:  "OS-A & AKR",
			expectIndex: 0,
		},
		{
			name: "分割符测试",
			arg: arg{
				buildSeparator: " ",
				querySeparator: " ",
				words:          []string{"hello world"},
				text:           "  apple hello world",
			},
			expectWord:  "hello world",
			expectIndex: 8,
		},
		{
			name: "分割符测试-2",
			arg: arg{
				buildSeparator: " ",
				querySeparator: " ",
				words:          []string{"hello world"},
				text:           " apple hello world",
			},
			expectWord:  "hello world",
			expectIndex: 7,
		},
		{
			name: "特殊字符分割符测试",
			arg: arg{
				buildSeparator: "，",
				querySeparator: "，",
				words:          []string{"你好，世界"},
				text:           "中中中，你好，世界",
			},
			expectWord:  "你好，世界",
			expectIndex: 12,
		},
		{
			name: "忽略大小写用例",
			arg: arg{
				ignoreCase: true,
				words:      []string{"mAN"},
				text:       "hellomAnorld",
			},
			expectWord:  "mAn",
			expectIndex: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := NewAC(&acoptions.BuildOptions{
				IgnoreCase: &tt.arg.ignoreCase,
				Separator:  &tt.arg.buildSeparator,
			})
			ac.AddWords(tt.arg.words...)
			word, index := ac.FindFirst(tt.arg.text, &acoptions.QueryOptions{
				Separator: &tt.arg.querySeparator,
			})
			if word != tt.expectWord || index != tt.expectIndex {
				t.Errorf("%s, expect[word=%s,index=%d],got[word=%s,index=%d]", tt.name, tt.expectWord, tt.expectIndex, word, index)
			}
		})
	}
}

func TestAC_BitMap(t *testing.T) {
	type arg struct {
		words              []string
		text               string
		characterReplace   map[byte]byte
		sentenceSeparators []rune
		buildSeparator     string
		querySeparator     string
	}
	tests := []struct {
		name   string
		arg    arg
		bitMap []bool
	}{
		{
			name: "普通测试",
			arg: arg{
				words: []string{"llo", "el"},
				text:  "hello",
			},
			bitMap: []bool{false, true, true, true, true},
		},
		{
			name: "替换符测试",
			arg: arg{
				words:            []string{"hello world"},
				text:             "hello\nworld",
				buildSeparator:   " ",
				querySeparator:   " ",
				characterReplace: map[byte]byte{'\n': ' '},
			},
			bitMap: []bool{true, true, true, true, true, true, true, true, true, true, true},
		},
		{
			name: "替换符加句子分割测试",
			arg: arg{
				words:              []string{"he llo", "w"},
				text:               " he\nllo，w ",
				buildSeparator:     " ",
				querySeparator:     " ",
				characterReplace:   map[byte]byte{'\n': ' '},
				sentenceSeparators: []rune{'，', ','},
			},
			bitMap: []bool{false, true, true, true, true, true, true, false, false, false, true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := NewAC(&acoptions.BuildOptions{
				Separator: &tt.arg.buildSeparator,
			})
			ac.AddWords(tt.arg.words...)
			bitMap, _ := ac.BitMap(tt.arg.text, &acoptions.QueryOptions{
				Separator:          &tt.arg.querySeparator,
				CharacterReplace:   tt.arg.characterReplace,
				SentenceSeparators: tt.arg.sentenceSeparators,
			})
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
