package {{.TableUrl}}logic

import (
"github.com/qiaogw/gocode/common/modelx"
	"github.com/qiaogw/gocode/common/errx"
	"context"
	"github.com/pkg/errors"

	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	{{ if .HasTimer }}"github.com/qiaogw/gocode/common/timex"{{ end }}
)

type Update{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Update{{.Table}}Logic {
	return &Update{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Update{{.Table}} 更新{{.TableComment}}
func (l *Update{{.Table}}Logic) Update{{.Table}}(in *{{.Db}}.Update{{.Table}}Request) (*{{.Db}}.Update{{.Table}}Response, error) {
	// 查询{{.TableComment}}是否存在
	res, err := l.svcCtx.{{.Table}}Model.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errors.Wrapf(errx.NewErrCode(errx.NoData), "该{{.TableComment}}不存在，id: %v", in.Id)
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData), "数据库查询 {{.TableComment}} 失败，id: %v,err:%v", in.Id,err)
	}
	{{- range  .Columns }}
		{{-  if .IsModelTime -}}
		{{- else if .IsControl -}}
        {{- else if .IsPage}}
		{{- else if .IsPk }}
		{{- else if eq .DataType "time.Time"}}
			res.{{.FieldName}}=timex.DatetimeStrToTime(in.{{.FieldName}})
		{{- else}}
			res.{{.FieldName}}=in.{{.FieldName}}
		{{- end}}
	{{- end }}
	res.UpdateBy = in.UpdateBy

	err = l.svcCtx.{{.Table}}Model.Update(l.ctx, res)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DbError),
"数据库更新 {{.TableComment}} 失败 , err:%v ,data : %+v  ", err, res)
	}
	var rep {{.Db}}.Update{{.Table}}Response
	return &rep, nil
}
