package logic

import (
"context"
"github.com/qiaogw/gocode/service/api/internal/svc"
"github.com/qiaogw/gocode/service/api/internal/types"
"github.com/qiaogw/gocode/service/common/captchax"
"github.com/qiaogw/gocode/service/common/errorx"
"time"

"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
logx.Logger
ctx    context.Context
svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
return &CaptchaLogic{
Logger: logx.WithContext(ctx),
ctx:    ctx,
svcCtx: svcCtx,
}
}

func (l *CaptchaLogic) Captcha(req *types.CaptchaRequest) (resp *types.CaptchaResponse, err error) {
cs := captchax.NewCacheStore(l.svcCtx.Cache, time.Second*120)
captchax.SetStore(cs)
id, b64s, err := captchax.DriverDigitFunc()
if err != nil {
logx.Error("验证码获取失败!", err)
return nil, errorx.NewDefaultError("验证码获取失败!")
}
return &types.CaptchaResponse{
CaptchaId:     id,
PicPath:       b64s,
CaptchaLength: l.svcCtx.Config.Captcha.KeyLong,
}, nil
}
