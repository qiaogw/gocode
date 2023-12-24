package operalog

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOperaLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOperaLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOperaLogLogic {
	return &UpdateOperaLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOperaLogLogic) UpdateOperaLog(req *types.UpdateOperaLogRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.OperaLogRpc.UpdateOperaLog(l.ctx, &admin.UpdateOperaLogRequest{
		Id:            req.Id,
		Title:         req.Title,
		BusinessType:  req.BusinessType,
		BusinessTypes: req.BusinessTypes,
		Method:        req.Method,
		RequestMethod: req.RequestMethod,
		OperatorType:  req.OperatorType,
		OperName:      req.OperName,
		DeptName:      req.DeptName,
		OperUrl:       req.OperUrl,
		OperIp:        req.OperIp,
		OperLocation:  req.OperLocation,
		OperParam:     req.OperParam,
		Status:        req.Status,
		OperTime:      req.OperTime,
		JsonResult:    req.JsonResult,
		Remark:        req.Remark,
		LatencyTime:   req.LatencyTime,
		UserAgent:     req.UserAgent,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "更新成功",
	}, nil
}
