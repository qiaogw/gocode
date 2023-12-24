package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *types.CommonResponse, err error) {

	rpcReq := admin.UpdateUserRequest{
		Id:       req.Id,
		DeptId:   req.DeptId,
		PostId:   req.PostId,
		Uuid:     req.Uuid,
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
		Mobile:   req.Mobile,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
		Email:    req.Email,
		Sort:     req.Sort,
		Remark:   req.Remark,
		Status:   req.Status,
		RoleId:   req.RoleId,
	}
	_ = copier.Copy(&rpcReq.Roles, req.Roles)
	if req.RoleId == "" {
		rpcReq.RoleId = req.Roles[0].Id
	}
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &rpcReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "更新成功",
	}, nil
}
