package auth

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"
	"net/http"
	"strings"
	"sub-admin/admin/api/internal/svc"
	"sub-admin/admin/api/internal/types"
	"sub-admin/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(req *types.CaptchaRequest, r *http.Request) (resp *types.VerifyResponse, err error) {
	reqUrl := r.Header.Get("X-Forwarded-Uri")
	token := r.Header.Get("Authorization")
	if strings.Contains(reqUrl, "?") {
		reqUrl = strings.Split(reqUrl, "?")[0]
	}

	var realUserId string
	if l.urlNoAuth(reqUrl) {
		//	不需要验证的url
		if len(token) < 0 {
			userid, err := jwtx.GetUserIdFromToken(token, l.svcCtx.Config.Auth.AccessSecret)
			if err != nil {
				return nil, errors.Wrapf(errx.NewErrCode(errx.ErrAuth),
					"Authorization:%s,reqUrl:%s", token, reqUrl)
			}
			if userid == "" {
				return nil, errors.Wrapf(errx.NewErrCode(errx.ErrAuth),
					"urlNoAuth，userid is 0，Authorization:%s,reqUrl:%s", token, reqUrl)
			}
			realUserId = userid
		}
	} else {
		userid, err := l.isPass(r)
		if err != nil {
			return nil, errors.Wrapf(errx.NewErrCode(errx.ErrAuth),
				"Authorization:%s,reqUrl:%s", token, reqUrl)
		}
		if userid == "" {
			return nil, errors.Wrapf(errx.NewErrCode(errx.ErrAuth),
				"urlAuth，userid is 0，Authorization:%s,reqUrl:%s", token, reqUrl)
		}
		realUserId = userid
	}
	return &types.VerifyResponse{
		Token:  token,
		UserId: realUserId,
		Ok:     true,
	}, nil
}

// 当前url是否需要授权验证
func (l *VerifyLogic) urlNoAuth(path string) bool {
	for _, val := range l.svcCtx.Config.NoAuthUrls {
		if val == path {
			return true
		}
	}
	return false
}

// 验证是否通过
func (l *VerifyLogic) isPass(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	userid, err := jwtx.GetUserIdFromToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return "", err
	}
	resp, err := l.svcCtx.AuthRpc.Verify(l.ctx, &admin.ValidateTokenReq{
		Path:   r.Header.Get("X-Forwarded-Uri"),
		Method: r.Header.Get("X-Forwarded-Method"),
		UserId: userid,
		Token:  token,
	})
	if err != nil || !resp.Ok {
		return "", err
	}
	return userid, nil
}
