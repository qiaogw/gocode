package {{.TableUrl}}logic

import (
	"context"

	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBusyName{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBusyName{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBusyName{{.Table}}Logic {
	return &GetBusyName{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetBusyName{{.Table}} 提取业务表名称
func (l *GetBusyName{{.Table}}Logic) GetBusyName{{.Table}}(in *{{.Db}}.NullRequest) (*{{.Db}}.BusyNameResponse, error) {
	res := l.svcCtx.{{.Table}}Model.GetName()

	return &{{.Db}}.BusyNameResponse{
		Name: res,
	}, nil
}
