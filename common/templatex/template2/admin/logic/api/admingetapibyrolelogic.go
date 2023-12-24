package api

import (
	"context"
	"github.com/pkg/errors"
	errorx2 "github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiByRoleLogic {
	return &GetApiByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiByRoleLogic) GetApiByRole(req *types.GetApiRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.AuthRpc.GetApiByRole(l.ctx, &admin.GetRoleRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errorx2.Success,
		Msg:  "查询成功",
		Data: res,
	}, nil
	return
}
