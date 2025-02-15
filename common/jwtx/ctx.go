package jwtx

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

// GetUserIdFromCtx 从上下文中获取用户ID
// 从 ctx 中取出存储的用户ID，键为 CtxKeyJwtUserId
func GetUserIdFromCtx(ctx context.Context) string {
	var uid string
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(string); ok {
		uid = jsonUid
	}
	return uid
}

// GetRoleIdFromCtx 从上下文中获取角色ID
// 从 ctx 中取出存储的角色ID，键为 CtxKeyJwtRoleId
func GetRoleIdFromCtx(ctx context.Context) string {
	var id string
	if jsonUid, ok := ctx.Value(CtxKeyJwtRoleId).(string); ok {
		id = jsonUid
	}
	return id
}

// GetTokenStrFromCtx 从上下文中获取Token字符串
// 从 ctx 中取出存储的Token字符串，键为 CtxKeyJwtToken
func GetTokenStrFromCtx(ctx context.Context) string {
	var id string
	if jsonUid, ok := ctx.Value(CtxKeyJwtToken).(string); ok {
		id = jsonUid
	}
	return id
}

// GetUserIdFromToken 从JWT token中解析出用户ID
// 参数 tokenString 为JWT字符串，SigningKey 为签名密钥
func GetUserIdFromToken(tokenString, SigningKey string) (id string, err error) {
	claims, err := ParseToken(tokenString, SigningKey)
	if err != nil {
		return
	}
	id = claims.UserId
	return
}

// GetRoleIdFromToken 从JWT token中解析出角色ID
// 参数 tokenString 为JWT字符串，SigningKey 为签名密钥
func GetRoleIdFromToken(tokenString, SigningKey string) (id string, err error) {
	claims, err := ParseToken(tokenString, SigningKey)
	if err != nil {
		return
	}
	id = claims.RoleId
	return
}

// GetDeptIdFromToken 从JWT token中解析出部门ID
// 参数 tokenString 为JWT字符串，SigningKey 为签名密钥
func GetDeptIdFromToken(tokenString, SigningKey string) (id string, err error) {
	claims, err := ParseToken(tokenString, SigningKey)
	if err != nil {
		return
	}
	id = claims.DeptId
	return
}

// GetClaimsFromToken 从JWT token中解析出所有的声明信息（Claims）
// 返回解析后的 SysJwtClaims 结构体
func GetClaimsFromToken(tokenString, SigningKey string) (data *SysJwtClaims, err error) {
	claims, err := ParseToken(tokenString, SigningKey)
	if err != nil {
		return
	}
	data = claims
	return
}

// CheckToken 校验JWT token的有效性，并在需要时刷新token
// 参数 token 为原始token，secretKey 为签名密钥
// 如果 token 已经过期但刷新时间未过，则生成新的token，返回 true 和新的 refreshToken
// 如果 token 有效，则返回 true 和空字符串；否则返回 false
func CheckToken(token, secretKey string) (ok bool, refreshToken string) {
	claims, err := ParseToken(token, secretKey)
	if err != nil {
		// 此处可根据需要添加过期后检查Refresh的逻辑，目前直接返回空结果
		return
	}
	now := time.Now().Unix()
	// 如果刷新时间已过，则生成新的token
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

// ParseToken 解析JWT token并返回 SysJwtClaims
// 参数 tokenString 为JWT字符串，SigningKey 为签名密钥
// 如果token有效则返回解析后的 SysJwtClaims，否则返回错误信息
func ParseToken(tokenString, SigningKey string) (*SysJwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &SysJwtClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte(SigningKey), nil
		})
	// 如果 token 不为空，则尝试解析 Claims
	if token != nil {
		if claims, ok := token.Claims.(*SysJwtClaims); ok {
			// 如果存在错误且 token 无效，检查是否因过期导致
			if err != nil && !token.Valid {
				if vError, vOk := err.(*jwt.ValidationError); vOk {
					if vError.Errors&jwt.ValidationErrorExpired != 0 {
						// 如果因过期导致，则返回 claims 并提示“checkRefresh”
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
