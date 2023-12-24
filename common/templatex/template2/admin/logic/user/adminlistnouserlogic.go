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

type ListNoUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListNoUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNoUserLogic {
	return &ListNoUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListNoUserLogic) ListNoUser(req *types.ListNoUserRequest) (resp *types.CommonResponse, err error) {
	res, err := l.svcCtx.UserRpc.ListNoUser(l.ctx, &admin.ListNoUserRequest{
		Id:         req.Id,
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
