package {{.TableUrl}}logic

import (
	"context"

	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBusyNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBusyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBusyNameLogic {
	return &GetBusyNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetBusyName 提取业务表名称
func (l *GetBusyNameLogic) GetBusyName(in *{{.Db}}.NullRequest) (*{{.Db}}.BusyNameResponse, error) {
	res := l.svcCtx.{{.Table}}Model.GetName()

	return &{{.Db}}.BusyNameResponse{
		Name: res,
	}, nil
}
