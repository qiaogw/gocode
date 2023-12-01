package {{.TableUrl}}logic

import (

	"github.com/qiaogw/gocode/common/errx"
	"context"
"github.com/pkg/errors"


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
func (l *Delete{{.Table}}Logic) Delete{{.Table}}(in *{{.Db}}.Delete{{.Table}}Request) (*{{.Db}}.NullResponse, error) {

	err := l.svcCtx.{{.Table}}Model.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
		"数据库删除 {{.TableComment}} 失败，id: %v,err:%v", in.Id,err)
	}

	return &{{.Db}}.NullResponse{}, nil
}
