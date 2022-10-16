package logic

import (
	"{{.PKG}}/common/modelx"
	"{{.PKG}}/common/errorx"
	"context"
	"google.golang.org/grpc/status"
"github.com/jinzhu/copier"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

)

type Get{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGet{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Get{{.Table}}Logic {
	return &Get{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get{{.Table}} 提取单条 {{.TableComment}}
func (l *Get{{.Table}}Logic) Get{{.Table}}(in *{{.Db}}.Get{{.Table}}Request) (*{{.Db}}.Get{{.Table}}Response, error) {
	// 查询{{.TableComment}} 是否存在
	res, err := l.svcCtx.{{.Table}}Model.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errorx.NewCodeError(errorx.NoData, "该{{.TableComment}} 不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	var rep {{.Db}}.Get{{.Table}}Response
	_ = copier.Copy(&rep, res)


	return &rep, nil
}
