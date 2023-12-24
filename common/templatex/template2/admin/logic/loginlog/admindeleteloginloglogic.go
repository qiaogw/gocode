package loginlog

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type DeleteLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLoginLogLogic {
	return &DeleteLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLoginLogLogic) DeleteLoginLog(req *types.DeleteLoginLogRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.LoginLogRpc.DeleteLoginLog(l.ctx, &admin.DeleteLoginLogRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "删除成功",
	}, nil
}
