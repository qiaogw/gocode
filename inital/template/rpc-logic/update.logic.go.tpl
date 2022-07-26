package logic

import (
	"{{.ParentPkg}}/common/global"
	"context"
	"google.golang.org/grpc/status"

	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	{{ if .HasTimer }}"{{.ParentPkg}}/common/timex"{{ end }}
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
		if err == global.ErrNotFound {
			return nil, status.Error(100, "{{.TableComment}}不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	{{- range  .Columns }}
		{{- if eq .DataType "time.Time"}}
			res.{{.FieldName}}=timex.DatetimeStrToTime(in.{{.FieldName}})
		{{- else}}
			{{- if .IsPage}}
			{{- else}}
			res.{{.FieldName}}=in.{{.FieldName}}
			{{- end}}
		{{- end}}
	{{- end }}

	err = l.svcCtx.{{.Table}}Model.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &{{.Db}}.Update{{.Table}}Response{}, nil
}
