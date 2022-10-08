package logic

import (
"context"
"github.com/qiaogw/gocode/common/captchax"
"github.com/qiaogw/gocode/common/errorx"
"github.com/qiaogw/gocode/common/jwtx"
"{{.ParentPkg}}/rpc/{{.Db}}"
"time"

"{{.ParentPkg}}/api/internal/svc"
"{{.ParentPkg}}/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
if !captchax.Verify(req.CaptchaId, req.Captcha, true) {
return nil, errorx.NewDefaultError("验证码错误")
}
res, err := l.svcCtx.AdminRpc.Login(l.ctx, &{{.Db}}.LoginRequest{
Mobile:   req.Mobile,
Password: req.Password,
})
if err != nil {
return nil, err
}

now := time.Now().Unix()
accessExpire := l.svcCtx.Config.Auth.AccessExpire

accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
if err != nil {
return nil, err
}

return &types.LoginResponse{
AccessToken:  accessToken,
AccessExpire: now + accessExpire,
}, nil
}
