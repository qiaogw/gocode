package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"
	"sub-admin/admin/rpc/admin"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPasswordLogic {
	return &SetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetPasswordLogic) SetPassword(req *types.SetPasswordRequest) (resp *types.CommonResponse, err error) {
	uid := jwtx.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.UserRpc.SetPassword(l.ctx, &admin.SetPasswordRequest{
		Id:          uid,
		OldPassword: req.OldPassword,
		Password:    req.Password,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "密码修改成功",
	}, nil
}
