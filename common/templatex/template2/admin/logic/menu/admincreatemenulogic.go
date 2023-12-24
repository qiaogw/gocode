package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type CreateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuLogic) CreateMenu(req *types.CreateMenuRequest) (resp *types.CommonResponse, err error) {

	userId := jwtx.GetUserIdFromCtx(l.ctx)
	var rpcReq admin.CreateMenuRequest
	_ = copier.Copy(&rpcReq, req)
	rpcReq.UpdateBy = userId
	_, err = l.svcCtx.MenuRpc.CreateMenu(l.ctx, &rpcReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "添加成功",
	}, nil

}
