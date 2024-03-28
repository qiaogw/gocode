package {{.TableUrl}}

import (
	"context"
	"github.com/pkg/errors"
	"{{.ParentPkg}}/rpc/{{.Db}}"

	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Export{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExport{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Export{{.Table}}Logic {
	return &Export{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
// Export{{.Table}} 导出{{.TableComment}}
func (l *Export{{.Table}}Logic) Export{{.Table}}(req *types.List{{.Table}}Request) (resp []byte, err error) {
	res, err := l.svcCtx.{{.Table}}Rpc.Export{{.Table}}(l.ctx, &{{.Db}}.ExportRequest{
		PageIndex:  req.PageIndex,
		PageSize:   req.PageSize,
		SearchKey:  req.SearchKey,
		SortBy:     req.SortBY,
		Descending: req.Descending,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return res.Data, nil
}
