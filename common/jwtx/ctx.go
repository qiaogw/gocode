package jwtx

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
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
