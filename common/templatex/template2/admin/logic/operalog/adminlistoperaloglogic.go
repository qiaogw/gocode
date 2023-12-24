package operalog

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ListOperaLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOperaLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOperaLogLogic {
	return &ListOperaLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOperaLogLogic) ListOperaLog(req *types.ListOperaLogRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.OperaLogRpc.ListOperaLog(l.ctx, &admin.ListOperaLogRequest{
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
		PageIndex:     req.PageIndex,
		PageSize:      req.PageSize,
		SearchKey:     req.SearchKey,
		SortBy:        req.SortBY,
		Descending:    req.Descending,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "查询成功",
		Data: res,
	}, nil
}
