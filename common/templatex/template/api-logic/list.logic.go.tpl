package {{.TableUrl}}

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"
	"github.com/zeromicro/go-zero/core/logx"

	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
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
// List{{.Table}} 查询{{.TableComment}}
func (l *List{{.Table}}Logic) List{{.Table}}(req *types.List{{.Table}}Request) (resp *types.CommonResponse, err error) {
  token := jwtx.GetTokenStrFromCtx(l.ctx)
	res, err := l.svcCtx.{{.Table}}Rpc.List{{.Table}}(l.ctx, &{{.Db}}.List{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
			{{- else}}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
		Token:      token,
		SearchKey: req.SearchKey,
		SortBy:     req.SortBY,
		Descending: req.Descending,
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
