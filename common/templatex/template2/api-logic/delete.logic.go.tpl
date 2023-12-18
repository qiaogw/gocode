package {{.TableUrl}}

import (
	"context"

	"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"

	"github.com/qiaogw/gocode/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
)

type Delete{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Delete{{.Table}}Logic {
	return &Delete{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete{{.Table}}Logic) Delete{{.Table}}(req *types.Delete{{.Table}}Request) (resp *types.CommonResponse, err error) {

	err = l.svcCtx.{{.Table}}Model.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData),
		"数据库删除 {{.TableComment}} 失败，id: %v,err:%v", req.Id,err)
	}
	
	return &types.CommonResponse{
		Code: errx.Success,
		Msg: "删除成功",
	}, nil
}
