package logic

import (
	"context"
	"github.com/pkg/errors"

	"{{.ParentPkg}}/model"
	"{{.ParentPkg}}/rpc/{{.PackageName}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
	{{ if .HasTimer }}"{{.PKG}}/common/timex"{{- end }}
	"{{.PKG}}/common/errx"
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
func (l *Create{{.Table}}Logic) Create{{.Table}}(in *{{.PackageName}}.Create{{.Table}}Request) (*{{.PackageName}}.Create{{.Table}}Response, error) {

{{$table:=.Table}}
{{$tableComment:=.TableComment}}
{{- range  .CacheKeys}}
	// 判断该{{.Field}}记录是否已经存在
	_, err := l.svcCtx.{{$table}}Model.FindOneBy{{.Field}}(l.ctx,in.{{.Field}})
	if err == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.Duplicate), "该{{$tableComment}}已存在")
	}
{{- end}}

new{{.Table}} := model.{{.Table}}{
		{{- range  .Columns }}
		{{- if .IsPk }}
		{{- else if .IsModelTime -}}
		{{- else if .IsControl -}}
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
	new{{.Table}}.CreateBy = in.CreateBy
	res, err := l.svcCtx.{{.Table}}Model.Insert(l.ctx, &new{{.Table}})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DbError),
		"create {{$tableComment}} db insert fail , err:%v ,data : %+v  ", err, new{{.Table}})
	}
	var rep {{.PackageName}}.Create{{.Table}}Response
	_ = copier.Copy(&rep, res)
	return &rep, nil
}
