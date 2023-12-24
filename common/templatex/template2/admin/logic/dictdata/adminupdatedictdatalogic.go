package dictdata

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictDataLogic {
	return &UpdateDictDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictDataLogic) UpdateDictData(req *types.UpdateDictDataRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.DictDataRpc.UpdateDictData(l.ctx, &admin.UpdateDictDataRequest{
		Id:         req.Id,
		DictTypeId: req.DictTypeId,
		Sort:       req.Sort,
		Label:      req.Label,
		Value:      req.Value,
		IsDefault:  req.IsDefault,
		Enabled:    req.Enabled,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "更新成功",
	}, nil
}
