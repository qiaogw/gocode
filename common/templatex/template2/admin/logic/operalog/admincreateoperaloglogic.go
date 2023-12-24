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

type CreateOperaLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOperaLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOperaLogLogic {
	return &CreateOperaLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOperaLogLogic) CreateOperaLog(req *types.CreateOperaLogRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.OperaLogRpc.CreateOperaLog(l.ctx, &admin.CreateOperaLogRequest{
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
		Msg:  "添加成功",
	}, nil
}
