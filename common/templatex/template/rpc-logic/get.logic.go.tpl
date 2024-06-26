package {{.TableUrl}}logic

import (
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/qiaogw/gocode/common/errx"
	"context"
"github.com/pkg/errors"

"github.com/jinzhu/copier"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
"github.com/qiaogw/gocode/common/timex"
)

type Get{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGet{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Get{{.Table}}Logic {
	return &Get{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get{{.Table}} 提取单条 {{.TableComment}}
func (l *Get{{.Table}}Logic) Get{{.Table}}(in *{{.Db}}.Get{{.Table}}Request) (*{{.Db}}.Get{{.Table}}Response, error) {
	// 查询{{.TableComment}} 是否存在
	res, err := l.svcCtx.{{.Table}}Model.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == modelx.ErrNotFound {
return nil, errx.NewErrCodeMsg(errx.ErrReq, "该{{.TableComment}}不存在")
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
"数据库提取 {{.TableComment}} 失败:%v",err)
	}
	var rep {{.Db}}.Get{{.Table}}Response
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
	return &rep, nil
}
