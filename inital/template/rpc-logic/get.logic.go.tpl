package logic

import (
	"{{.ParentPkg}}/common/global"
	"context"
	"google.golang.org/grpc/status"

	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"

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
		if err == global.ErrNotFound {
			return nil, status.Error(100, "{{.TableComment}}不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &{{.Db}}.Get{{.Table}}Response{
      {{- range  .Columns }}
		  {{- if eq .DataType "time.Time"}}
			  {{.FieldName}}: res.{{.FieldName}}.String(),
		  {{- else}}
			  {{- if .IsPage}}
			  {{- else}}
			  {{.FieldName}}: res.{{.FieldName}},
			  {{- end}}
		  {{- end}}
      {{- end -}}
	}, nil
}
