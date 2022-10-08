
type CommonResponse {
    Code uint32 `json:"code"`
    Data interface{} `json:"data,omitempty"`
    Msg string `json:"msg"`
}