package jwtx

import (
	"github.com/golang-jwt/jwt"
)

const (
	CtxKeyJwtAccessExpire = "exp"
	CtxKeyJwtIssuedAt     = "iat"
	CtxKeyJwtIssuer       = "iss"
	CtxKeyJwtUserId       = "userId"
	CtxKeyJwtUserName     = "userName"
	CtxKeyJwtRoleId       = "roleId"
	CtxKeyJwtDeptId       = "deptId"
	CtxKeyJwtNickName     = "nickName"
	CtxKeyRefreshAt       = "refreshAt"
	CtxKeyExpire          = "expire"
	CtxKeyJwtToken        = "tokenStr"
)

type Jwt struct {
	SigningKey []byte
}
type SysJwtClaims struct {
	UserId    string `json:"userId"`
	RoleId    string `json:"roleId"`
	DeptId    string `json:"deptId"`
	UserName  string `json:"userName"`
	NickName  string `json:"nickName"`
	RefreshAt int64  `json:"refreshAt"`
	Expire    int64  `json:"expire"`
	TokenStr  string `json:"tokenStr"`
	jwt.StandardClaims
}

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
func MakeJwt(secretKey string, claim *SysJwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secretKey))
}
