// author gmfan
// date 2024/2/18
package errors

import stderrors "errors"

// Unwrap 标准库
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}

// Is 标准库
func Is(err, target error) bool {
	return stderrors.Is(err, target)
}

// As 标准库
func As(err error, target any) bool {
	return stderrors.As(err, target)
}
