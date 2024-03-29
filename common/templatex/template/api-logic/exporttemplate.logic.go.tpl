package {{.TableUrl}}

import (
	"context"
	"github.com/pkg/errors"
	"{{.ParentPkg}}/rpc/{{.Db}}"

	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

// ExportTemplate{{.Table}} 获取导入模板{{.TableComment}}
func (l *ExportTemplate{{.Table}}Logic) ExportTemplate{{.Table}}(req *types.NullRequest) (resp []byte, err error) {
	res, err := l.svcCtx.{{.Table}}Rpc.ExportTemplate{{.Table}}(l.ctx, &{{.Db}}.NullRequest{})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return res.Data, nil
}
