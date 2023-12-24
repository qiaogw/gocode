package loginlog

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLoginLogLogic {
	return &UpdateLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLoginLogLogic) UpdateLoginLog(req *types.UpdateLoginLogRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.LoginLogRpc.UpdateLoginLog(l.ctx, &admin.UpdateLoginLogRequest{
		Id:            req.Id,
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
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "更新成功",
	}, nil
}
