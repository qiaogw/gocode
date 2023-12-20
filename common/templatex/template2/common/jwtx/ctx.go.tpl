package jwtx

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

func GetUserIdFromCtx(ctx context.Context) string {
	var uid string
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(string); ok {
		uid = jsonUid
	}

	return uid
}

func GetRoleIdFromCtx(ctx context.Context) string {
	var id string
	if jsonUid, ok := ctx.Value(CtxKeyJwtRoleId).(string); ok {
		id = jsonUid
	}
	return id
}

func GetUserIdFromToken(tokenString, SigningKey string) (id string, err error) {
	claims, err := ParseToken(tokenString, SigningKey)
	if err != nil {
		return
	}
	id = claims.UserId
	return
}
func GetRoleIdFromToken(tokenString, SigningKey string) (id string, err error) {
	claims, err := ParseToken(tokenString, SigningKey)
	if err != nil {
		return
	}
	id = claims.RoleId
	return
}

func CheckToken(token, secretKey string) (ok bool, refreshToken string) {
	claims, err := ParseToken(token, secretKey)
	if err != nil {
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
