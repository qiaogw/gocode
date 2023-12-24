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

type GetOperaLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOperaLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperaLogLogic {
	return &GetOperaLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOperaLogLogic) GetOperaLog(req *types.GetOperaLogRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.OperaLogRpc.GetOperaLog(l.ctx, &admin.GetOperaLogRequest{
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
