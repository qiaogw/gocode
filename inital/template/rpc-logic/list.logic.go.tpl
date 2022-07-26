package logic

import (
	"{{.ParentPkg}}/common/global"
	"context"
	"google.golang.org/grpc/status"

	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"{{.ParentPkg}}/model"
{{ if .HasTimer }}"{{.ParentPkg}}/common/timex"{{ end }}
	"github.com/zeromicro/go-zero/core/logx"

)

type List{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewList{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *List{{.Table}}Logic {
	return &List{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// List{{.Table}} 条件查询 {{.TableComment}} 列表
func (l *List{{.Table}}Logic) List{{.Table}}(in *{{.Db}}.List{{.Table}}Request) (*{{.Db}}.List{{.Table}}Response, error) {
	// 查询{{.TableComment}}
	var qData model.List{{.Table}}Req
	{{- range  .Columns }}
		{{- if .IsPk }}
		{{- else}}
			{{- if eq .DataType "time.Time"}}
				qData.{{.FieldName}} =timex.DatetimeStrToTime(in.{{.FieldName}})
			{{- else}}
				qData.{{.FieldName}} = in.{{.FieldName}}
			{{- end}}
		{{- end}}
	{{- end }}
	list, err := l.svcCtx.{{.Table}}Model.FindAll(l.ctx, &qData)
	if err != nil {
		if err == global.ErrNotFound {
			return nil, status.Error(100, "{{.TableComment}} 不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	
	dataList := make([]*{{.Db}}.Get{{.Table}}Response, 0)
	for _, item := range list {
		dataList = append(dataList, &{{.Db}}.Get{{.Table}}Response{
			{{- range  .Columns }}
				{{- if eq .DataType "time.Time"}}
					{{.FieldName}}: item.{{.FieldName}}.String(),
				{{- else}}
					{{- if .IsPage}}
					{{- else}}
					{{.FieldName}}: item.{{.FieldName}},
					{{- end}}
					{{- end}}
			{{- end }}
		})
	}

	return &{{.Db}}.List{{.Table}}Response{
		Data: dataList,
	}, nil

}