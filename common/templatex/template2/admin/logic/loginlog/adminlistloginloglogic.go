package loginlog

import (
	"context"
	"github.com/qiaogw/gocode/common/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"
)

type ListLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLoginLogLogic {
	return &ListLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLoginLogLogic) ListLoginLog(req *types.ListLoginLogRequest) (resp *types.CommonResponse, err error) {

	res, err := l.svcCtx.LoginLogRpc.ListLoginLog(l.ctx, &admin.ListLoginLogRequest{
		Username:      req.Username,
		Status:        req.Status,
		Ipaddr:        req.Ipaddr,
		LoginLocation: req.LoginLocation,
		Browser:       req.Browser,
		Os:            req.Os,
		Platform:      req.Platform,
		LoginTime:     req.LoginTime,
		Remark:        req.Remark,
		Msg:           req.Msg,
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
