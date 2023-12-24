package dictdata

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ListDictDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDictDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDictDataLogic {
	return &ListDictDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDictDataLogic) ListDictData(req *types.ListDictDataRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.DictDataRpc.ListDictData(l.ctx, &admin.ListDictDataRequest{
		DictTypeId: req.DictTypeId,
		Sort:       req.Sort,
		Label:      req.Label,
		Value:      req.Value,
		IsDefault:  req.IsDefault,
		Enabled:    req.Enabled,
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
