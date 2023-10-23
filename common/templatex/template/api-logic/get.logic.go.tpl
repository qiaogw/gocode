package {{.TableUrl}}

import (
	"context"

"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.PackageName}}"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Get{{.Table}}Logic {
	return &Get{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get{{.Table}}Logic) Get{{.Table}}(req *types.Get{{.Table}}Request) (resp *types.Get{{.Table}}Response, err error) {
	res, err := l.svcCtx.{{.Service}}Rpc.Get{{.Table}}(l.ctx, &{{.PackageName}}.Get{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

return &types.Get{{.Table}}Response{
	{{- range  .Columns }}
	{{- if eq .FieldName "DeletedAt"}}
	{{- else if .IsPage }}
	{{- else }}
			{{.FieldName}}: res.{{.FieldName}},
	{{- end }}
	{{- end }}
}, nil

}
