package logic

import (
	"context"
	
	"{{.ParentPkg}}/common/errorx"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type Delete{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Delete{{.Table}}Logic {
	return &Delete{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete{{.Table}}Logic) Delete{{.Table}}(req *types.Delete{{.Table}}Request) (resp *types.Delete{{.Table}}Response, err error) {
	l.Logger.Infof("l.svcCtx.{{.Service}}Rpc is %v\n", l.svcCtx.{{.Service}}Rpc)
	_, err = l.svcCtx.{{.Service}}Rpc.Delete{{.Table}}(l.ctx, &{{.Db}}.Delete{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	
	return &types.Delete{{.Table}}Response{}, nil
}
