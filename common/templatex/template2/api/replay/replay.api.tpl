
type CommonResponse {
    Code uint32 `json:"code"`
    Data interface{} `json:"data,omitempty"`
    Msg string `json:"msg"`
}

//导入请求
type ImportRequest {
    UpFile interface{} `json:"upFile"`
}
type ExportResponse {
    Byte []byte `json:"byte"`
}
// 空请求
type NullRequest {}

// 空回复
type NullResponse {}
