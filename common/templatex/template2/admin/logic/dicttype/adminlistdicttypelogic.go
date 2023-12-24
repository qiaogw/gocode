package dicttype

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ListDictTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDictTypeLogic {
	return &ListDictTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDictTypeLogic) ListDictType(req *types.ListDictTypeRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.DictTypeRpc.ListDictType(l.ctx, &admin.ListDictTypeRequest{
		Name:       req.Name,
		Type:       req.Type,
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
