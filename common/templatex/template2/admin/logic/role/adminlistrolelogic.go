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

type ListRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRoleLogic) ListRole(req *types.ListRoleRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.RoleRpc.ListRole(l.ctx, &admin.ListRoleRequest{
		Name:         req.Name,
		Code:         req.Code,
		Sort:         req.Sort,
		IsAdmin:      req.IsAdmin,
		DataScope:    req.DataScope,
		DefaultRoute: req.DefaultRoute,
		Remark:       req.Remark,
		PageIndex:    req.PageIndex,
		PageSize:     req.PageSize,
		SearchKey:    req.SearchKey,
		SortBy:       req.SortBY,
		Descending:   req.Descending,
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
