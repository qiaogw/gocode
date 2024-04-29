package errx

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[NoData] = "数据不存在"
	message[Duplicate] = "数据或主键重复,请检查输入数据"
	message[ServerRpcError] = "服务器开小差啦,稍后再来试一试"
	message[PrimaryError] = "不支持无主键或联合主键"

	message[ErrGenerateTokenError] = "生成token失败"
	message[ErrUsernamePwdError] = "账号或密码不正确"
	message[ErrAuth] = "无访问权限"
	message[ErrReq] = "数据请求错误"
	message[ErrTimeout] = "服务器响应超时"
	message[FileOrDirectoryNotExist] = "文件或目录不存在"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
