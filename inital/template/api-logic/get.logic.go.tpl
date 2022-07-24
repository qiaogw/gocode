package logic

import (
	"context"
	
	"{{.ParentPkg}}/common/errorx"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	
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
	l.Logger.Infof("l.svcCtx.{{.Service}}Rpc is %v\n", l.svcCtx.{{.Service}}Rpc)
	res, err := l.svcCtx.{{.Service}}Rpc.Get{{.Table}}(l.ctx, &{{.Db}}.Get{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	
	return &types.Get{{.Table}}Response{
		{{- range  .Columns }}
			{{.FieldName}}: res.{{.FieldName}},
		{{- end -}}
	}, nil
}
