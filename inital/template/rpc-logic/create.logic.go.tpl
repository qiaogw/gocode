package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"{{.ParentPkg}}/model"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type Create{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Create{{.Table}}Logic {
	return &Create{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create{{.Table}} 创建 {{.TableComment}}
func (l *Create{{.Table}}Logic) Create{{.Table}}(in *{{.Db}}.Create{{.Table}}Request) (*{{.Db}}.Create{{.Table}}Response, error) {
	new{{.Table}} := model.{{.Table}}{
		{{- range  .Columns }}
		{{- if .IsPk }}
		{{- else}}
			{{.FieldName}}: in.{{.FieldName}},
		{{- end}}
		{{- end }}
	}

	res, err := l.svcCtx.{{.Table}}Model.Insert(l.ctx, &new{{.Table}})
	if err != nil {
		logx.Infof("l.svcCtx.{{.Table}}Model.Insert err is %v\n", err)
		return nil, status.Error(500, err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}


	return &{{.Db}}.Create{{.Table}}Response{
	  Id: id,
	}, nil
}
