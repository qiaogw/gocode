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

type CreateMigrationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMigrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMigrationLogic {
	return &CreateMigrationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMigrationLogic) CreateMigration(req *types.CreateMigrationRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.MigrationRpc.CreateMigration(l.ctx, &admin.CreateMigrationRequest{
		Version: req.Version,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "添加成功",
	}, nil
}
