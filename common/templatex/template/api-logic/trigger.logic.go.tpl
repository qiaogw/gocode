package {{.TableUrl}}

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/condition"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"
	"github.com/zeromicro/go-zero/core/logx"

	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"

{{- if .IsFlow }}
	"{{.Pkg}}/fsm/fsmx"
	"{{.Pkg}}/fsm/rpc/fsm"
{{- end}}

)

type TriggerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTriggerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TriggerLogic {
	return &TriggerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TriggerLogic) Trigger(req *types.UpdateNodeInstanceRequest) (resp *types.CommonResponse, err error) {
	userId := jwtx.GetUserIdFromCtx(l.ctx)
	var nodeReq fsm.UpdateNodeInstanceRequest
	_ = copier.Copy(&nodeReq, req)
	nodeReq.UpdateBy = userId

	res, err := l.svcCtx.FlowInstanceRpc.Trigger(l.ctx, &nodeReq)
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
	upReq.FlowStatus = res.FlowStatus
	upReq.Remark = res.Remark

	if res.CurrentNode.Type == fsmx.ConditionState {
		crn := res.CurrentNode.Id
		events, err := l.svcCtx.FlowInstanceRpc.GetFlowEventCondition(l.ctx, &nodeReq)
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		for _, o := range events.List {
			var conds []condition.Condition
			_ = copier.Copy(&conds, o.ConditionList)
			if condition.EvaluateCondition(data, conds) {
				nodeReq.Event = o.Name
				evr, err := l.svcCtx.FlowInstanceRpc.Event(l.ctx, &nodeReq)
				if err != nil {
					return nil, errors.Wrapf(err, "req: %+v", req)
				}
				crn = evr.CurrentNode.Id
				upReq.FlowStatus = evr.FlowStatus
				upReq.Remark = evr.Remark
				break
			}
		}
		if res.CurrentNode.Id == crn {
			logx.Info("【condition no")
		}
	}

	_, err = l.svcCtx.{{.Table}}Rpc.Update{{.Table}}(l.ctx, &upReq)
	if err != nil {
		return nil,errors.Wrapf(err, "req: %+v", req)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "审批成功",
	}, nil
}
