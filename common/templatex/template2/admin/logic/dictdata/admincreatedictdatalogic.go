package dictdata

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type CreateDictDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDictDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictDataLogic {
	return &CreateDictDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDictDataLogic) CreateDictData(req *types.CreateDictDataRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.DictDataRpc.CreateDictData(l.ctx, &admin.CreateDictDataRequest{
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
		Msg:  "添加成功",
	}, nil
}
