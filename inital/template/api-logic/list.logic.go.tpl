package logic

import (
	"context"

"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.PKG}}/common/errx"
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
	l.Logger.Infof("l.svcCtx.{{.Service}}Rpc is %v\n", l.svcCtx.{{.Service}}Rpc)
	res, err := l.svcCtx.{{.Service}}Rpc.List{{.Table}}(l.ctx, &{{.Db}}.List{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
			{{- else}}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
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
