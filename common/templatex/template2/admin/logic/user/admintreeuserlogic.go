package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/rpc/admin"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeUserLogic {
	return &TreeUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeUserLogic) TreeUser(req *types.ListUserRequest) (resp *types.CommonResponse, err error) {
	res, err := l.svcCtx.UserRpc.TreeUser(l.ctx, &admin.ListUserRequest{
		DeptId: req.DeptId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "查询成功",
		Data: res,
	}, nil
}
