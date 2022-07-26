package logic

import (
	"context"
	
	"{{.ParentPkg}}/common/errorx"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type Update{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Update{{.Table}}Logic {
	return &Update{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Update{{.Table}}Logic) Update{{.Table}}(req *types.Update{{.Table}}Request) (resp *types.Update{{.Table}}Response, err error) {
	l.Logger.Infof("l.svcCtx.{{.Service}}Rpc is %v\n", l.svcCtx.{{.Service}}Rpc)
	_, err = l.svcCtx.{{.Service}}Rpc.Update{{.Table}}(l.ctx, &{{.Db}}.Update{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
			{{- else}}
				{{- if .IsPage}}
				{{- else}}
				{{.FieldName}}: req.{{.FieldName}},
				{{- end}}
			{{- end}}
		{{- end }}
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	
	return &types.Update{{.Table}}Response{}, nil
}
