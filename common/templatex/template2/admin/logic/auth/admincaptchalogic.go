package auth

import (
	"context"
	"github.com/qiaogw/gocode/common/captchax"
	"github.com/qiaogw/gocode/common/errx"
	"time"

	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"

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

type CaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}

func (l *CaptchaLogic) Captcha(req *types.CaptchaRequest) (resp *types.CommonResponse, err error) {
	cs := captchax.NewCacheStore(l.svcCtx.Cache, time.Second*120)
	captchax.SetStore(cs)
	id, b64s, err := captchax.DriverDigitFunc()
	if err != nil {
		logx.Error("验证码获取失败!", err)
		return nil, errx.NewDefaultError("验证码获取失败!")
	}
	res := &CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: l.svcCtx.Config.Captcha.KeyLong,
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "查询成功",
		Data: res,
	}, nil

}
