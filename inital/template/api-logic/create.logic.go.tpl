package {{.TableUrl}}

import (
	"context"

"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.PKG}}/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
)

type Create{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Create{{.Table}}Logic {
	return &Create{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create{{.Table}}Logic) Create{{.Table}}(req *types.Create{{.Table}}Request) (resp *types.CommonResponse, err error) {
	l.Logger.Infof("l.svcCtx.{{.Service}}Rpc is %v\n", l.svcCtx.{{.Service}}Rpc)
	userId := jwtx.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.{{.Service}}Rpc.Create{{.Table}}(l.ctx, &{{.Db}}.Create{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
			{{- else if .IsModelTime -}}
			{{- else if .IsControl -}}
			{{- else}}
				{{- if .IsPage}}
				{{- else}}
				{{.FieldName}}: req.{{.FieldName}},
				{{- end}}
			{{- end}}
		{{- end }}
		CreateBy: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	
	return &types.CommonResponse{
		Code : errx.Success,
		Msg: "添加成功",
	}, nil
}
