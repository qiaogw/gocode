package migration

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMigrationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMigrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMigrationLogic {
	return &UpdateMigrationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMigrationLogic) UpdateMigration(req *types.UpdateMigrationRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.MigrationRpc.UpdateMigration(l.ctx, &admin.UpdateMigrationRequest{
		Id:      req.Id,
		Version: req.Version,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "更新成功",
	}, nil
}
