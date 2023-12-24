package api

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ListApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListApiLogic {
	return &ListApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListApiLogic) ListApi(req *types.ListApiRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.ApiRpc.ListApi(l.ctx, &admin.ListApiRequest{
		Title:      req.Title,
		Path:       req.Path,
		Method:     req.Method,
		Module:     req.Module,
		Remark:     req.Remark,
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
