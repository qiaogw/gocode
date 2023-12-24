package role

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.RoleRpc.CreateRole(l.ctx, &admin.CreateRoleRequest{
		Name:         req.Name,
		Code:         req.Code,
		Sort:         req.Sort,
		IsAdmin:      req.IsAdmin,
		DataScope:    req.DataScope,
		DefaultRoute: req.DefaultRoute,
		Remark:       req.Remark,
		DataFilter:   req.DataFilter,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "添加成功",
	}, nil
}
