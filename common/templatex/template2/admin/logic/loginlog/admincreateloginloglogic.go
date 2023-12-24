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

type CreateLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLoginLogLogic {
	return &CreateLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLoginLogLogic) CreateLoginLog(req *types.CreateLoginLogRequest) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.LoginLogRpc.CreateLoginLog(l.ctx, &admin.CreateLoginLogRequest{
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
		Msg:  "添加成功",
	}, nil
}
