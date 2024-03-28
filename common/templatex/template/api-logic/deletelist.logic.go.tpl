package {{.TableUrl}}

import (
	"context"

"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteList{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteList{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteList{{.Table}}Logic {
	return &DeleteList{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
// DeleteList{{.Table}} 删除多个{{.TableComment}}
func (l *DeleteList{{.Table}}Logic) DeleteList{{.Table}}(req *types.DeleteList{{.Table}}Request) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.{{.Table}}Rpc.DeleteList{{.Table}}(l.ctx, &{{.Db}}.DeleteList{{.Table}}Request{
		List: req.IdList,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	
	return &types.CommonResponse{
	Code: errx.Success,
	Msg: "删除成功",
}, nil
}
