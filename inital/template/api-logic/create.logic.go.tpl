package logic

import (
	"context"
	
	"{{.ParentPkg}}/common/errorx"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type Create{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Create{{.Table}}Logic {
	return &Create{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create{{.Table}}Logic) Create{{.Table}}(req *types.Create{{.Table}}Request) (resp *types.Create{{.Table}}Response, err error) {
	l.Logger.Infof("l.svcCtx.{{.Service}}Rpc is %v\n", l.svcCtx.{{.Service}}Rpc)
	res, err := l.svcCtx.{{.Service}}Rpc.Create{{.Table}}(l.ctx, &{{.Db}}.Create{{.Table}}Request{
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
	
	return &types.Create{{.Table}}Response{
		Id:     res.Id,
	}, nil
}
