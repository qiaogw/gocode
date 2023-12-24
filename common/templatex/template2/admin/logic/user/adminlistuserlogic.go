package user

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser(req *types.ListUserRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.UserRpc.ListUser(l.ctx, &admin.ListUserRequest{
		DeptId:     req.DeptId,
		PostId:     req.PostId,
		Uuid:       req.Uuid,
		Username:   req.Username,
		Password:   req.Password,
		NickName:   req.NickName,
		Mobile:     req.Mobile,
		Avatar:     req.Avatar,
		Gender:     req.Gender,
		Email:      req.Email,
		Sort:       req.Sort,
		Remark:     req.Remark,
		Status:     req.Status,
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
