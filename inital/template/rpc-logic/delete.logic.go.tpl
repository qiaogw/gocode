package logic

import (
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/common/errorx"
	"context"
	"google.golang.org/grpc/status"

	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"


	"github.com/zeromicro/go-zero/core/logx"
)

type Delete{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelete{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Delete{{.Table}}Logic {
	return &Delete{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Delete{{.Table}} 删除 {{.TableComment}}
func (l *Delete{{.Table}}Logic) Delete{{.Table}}(in *{{.Db}}.Delete{{.Table}}Request) (*{{.Db}}.Delete{{.Table}}Response, error) {
	// 查询 {{.TableComment}}是否存在
	res, err := l.svcCtx.{{.Table}}Model.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == global.ErrNotFound {
			return nil, errorx.NewCodeError(errorx.NoData, "该{{.TableComment}}不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.{{.Table}}Model.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &{{.Db}}.Delete{{.Table}}Response{}, nil
}
