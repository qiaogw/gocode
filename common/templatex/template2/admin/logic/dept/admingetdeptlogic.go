package dept

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type GetDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeptLogic {
	return &GetDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeptLogic) GetDept(req *types.GetDeptRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.DeptRpc.GetDept(l.ctx, &admin.GetDeptRequest{
		Id: req.Id,
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
