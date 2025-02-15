package jwtx

import (
	"github.com/golang-jwt/jwt"
)

// 定义 JWT 上下文中使用的键常量
const (
	CtxKeyJwtAccessExpire = "exp"       // 访问令牌的过期时间
	CtxKeyJwtIssuedAt     = "iat"       // 签发时间
	CtxKeyJwtIssuer       = "iss"       // 签发者
	CtxKeyJwtUserId       = "userId"    // 用户 ID
	CtxKeyJwtUserName     = "userName"  // 用户名
	CtxKeyJwtRoleId       = "roleId"    // 角色 ID
	CtxKeyJwtNickName     = "nickName"  // 昵称
	CtxKeyRefreshAt       = "refreshAt" // 刷新时间
	CtxKeyExpire          = "expire"    // 过期时长
	CtxKeyJwtToken        = "tokenStr"  // Token 字符串
)

// Jwt 结构体包含用于签名的密钥
type Jwt struct {
	SigningKey []byte // 用于签名的密钥
}

// SysJwtClaims 定义了系统 JWT 的声明内容，包含用户身份信息及标准 JWT 声明
type SysJwtClaims struct {
	UserId             string `json:"userId"`    // 用户 ID
	RoleId             string `json:"roleId"`    // 角色 ID
	DeptId             string `json:"deptId"`    // 部门 ID
	UserName           string `json:"userName"`  // 用户名
	NickName           string `json:"nickName"`  // 昵称
	RefreshAt          int64  `json:"refreshAt"` // 刷新时间
	Expire             int64  `json:"expire"`    // 过期时长（秒）
	TokenStr           string `json:"tokenStr"`  // 生成的 Token 字符串
	jwt.StandardClaims        // 标准 JWT 声明（包括 ExpiresAt、IssuedAt、Issuer 等字段）
}

// GetToken 使用 jwt.MapClaims 构造 JWT 声明，生成 token 字符串
// 参数：
//
//	secretKey - 签名密钥
//	username  - 用户名
//	nickName  - 昵称
//	issuer    - 签发者
//	iat       - 签发时间（Unix 时间戳）
//	seconds   - 有效时长（秒）
//	uid       - 用户 ID（int64 类型）
//	roleId    - 角色 ID（int64 类型）
//
// 返回 token 字符串及可能的错误
func GetToken(secretKey, username, nickName, issuer string, iat, seconds, uid, roleId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims[CtxKeyJwtAccessExpire] = iat + seconds
	claims[CtxKeyJwtIssuedAt] = iat
	claims[CtxKeyJwtUserId] = uid
	claims[CtxKeyJwtRoleId] = roleId
	claims[CtxKeyJwtNickName] = nickName
	claims[CtxKeyJwtUserName] = username
	claims[CtxKeyJwtIssuer] = issuer
	claims[CtxKeyRefreshAt] = seconds + iat/2
	claims[CtxKeyExpire] = seconds + iat
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// GetTokenClaims 使用 SysJwtClaims 结构体构造 JWT 声明，生成 token 字符串
// 参数：
//
//	secretKey - 签名密钥
//	username  - 用户名
//	nickName  - 昵称
//	issuer    - 签发者
//	iat       - 签发时间（Unix 时间戳）
//	seconds   - 有效时长（秒）
//	uid       - 用户 ID（字符串）
//	roleId    - 角色 ID（字符串）
//	deptId    - 部门 ID（字符串）
//
// 返回生成的 token 字符串及可能的错误
func GetTokenClaims(secretKey, username, nickName, issuer string, iat, seconds int64, uid, roleId, deptId string) (string, error) {
	var claims SysJwtClaims
	claims.ExpiresAt = iat + seconds
	claims.IssuedAt = iat
	claims.UserId = uid
	claims.RoleId = roleId
	claims.DeptId = deptId
	claims.NickName = nickName
	claims.UserName = username
	claims.Issuer = issuer
	claims.RefreshAt = iat + seconds/2
	claims.Expire = seconds
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	claims.TokenStr = tokenStr
	return tokenStr, err
}

// MakeJwt 根据传入的 SysJwtClaims 生成 JWT token 字符串
// 参数：
//
//	secretKey - 签名密钥
//	claim     - 指向 SysJwtClaims 结构体的指针，包含完整的声明信息
//
// 返回生成的 token 字符串及可能的错误
func MakeJwt(secretKey string, claim *SysJwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secretKey))
}
