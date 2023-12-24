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

type DeleteMigrationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMigrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMigrationLogic {
	return &DeleteMigrationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMigrationLogic) DeleteMigration(req *types.DeleteMigrationRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.MigrationRpc.DeleteMigration(l.ctx, &admin.DeleteMigrationRequest{
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
