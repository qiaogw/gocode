package errx

// OK 成功返回
const OK uint32 = 200

// 前3位代表业务,后三位代表具体功能

// 全局错误码
const (
	ServerCommonError       uint32 = 100001
	RequestParamError       uint32 = 100002
	TokenExpireError        uint32 = 100003
	TokenGenerateError      uint32 = 100004
	DbError                 uint32 = 100005
	NoData                  uint32 = 100006
	Duplicate               uint32 = 100007
	PrimaryError            uint32 = 100008
	Success                 uint32 = 0 //请求成功
	ServerRpcError          uint32 = 500
	ErrUsernamePwdError     uint32 = 110000
	ErrGenerateTokenError   uint32 = 110001
	ErrAuth                 uint32 = 401
	ErrReq                  uint32 = 100009
	ErrTimeout              uint32 = 110
	FileOrDirectoryNotExist uint32 = 100010
)
