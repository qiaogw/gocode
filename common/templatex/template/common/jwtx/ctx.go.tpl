package jwtx

import (
"context"
"encoding/json"
"errors"
"fmt"
"github.com/golang-jwt/jwt"
"github.com/zeromicro/go-zero/core/logx"
"time"
)

func GetUserIdFromCtx(ctx context.Context) int64 {
var uid int64
if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
if int64Uid, err := jsonUid.Int64(); err == nil {
uid = int64Uid
} else {
logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
}
}
return uid
}

func GetRoleIdFromCtx(ctx context.Context) int64 {
var id int64
if jsonUid, ok := ctx.Value(CtxKeyJwtRoleId).(json.Number); ok {
if int64Uid, err := jsonUid.Int64(); err == nil {
id = int64Uid
} else {
logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
}
}
return id
}

func CheckToken(token, secretKey string) (ok bool, refreshToken string) {
claims, err := ParseToken(token, secretKey)
if err != nil {
////ExpiresAt已经过期，检查Refresh是否过期
//if err.Error() == "checkRefresh" {
//	//如果ExpiresAt在当前时间之前（已过期），但RefreshAt在当前时间之后（还未过期）
//	if claims.ExpiresAt < time.Now().Unix() && claims.RefreshAt > time.Now().Unix() {
//		//重新生成新的token，并插入Header里
//		jt := jwt.New(jwt.SigningMethodHS256)
//		jt.Claims = claims
//		refreshToken, _ = jt.SignedString([]byte(secretKey))
//		if refreshToken != "" {
//			return true, refreshToken
//		}
//	}
//	//双重超期
//	if claims.ExpiresAt < time.Now().Unix() && claims.RefreshAt < time.Now().Unix() {
//		return
//	}
//}
return
}
now := time.Now().Unix()
if claims.RefreshAt < now {
jt := jwt.New(jwt.SigningMethodHS256)
claims.ExpiresAt = now + claims.Expire
claims.RefreshAt = now + claims.Expire/2
claims.IssuedAt = now
jt.Claims = claims
refreshToken, _ = jt.SignedString([]byte(secretKey))
}
return true, refreshToken
}

func ParseToken(tokenString, SigningKey string) (*SysJwtClaims, error) {
token, err := jwt.ParseWithClaims(tokenString, &SysJwtClaims{},
func(token *jwt.Token) (i interface{}, e error) {
return []byte(SigningKey), nil
})

//拿到token
if token != nil {
if claims, ok := token.Claims.(*SysJwtClaims); ok {
//如果上面的err不为空，判断是否Expired超期，也返回claims，其他情况返回空
if err != nil && !token.Valid {
if vError, vOk := err.(*jwt.ValidationError); vOk {
if vError.Errors&jwt.ValidationErrorExpired != 0 {
return claims, errors.New("checkRefresh")
}
} else {
return nil, errors.New("身份鉴别失败！")
}
} else if token.Valid {
return claims, nil
}
}
return nil, errors.New("身份鉴别失败！")
} else {
return nil, errors.New("身份鉴别失败！")
}
}
func MakeJwt(secret string) (*Jwt, error) {
if secret == "" {
return nil, errors.New("没有找到Jwt密钥配置，请重新初始化数据库！")
}
return &Jwt{
[]byte(secret),
}, nil
}
