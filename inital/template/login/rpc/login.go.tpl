package logic

import (
"context"
"github.com/qiaogw/gocode/common/cryptx"
"github.com/qiaogw/gocode/global"
"google.golang.org/grpc/status"

"{{.ParentPkg}}/rpc/{{.Db}}"
"{{.ParentPkg}}/rpc/internal/svc"

"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
ctx    context.Context
svcCtx *svc.ServiceContext
logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
return &LoginLogic{
ctx:    ctx,
svcCtx: svcCtx,
Logger: logx.WithContext(ctx),
}
}

func (l *LoginLogic) Login(in *{{.Db}}.LoginRequest) (*{{.Db}}.LoginResponse, error) {
// 查询用户是否存在
res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
if err != nil {
if err == global.ErrNotFound {
return nil, status.Error(100, "用户不存在")
}
return nil, status.Error(500, err.Error())
}

// 判断密码是否正确
password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
if password != res.Password {
return nil, status.Error(100, "密码错误")
}

return &{{.Db}}.LoginResponse{
Id: res.Id,
}, nil
}
