package errx

import "fmt"

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32 `json:"code"`
	errMsg  string `json:"msg"`
}

// GetErrCode 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d,ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewDefaultError(errMsg string) *CodeError {
	return &CodeError{errCode: ServerCommonError, errMsg: errMsg}
}
func NewErrorf(errCode uint32, format string, args ...interface{}) *CodeError {
	return &CodeError{errCode: errCode, errMsg: fmt.Sprintf(format, args...)}
}
