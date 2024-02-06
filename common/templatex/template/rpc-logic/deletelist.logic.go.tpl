package {{.TableUrl}}logic

import (

	"github.com/qiaogw/gocode/common/errx"
	"context"
	"github.com/pkg/errors"


	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"


	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteList{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteList{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Delete{{.Table}}Logic {
	return &DeleteList{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteList{{.Table}} 批量删除 {{.TableComment}}
func (l *DeleteList{{.Table}}Logic) DeleteList{{.Table}}(in *{{.Db}}.DeleteList{{.Table}}Request) (*{{.Db}}.NullResponse, error) {

	err := l.svcCtx.{{.Table}}Model.DeleteList(l.ctx, in.List)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
		"数据库批量删除 {{.TableComment}} 失败:%v", err)
	}

	return &{{.Db}}.NullResponse{}, nil
}
