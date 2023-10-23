package {{.TableUrl}}

import (
	"context"

"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.PackageName}}"

	"github.com/zeromicro/go-zero/core/logx"
)

type List{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewList{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *List{{.Table}}Logic {
	return &List{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *List{{.Table}}Logic) List{{.Table}}(req *types.List{{.Table}}Request) (resp *types.List{{.Table}}Response, err error) {
	res, err := l.svcCtx.{{.Service}}Rpc.List{{.Table}}(l.ctx, &{{.PackageName}}.List{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
			{{- else}}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
		SearchKey: req.SearchKey,
		SortBy:     req.SortBY,
		Descending: req.Descending,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var dataList []*types.Get{{.Table}}Response

	for _, v := range res.List {
		var dm types.Get{{.Table}}Response
		{{- range  .Columns }}
				{{- if eq .FieldName "DeletedAt"}}
				{{- else if .IsPage }}
				{{- else }}
					dm.{{.FieldName}}=v.{{.FieldName}}
				{{- end }}
		{{- end }}
dataList = append(dataList, &dm)
	}

	return &types.List{{.Table}}Response{
			Count: res.Count,
		List: dataList,
		}, nil
}
