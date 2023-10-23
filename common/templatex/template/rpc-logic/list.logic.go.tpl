package logic

import (
"{{.PKG}}/common/modelx"
	"{{.PKG}}/common/errx"
"{{.PKG}}/common/timex"
	"context"
"github.com/pkg/errors"

"github.com/jinzhu/copier"
	"{{.ParentPkg}}/rpc/{{.PackageName}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"{{.ParentPkg}}/model"
{{ if .HasTimer }}"{{.PKG}}/common/timex"{{ end }}
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
func (l *List{{.Table}}Logic) List{{.Table}}(in *{{.PackageName}}.List{{.Table}}Request) (*{{.PackageName}}.List{{.Table}}Response, error) {
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
	qData.SearchKey = in.SearchKey
	qData.SortBY = in.SortBy
	qData.Descending = in.Descending
	list,count, err := l.svcCtx.{{.Table}}Model.FindAll(l.ctx, &qData)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
"{{.TableComment}}-该查询无数据，查询条件: %+v", qData)
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
"查询 {{.TableComment}} 数据库查询失败，查询条件: %v,err:%v", in,err)

}
	
	var dataList []*{{.PackageName}}.Get{{.Table}}Response

	for _, v := range list {
		var dm {{.PackageName}}.Get{{.Table}}Response
		_ = copier.Copy(&dm, v)
		{{- range  .Columns }}
			{{- if eq .DataType "time.Time"}}
				{{- if eq .FieldName "DeletedAt"}}
				{{- else }}
					if !v.{{.FieldName}}.IsZero() {
					dm.{{.FieldName}}=timex.TimeToDatetimeStr(v.{{.FieldName}})
					}
				{{- end }}
			{{- end}}
		{{- end }}
dataList = append(dataList, &dm)
	}


	return &{{.PackageName}}.List{{.Table}}Response{
		List: dataList,
		Count: count,
	}, nil

}