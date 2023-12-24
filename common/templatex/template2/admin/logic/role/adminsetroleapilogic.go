package role

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/rpc/admin"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoleApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoleApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleApiLogic {
	return &SetRoleApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoleApiLogic) SetRoleApi(req *types.SetRoleRequest) (resp *types.CommonResponse, err error) {

	var rpcReq admin.SetRoleRequest
	rpcReq.RoleId = req.Id

	for _, v := range req.Permission {
		rpcReq.Permission = append(rpcReq.Permission, &admin.DeleteRoleRequest{
			Id: v,
		})
	}
	_, err = l.svcCtx.RoleRpc.SetRoleApi(l.ctx, &rpcReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "授权成功",
	}, nil
}
