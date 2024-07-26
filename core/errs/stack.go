// author gmfan
// date 2023/2/24

package errs

import (
	"encoding/json"
	stderrors "errors"
	"runtime"
)

// 堆栈信息结构
type stack struct {
	pc     uintptr
	File   string `json:"file"`
	Line   int    `json:"line"`
	Remark string `json:"remark"`
}

func caller(remark string) *stack {
	st := &stack{
		Remark: remark,
	}
	st.pc, st.File, st.Line, _ = runtime.Caller(2)
	return st
}

// stackError 包含错误的堆栈信息
type stackError struct {
	Cause  error    `json:"cause"`
	Stacks []*stack `json:"stacks"`
}

func (s *stackError) MarshalJSON() ([]byte, error) {
	cause := ""
	if s.Cause != nil {
		cause = s.Cause.Error()
	}
	m := map[string]any{
		"cause":  cause,
		"stacks": s.Stacks,
	}
	return json.Marshal(m)
}

// New 创建一个包含堆栈信息的 error
func New(msg string) error {
	return &stackError{
		Cause:  stderrors.New(msg),
		Stacks: []*stack{caller("")},
	}
}

func (s *stackError) Error() string {
	return s.Cause.Error()
}
