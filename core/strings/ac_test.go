// author lby
// date 2023/3/29

package strings

import (
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
		ac := NewAC()
		ac.AddWords(tt.words...)
		word, index := ac.FindFirst(tt.text)
		if word != tt.expectWord || index != tt.expectIndex {
			t.Errorf("%s, expect[word=%s,index=%d],got[word=%s,index=%d]", tt.name, tt.expectWord, tt.expectIndex, word, index)
		}
	}
}
