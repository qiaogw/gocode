package operalog

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type DeleteOperaLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOperaLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOperaLogLogic {
	return &DeleteOperaLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOperaLogLogic) DeleteOperaLog(req *types.DeleteOperaLogRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.OperaLogRpc.DeleteOperaLog(l.ctx, &admin.DeleteOperaLogRequest{
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
