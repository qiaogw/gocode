package jwtx

import (
	"github.com/golang-jwt/jwt"
)

const (
	CtxKeyJwtAccessExpire = "exp"
	CtxKeyJwtIssuedAt     = "iat"
	CtxKeyJwtIssuer       = "iss"
	CtxKeyJwtUserId       = "userid"
	CtxKeyJwtUserName     = "username"
	CtxKeyJwtRoleId       = "roleId"
	CtxKeyJwtNickName     = "nickName"
)

func GetToken(secretKey, username, nickName, issuer string, iat, seconds, uid, roleId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims[CtxKeyJwtAccessExpire] = iat + seconds
	claims[CtxKeyJwtIssuedAt] = iat
	claims[CtxKeyJwtUserId] = uid
	claims[CtxKeyJwtRoleId] = roleId
	claims[CtxKeyJwtNickName] = nickName
	claims[CtxKeyJwtUserName] = username
	claims[CtxKeyJwtIssuer] = issuer
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
