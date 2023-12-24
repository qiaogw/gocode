package dicttype

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictTypeLogic {
	return &UpdateDictTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictTypeLogic) UpdateDictType(req *types.UpdateDictTypeRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.DictTypeRpc.UpdateDictType(l.ctx, &admin.UpdateDictTypeRequest{
		Id:      req.Id,
		Name:    req.Name,
		Type:    req.Type,
		Enabled: req.Enabled,
		Remark:  req.Remark,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "更新成功",
	}, nil
}
