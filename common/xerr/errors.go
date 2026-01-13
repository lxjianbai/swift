package xerr

import "fmt"

type CodeError struct {
	code int
	msg  string
}

// 核心：实现此接口即视为“可信任的业务错误”
func (e *CodeError) Code() int       { return e.code }
func (e *CodeError) Message() string { return e.msg }
func (e *CodeError) Error() string   { return fmt.Sprintf("Code:%d, Msg:%s", e.code, e.msg) }

func New(code int, msg string) error {
	return &CodeError{code: code, msg: msg}
}

func FromCode(code int, customMsg ...string) error {
	msg := MapErrMsg(code)
	if len(customMsg) > 0 {
		msg = customMsg[0]
	}
	return &CodeError{code: code, msg: msg}
}
