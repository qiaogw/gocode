package config

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type DeleteConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigLogic {
	return &DeleteConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteConfigLogic) DeleteConfig(req *types.DeleteConfigRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.ConfigRpc.DeleteConfig(l.ctx, &admin.DeleteConfigRequest{
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
