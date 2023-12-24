package dicttype

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type DeleteDictTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictTypeLogic {
	return &DeleteDictTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDictTypeLogic) DeleteDictType(req *types.DeleteDictTypeRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.DictTypeRpc.DeleteDictType(l.ctx, &admin.DeleteDictTypeRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "删除成功",
	}, nil
}
