// author gmfan
// date 2023/2/24

package errs

import (
	"errors"
	"fmt"
	"github.com/tkgfan/got/core/structs"
)

// Cause 返回由 errors.New、errors.Wrap、errors.Wrapf 中包裹的 Cause error。
// error 如果不是上述方法产生的则会返回其本身
func Cause(err error) error {
	var e *stackError
	if errors.As(err, &e) {
		return e.Cause
	}
	return err
}

// Wrap 返回包含堆栈信息的 error
func Wrap(err error) error {
	if structs.IsNil(err) {
		return nil
	}
	var se *stackError
	if errors.As(err, &se) {
		se.Stacks = append(se.Stacks, caller(""))
		return se
	}
	return &stackError{
		Cause:  err,
		Stacks: []*stack{caller("")},
	}
}

// Wrapf 返回包含堆栈信息的 error。format 格式化信息会保存到堆栈信息中
func Wrapf(err error, format string, args ...any) error {
	if structs.IsNil(err) {
		return nil
	}

	remark := fmt.Sprintf(format, args...)
	var st *stackError
	if errors.As(err, &st) {
		st.Stacks = append(st.Stacks, caller(remark))
		return st
	}

	return &stackError{
		Cause:  err,
		Stacks: []*stack{caller(remark)},
	}
}
