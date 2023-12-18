package {{.TableUrl}}

import (
	"context"

	"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/jinzhu/copier"
"github.com/qiaogw/gocode/common/timex"
)

type Get{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Get{{.Table}}Logic {
	return &Get{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get{{.Table}}Logic) Get{{.Table}}(req *types.Get{{.Table}}Request) (resp *types.CommonResponse, err error) {

	// 查询{{.TableComment}} 是否存在
	res, err := l.svcCtx.{{.Table}}Model.FindOne(l.ctx, req.Id)
	if err != nil {
		if err == modelx.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
			"该{{.TableComment}}不存在，id is %v", req.Id)
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
			"数据库提取 {{.TableComment}} 失败l，id: %v,err:%v", req.Id,err)
		}
	var rep types.Get{{.Table}}Response
	_ = copier.Copy(&rep, res)
	{{- range  .Columns }}
		{{- if eq .DataType "time.Time"}}
			{{- if eq .FieldName "DeletedAt"}}
				{{- else }}
			if !res.{{.FieldName}}.IsZero() {
			rep.{{.FieldName}}=timex.TimeToDatetimeStr(res.{{.FieldName}})
			}
				{{- end }}
		{{- end}}
	{{- end }}

return &types.CommonResponse{
	Code: errx.Success,
	Msg: "查询成功",
	Data: rep,
}, nil

}
