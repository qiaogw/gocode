package auth

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/captchax"
	"github.com/qiaogw/gocode/common/errx"
	"sub-admin/admin/rpc/admin"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.CommonResponse, err error) {
	if !captchax.Verify(req.CaptchaId, req.Captcha, true) {
		return nil, errx.NewDefaultError("验证码错误")
	}
	res, err := l.svcCtx.AuthRpc.Login(l.ctx, &admin.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "登录成功",
		Data: res,
	}, nil
}
