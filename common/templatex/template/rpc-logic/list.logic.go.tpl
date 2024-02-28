package {{.TableUrl}}logic

import (
"github.com/qiaogw/gocode/common/modelx"
	"github.com/qiaogw/gocode/common/errx"
	"context"
"github.com/pkg/errors"

"github.com/jinzhu/copier"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"{{.ParentPkg}}/model"
"github.com/qiaogw/gocode/common/gormx"
"github.com/qiaogw/gocode/common/timex"
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
	qData.SearchKey = in.SearchKey
	qData.SortBY = gormx.GetSortBy(qData.{{.Table}}Order, in.SortBy)
	qData.Descending = in.Descending
	list,count, err := l.svcCtx.{{.Table}}Model.FindAll(l.ctx, &qData)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
"{{.TableComment}}-该查询无数据，查询条件: %+v", qData)
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
"查询 {{.TableComment}} 失败:%v",err)

}
	
	var dataList[]*{{.Db}}.Get{{.Table}}Response

	for _, v := range list {
		var dm {{.Db}}.Get{{.Table}}Response
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


	return &{{.Db}}.List{{.Table}}Response{
		List: dataList,
		Count: count,
	}, nil

}