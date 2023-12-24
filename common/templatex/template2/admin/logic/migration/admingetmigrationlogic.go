package migration

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type GetMigrationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMigrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMigrationLogic {
	return &GetMigrationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMigrationLogic) GetMigration(req *types.GetMigrationRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.MigrationRpc.GetMigration(l.ctx, &admin.GetMigrationRequest{
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
