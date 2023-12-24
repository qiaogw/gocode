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

type SetMeRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMeRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMeRoleLogic {
	return &SetMeRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMeRoleLogic) SetMeRole(req *types.SetMeRoleRequest) (resp *types.CommonResponse, err error) {
	uid := jwtx.GetUserIdFromCtx(l.ctx)
	res, err := l.svcCtx.UserRpc.SetMeRole(l.ctx, &admin.SetMeRoleRequest{
		Id:     uid,
		RoleId: req.RoleId,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "切换成功",
		Data: res,
	}, nil
}
