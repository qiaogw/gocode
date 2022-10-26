package logic

import (
"{{.PKG}}/common/modelx"
	"{{.PKG}}/common/errx"
	"context"
"github.com/pkg/errors"
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
		if err == modelx.ErrNotFound {
			return nil, errors.Wrapf(errx.NewErrCode(errx.NoData), "该{{.TableComment}}不存在，id: %v", in.Id)
		}
		return nil,  errors.Wrapf(errx.NewErrCode(errx.NoData),
"查询 {{.TableComment}} db fail，id: %v,err:%v", in.Id)
	}

	err = l.svcCtx.{{.Table}}Model.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
"删除 {{.TableComment}} db fail，id: %v,err:%v", in.Id,err)
}

	return &{{.Db}}.Delete{{.Table}}Response{}, nil
}
