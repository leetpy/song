package errorx

import "fmt"

type CodeError struct {
	Code int
	Msg  string
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func (e *CodeError) GetCode() int {
	return e.Code
}

func (e *CodeError) GetMsg() string {
	return e.Msg
}

// 创建错误
func NewCodeError(code int, msg string) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}

// 使用预定义错误码
func NewCodeErrorWithCode(code int) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  GetErrMsg(code),
	}
}
