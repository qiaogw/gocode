package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ListMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMenuLogic {
	return &ListMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMenuLogic) ListMenu(req *types.ListMenuRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.MenuRpc.ListMenu(l.ctx, &admin.ListMenuRequest{
		Name:       req.Name,
		Title:      req.Title,
		Icon:       req.Icon,
		Path:       req.Path,
		Type:       req.Type,
		Component:  req.Component,
		ParentId:   req.ParentId,
		Sort:       req.Sort,
		KeepAlive:  req.KeepAlive,
		Hidden:     req.Hidden,
		IsFrame:    req.IsFrame,
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
	var data types.ListMenuResponse
	_ = copier.Copy(&data, res)
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "查询成功",
		Data: data,
	}, nil
}
