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

type SetRoleUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoleUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleUsersLogic {
	return &SetRoleUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoleUsersLogic) SetRoleUsers(req *types.UpdateRoleUsersRequest) (resp *types.CommonResponse, err error) {
	var rpcReq admin.UpdateRoleUsersRequest
	rpcReq.Id = req.Id
	for _, v := range req.Ids {
		rpcReq.Ids = append(rpcReq.Ids, &admin.GetUserRequest{
			Id: v.Id,
		})
	}
	_, err = l.svcCtx.RoleRpc.SetRoleUsers(l.ctx, &rpcReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "更新成功",
	}, nil
}
