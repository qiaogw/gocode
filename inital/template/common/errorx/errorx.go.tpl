package errorx

type CodeError struct {
Code uint32      `json:"code"`
Msg  string      `json:"msg"`
Data interface{} `json:"data,omitempty"`
}

type CodeErrorResponse struct {
Code uint32      `json:"code"`
Msg  string      `json:"msg"`
Data interface{} `json:"data,omitempty"`
}

func NewCodeError(code uint32, msg string) error {
return &CodeError{Code: code, Msg: msg}
}

func NewErrorData(code uint32, msg string, data interface{}) error {
return &CodeError{Code: code, Msg: msg, Data: data}
}

func NewDefaultError(msg string) error {
return NewCodeError(ServerCommonError, msg)
}

// GetErrCode 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
return e.Code
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
return e.Msg
}

func (e *CodeError) Error() string {
return e.Msg
}

func (e *CodeError) Info() *CodeErrorResponse {
return &CodeErrorResponse{
Code: e.Code,
Msg:  e.Msg,
Data: e.Data,
}
}
