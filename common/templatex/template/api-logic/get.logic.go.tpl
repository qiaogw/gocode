package {{.TableUrl}}

import (
	"context"

"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
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

	res, err := l.svcCtx.{{.Table}}Rpc.Get{{.Table}}(l.ctx, &{{.Db}}.Get{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

return &types.CommonResponse{
	Code: errx.Success,
	Msg: "查询成功",
	Data: res,
}, nil

}
