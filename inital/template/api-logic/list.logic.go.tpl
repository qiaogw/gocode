package logic

import (
	"context"
	
	"{{.ParentPkg}}/common/errorx"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	
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

func (l *List{{.Table}}Logic) List{{.Table}}(req *types.List{{.Table}}Request) (resp []*types.List{{.Table}}Response, err error) {
	l.Logger.Infof("l.svcCtx.{{.Service}}Rpc is %v\n", l.svcCtx.{{.Service}}Rpc)
	res, err := l.svcCtx.{{.Service}}Rpc.List{{.Table}}(l.ctx, &{{.Db}}.List{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
			{{- else}}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	dataList := make([]*types.List{{.Table}}Response, 0)
	for _, item := range res.Data {
	dataList = append(dataList, &types.List{{.Table}}Response{
			{{- range  .Columns }}
				{{.FieldName}}: item.{{.FieldName}},
			{{- end }}
		})
	}

	return dataList, nil
}
