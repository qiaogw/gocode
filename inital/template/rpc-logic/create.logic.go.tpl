package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"{{.ParentPkg}}/model"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	{{ if .HasTimer }}"{{.ParentPkg}}/common/timex"{{ end }}
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

{{$table:=.Table}}
{{$tableComment:=.TableComment}}
{{- range  .CacheKeys}}
	// 判断该{{.Field}}记录是否已经存在
	_, err := l.svcCtx.{{$table}}Model.FindOneBy{{.Field}}(l.ctx,in.{{.Field}})
	if err == nil {
	return nil, status.Error(100, "该{{$tableComment}}已存在")
	}
{{- end}}

new{{.Table}} := model.{{.Table}}{
		{{- range  .Columns }}
		{{- if .IsPk }}
		{{- else}}
			{{- if eq .DataType "time.Time"}}
				{{.FieldName}}: timex.DatetimeStrToTime(in.{{.FieldName}}),
			{{- else}}
				{{- if .IsPage}}
				{{- else}}
				{{.FieldName}}: in.{{.FieldName}},
				{{- end}}
			{{- end}}
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
		if err.Error() == "LastInsertId is not supported by this driver"{
			logx.Infof("res is %v,err is %v\n", res,err)
		} else {
			return nil, status.Error(500, err.Error())
		}
	}


	return &{{.Db}}.Create{{.Table}}Response{
	  Id: id,
	}, nil
}
