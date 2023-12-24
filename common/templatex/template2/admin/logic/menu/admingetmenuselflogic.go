package menu

import (
	"context"
	"github.com/pkg/errors"
	errorx2 "github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"
	"sub-admin/admin/rpc/admin"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuSelfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuSelfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuSelfLogic {
	return &GetMenuSelfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuSelfLogic) GetMenuSelf(req *types.GetMenuSelfRequest) (resp *types.CommonResponse, err error) {

	//uid := l.ctx.Value("uid")
	rid := jwtx.GetRoleIdFromCtx(l.ctx)

	res, err := l.svcCtx.MenuRpc.GetMenuByRole(l.ctx, &admin.GetRoleRequest{
		Id: rid,
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
