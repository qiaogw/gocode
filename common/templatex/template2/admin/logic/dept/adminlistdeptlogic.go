package dept

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ListDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeptLogic {
	return &ListDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeptLogic) ListDept(req *types.ListDeptRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.DeptRpc.ListDept(l.ctx, &admin.ListDeptRequest{
		ParentId:   req.ParentId,
		Name:       req.Name,
		Sort:       req.Sort,
		Leader:     req.Leader,
		Phone:      req.Phone,
		Email:      req.Email,
		Enabled:    req.Enabled,
		PageIndex:  req.PageIndex,
		PageSize:   req.PageSize,
		SearchKey:  req.SearchKey,
		SortBy:     req.SortBY,
		Descending: req.Descending,
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
