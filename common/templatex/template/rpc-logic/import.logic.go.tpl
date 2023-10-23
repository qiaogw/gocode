package logic

import (
	"bytes"
	"context"
	"{{.ParentPkg}}/rpc/{{.PackageName}}"
	"{{.ParentPkg}}/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type Import{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImport{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Import{{.Table}}Logic {
	return &Import{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Import{{.Table}}Logic) Import{{.Table}}(in *{{.PackageName}}.ExportResponse) (*{{.PackageName}}.Delete{{.Table}}Response, error) {
	reader := bytes.NewReader(in.Data)
	err := l.svcCtx.{{.Table}}Model.Import(reader)

	return &{{.PackageName}}.Delete{{.Table}}Response{}, err
}
