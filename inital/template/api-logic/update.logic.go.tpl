package {{.TableUrl}}

import (
	"context"
"{{.PKG}}/common/errx"
"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
"{{.PKG}}/common/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
)

type Update{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Update{{.Table}}Logic {
	return &Update{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Update{{.Table}}Logic) Update{{.Table}}(req *types.Update{{.Table}}Request) (resp *types.CommonResponse, err error) {
	l.Logger.Infof("l.svcCtx.{{.Service}}Rpc is %v\n", l.svcCtx.{{.Service}}Rpc)
	userId := jwtx.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.{{.Service}}Rpc.Update{{.Table}}(l.ctx, &{{.Db}}.Update{{.Table}}Request{
		{{- range  .Columns }}
				{{- if .IsPage}}
				{{- else if .IsModelTime -}}
				{{- else if .IsControl -}}
				{{- else}}
				{{.FieldName}}: req.{{.FieldName}},
				{{- end}}
		{{- end }}
		UpdateBy: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
	Code: errx.Success,
	Msg: "更新成功",
	}, nil
}
