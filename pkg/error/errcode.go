package error

import (
	"fmt"
	"net/http"
)

type Error struct {
	code     int
	msg      string
	details  []string
	httpCode int
}

func NewError(code int, msg string) *Error {
	return &Error{
		code:     code,
		msg:      msg,
		httpCode: http.StatusOK,
	}
}

func NewErrWithBusinessError(error error) *Error {
	errMsg := "business error"
	if _, ok := error.(*BusinessError); ok {
		errMsg = error.Error()
	}

	return &Error{
		code:     500,
		msg:      errMsg,
		httpCode: http.StatusOK,
	}
}

func (e *Error) String() string {
	return fmt.Sprintf("错误码:%d,错误信息:%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) HttpCode() int {
	return e.httpCode
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// 添加额外信息
func (e *Error) WithDetails(details ...string) *Error {
	e.details = make([]string, 0, len(details))
	for _, detail := range details {
		e.details = append(e.details, detail)
	}

	return e
}

func (e *Error) WithHttpCode(code int) *Error {
	e.httpCode = code

	return e
}

func (e *Error) WithMsg(msg string) *Error {
	e.msg = msg

	return e
}
