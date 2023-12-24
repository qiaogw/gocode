package dept

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptLogic {
	return &UpdateDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDeptLogic) UpdateDept(req *types.UpdateDeptRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.DeptRpc.UpdateDept(l.ctx, &admin.UpdateDeptRequest{
		Id:       req.Id,
		ParentId: req.ParentId,
		Name:     req.Name,
		Sort:     req.Sort,
		Leader:   req.Leader,
		Phone:    req.Phone,
		Email:    req.Email,
		Enabled:  req.Enabled,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "更新成功",
	}, nil
}
