package {{.TableUrl}}

import (
	"context"
"github.com/qiaogw/gocode/common/errx"
"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"github.com/qiaogw/gocode/common/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
{{- if .IsFlow }}
		"{{.Pkg}}/fsm/rpc/fsm"
{{- end}}

)

type Trigger{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTrigger{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Trigger{{.Table}}Logic {
	return &Trigger{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Trigger{{.Table}}Logic) Trigger{{.Table}}(req *types.UpdateNodeInstanceRequest) (resp *types.CommonResponse, err error) {
	userId := jwtx.GetUserIdFromCtx(l.ctx)
	var nodeReq fsm.UpdateNodeInstanceRequest
	_ = copier.Copy(&nodeReq, req)
	nodeReq.UpdateBy = userId
	res, err := l.svcCtx.FlowInstanceRpc.TriggerFlowInstance(l.ctx, &nodeReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	data, err := l.svcCtx.{{.Table}}Rpc.Get{{.Table}}(l.ctx, &cms.Get{{.Table}}Request{
		Id: res.BusyId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	var upReq {{.Db}}.Update{{.Table}}Request
	_ = copier.Copy(&upReq, data)
	upReq.Status = res.FlowStatus
	upReq.Remark = res.Remark

	_, err = l.svcCtx.{{.Table}}Rpc.Update{{.Table}}(l.ctx, &upReq)
	if err != nil {
		return nil,errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "审批成功",
	}, nil
}
