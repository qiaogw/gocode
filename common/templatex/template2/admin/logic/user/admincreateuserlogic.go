package user

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

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) (resp *types.CommonResponse, err error) {

	rpcReq := admin.CreateUserRequest{
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
		if len(req.Roles) > 0 {
			rpcReq.RoleId = req.Roles[0].Id
		}
	}
	_, err = l.svcCtx.UserRpc.CreateUser(l.ctx, &rpcReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "添加成功",
	}, nil
}
