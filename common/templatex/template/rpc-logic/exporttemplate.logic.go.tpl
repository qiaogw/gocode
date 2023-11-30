package {{.TableUrl}}logic

import (
	"context"
	"github.com/qiaogw/gocode/common/toolx"
	"{{.ParentPkg}}/model"
"github.com/pkg/errors"
"github.com/qiaogw/gocode/common/errx"
	"{{.ParentPkg}}/rpc/{{.Db}}"
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

func (l *ExportTemplate{{.Table}}Logic) ExportTemplate{{.Table}}(in *{{.Db}}.NullRequest) (*{{.Db}}.ExportResponse, error) {
	var m model.{{.Table}}
	resp, err := toolx.ExportToWebTemplate(&m, m.TableName())
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ServerCommonError),
			"导出{{.TableComment}}excel导入模板 失败，,err:%v", err)
	}
	return &{{.Db}}.ExportResponse{
		Data: resp.Bytes(),
	}, nil
}
