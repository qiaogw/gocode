package user

import (
	"context"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportTemplateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportTemplateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportTemplateUserLogic {
	return &ExportTemplateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportTemplateUserLogic) ExportTemplateUser(req *types.NullRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
