package {{.TableUrl}}

import (
	"context"

	"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"

	"github.com/qiaogw/gocode/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/qiaogw/gocode/common/toolx"
	"{{.ParentPkg}}/model"
)

type ExportTemplate{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportTemplate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportTemplate{{.Table}}Logic {
	return &ExportTemplate{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportTemplate{{.Table}}Logic) ExportTemplate{{.Table}}(req *types.NullRequest) (resp []byte, err error) {
	var m model.{{.Table}}
	res, err := toolx.ExportToWebTemplate(&m, m.TableName())
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ServerCommonError),
			"导出{{.TableComment}}excel导入模板 失败，,err:%v", err)
	}

	return res.Bytes(), nil
}
