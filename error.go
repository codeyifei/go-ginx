package ginx

import (
	"github.com/pkg/errors"
)

type Error struct {
	code int
	err  error
}

func (e *Error) Unwrap() error {
	return errors.Unwrap(e.err)
}

func (e *Error) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

var _ error = (*Error)(nil)

// NewError 创建一个ginx错误
// Example：
// err := ginx.NewError(101, "参数错误")
func NewError(code int, message string) *Error {
	return &Error{
		code: code,
		err:  errors.New(message),
	}
}

// WrapError 包装成ginx错误
// Example:
// err = ginx.WrapError(err, 101, "参数错误")
func WrapError(err error, code int, message string) *Error {
	return &Error{
		code: code,
		err:  errors.Wrap(err, message),
	}
}
