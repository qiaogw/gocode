package {{.TableUrl}}

import (
	"context"
"github.com/qiaogw/gocode/common/jwtx"
"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/jinzhu/copier"

{{- if .IsFlow }}
	"{{.Pkg}}/fsm/rpc/fsmx"
	"{{.Pkg}}/fsm/rpc/client/flow"
{{- end}}
)

type Get{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Get{{.Table}}Logic {
	return &Get{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
// Get{{.Table}} 获取单个{{.TableComment}}
func (l *Get{{.Table}}Logic) Get{{.Table}}(req *types.Get{{.Table}}Request) (resp *types.CommonResponse, err error) {
	userId := jwtx.GetUserIdFromCtx(l.ctx)
	res, err := l.svcCtx.{{.Table}}Rpc.Get{{.Table}}(l.ctx, &{{.Db}}.Get{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
				{{.FieldName}}: req.{{.FieldName}},
			{{- end}}
		{{- end }}
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	var rep types.Get{{.Table}}Response
	_ = copier.Copy(&rep, res)


{{- if .IsFlow }}
	flowIns, err := l.svcCtx.FlowInstanceRpc.GetFlowInstanceByBusy(l.ctx, &flow.GetFlowInstanceByBusyRequest{
	BusyId:   res.Id,
	BusyName: res.BusyName,
	UserId:   userId,
	})
	if err != nil {
	return nil, errors.Wrapf(err, "req: %+v", req)
	}
	history, err := l.svcCtx.FlowInstanceRpc.GetFlowInstanceHistoryByBusy(l.ctx, &flow.GetFlowInstanceByBusyRequest{
	BusyId:   res.Id,
	BusyName: res.BusyName,
	UserId:   userId,
	})
	if err != nil {
	return nil, errors.Wrapf(err, "req: %+v", req)
	}
	var fins types.GetFlowInstanceResponse
	_ = copier.Copy(&fins, flowIns)
	var list []*types.GetFlowInstanceResponse
	_ = copier.Copy(&list, history.List)
	rep.FlowInstance = fins
	rep.FlowInstanceHistory = list
	_ = copier.Copy(&rep.CurrentNode, flowIns.CurrentNode)

{{- end}}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "查询成功",
		Data: rep,
	}, nil
}
