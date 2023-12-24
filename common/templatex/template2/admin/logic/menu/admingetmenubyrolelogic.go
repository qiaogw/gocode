package menu

import (
	"context"
	"github.com/pkg/errors"
	errorx2 "github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/rpc/admin"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuByRoleLogic {
	return &GetMenuByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuByRoleLogic) GetMenuByRole(req *types.GetMenuRequest) (resp *types.CommonResponse, err error) {

	//uid := l.ctx.Value("uid")

	res, err := l.svcCtx.MenuRpc.GetMenuByRole(l.ctx, &admin.GetRoleRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errorx2.Success,
		Msg:  "查询成功",
		Data: res,
	}, nil
}
