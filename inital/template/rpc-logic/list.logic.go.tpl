package logic

import (
"github.com/qiaogw/gocode/common/modelx"
	"github.com/qiaogw/gocode/common/errorx"
	"context"
	"google.golang.org/grpc/status"
"github.com/jinzhu/copier"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"{{.ParentPkg}}/model"
{{ if .HasTimer }}"github.com/qiaogw/gocode/common/timex"{{ end }}
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
	qData.SortBY = in.SortBy
	qData.Descending = in.Descending
	list,count, err := l.svcCtx.{{.Table}}Model.FindAll(l.ctx, &qData)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errorx.NewCodeError(errorx.NoData, "{{.TableComment}}-该查询无数据")
		}
		return nil, status.Error(500, err.Error())
	}
	
	dataList := make([]*{{.Db}}.Get{{.Table}}Response, 0)
	_ = copier.Copy(&dataList, list)

	return &{{.Db}}.List{{.Table}}Response{
		List: dataList,
		Count: count,
	}, nil

}