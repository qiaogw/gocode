package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ExportUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportUserLogic {
	return &ExportUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportUserLogic) ExportUser(req *types.ListUserRequest) (resp *types.ExportResponse, err error) {
	res, err := l.svcCtx.UserRpc.ExportUser(l.ctx, &admin.ExportRequest{
		PageIndex:  req.PageIndex,
		PageSize:   req.PageSize,
		SearchKey:  req.SearchKey,
		SortBy:     req.SortBY,
		Descending: req.Descending,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.ExportResponse{
		Byte: res.Data,
	}, nil
}
