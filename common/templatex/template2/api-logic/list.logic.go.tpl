package {{.TableUrl}}

import (
	"context"

	"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/model"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/timex"
	"github.com/qiaogw/gocode/common/modelx"
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

func (l *List{{.Table}}Logic) List{{.Table}}(req *types.List{{.Table}}Request) (resp *types.CommonResponse, err error) {

	// 查询{{.TableComment}}
	var qData model.List{{.Table}}Req
	{{- range  .Columns }}
		{{- if .IsPk }}
		{{- else}}
			{{- if eq .DataType "time.Time"}}
				qData.{{.FieldName}} =timex.DatetimeStrToTime(req.{{.FieldName}})
			{{- else}}
				qData.{{.FieldName}} = req.{{.FieldName}}
			{{- end}}
		{{- end}}
	{{- end }}
	qData.SearchKey = req.SearchKey
	qData.SortBY = req.SortBY
	qData.Descending = req.Descending
	list,count, err := l.svcCtx.{{.Table}}Model.FindAll(l.ctx, &qData)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
			"{{.TableComment}}-该查询无数据，查询条件: %+v", qData)
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
			"查询 {{.TableComment}} db fail，查询条件: %v,err:%v", req,err)
	}

	var dataList []*types.Get{{.Table}}Response

	for _, v := range list {
		var dm types.Get{{.Table}}Response
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
	var rep types.List{{.Table}}Response
	rep.List=dataList
	rep.Count=count
	return &types.CommonResponse{
			Code: errx.Success,
			Msg: "查询成功",
			Data: rep,
		}, nil
}
