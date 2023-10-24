package {{.TableUrl}}

import (
	"context"

"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.PackageName}}"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
)

type Delete{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Delete{{.Table}}Logic {
	return &Delete{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete{{.Table}}Logic) Delete{{.Table}}(req *types.Delete{{.Table}}Request) (resp *types.CommonResponse, err error) {
	_, err = l.svcCtx.{{.Service}}Rpc.Delete{{.Table}}(l.ctx, &{{.PackageName}}.Delete{{.Table}}Request{
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
	Msg: "删除成功",
}, nil
}
