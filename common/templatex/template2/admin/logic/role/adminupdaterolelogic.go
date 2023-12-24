package role

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleRequest) (resp *types.CommonResponse, err error) {

	//userId := jwtx.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.RoleRpc.UpdateRole(l.ctx, &admin.UpdateRoleRequest{
		Id:           req.Id,
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
		Msg:  "更新成功",
	}, nil
}
