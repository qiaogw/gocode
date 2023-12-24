package user

import (
	"context"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportUserLogic {
	return &ImportUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportUserLogic) ImportUser(req *types.ImportRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
