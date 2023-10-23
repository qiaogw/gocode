package logic

import (
	"context"
	"{{.PKG}}/common/toolx"
	"{{.ParentPkg}}/model"
"github.com/pkg/errors"
"{{.PKG}}/common/errx"
	"{{.ParentPkg}}/rpc/{{.PackageName}}"
	"{{.ParentPkg}}/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportTemplate{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExportTemplate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportTemplate{{.Table}}Logic {
	return &ExportTemplate{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExportTemplate{{.Table}}Logic) ExportTemplate{{.Table}}(in *{{.PackageName}}.NullRequest) (*{{.PackageName}}.ExportResponse, error) {
	var m model.{{.Table}}
	resp, err := toolx.ExportToWebTemplate(&m, m.TableName())
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ServerCommonError),
			"导出{{.TableComment}}excel导入模板 fail，,err:%v", err)
	}
	return &{{.PackageName}}.ExportResponse{
		Data: resp.Bytes(),
	}, nil
}
