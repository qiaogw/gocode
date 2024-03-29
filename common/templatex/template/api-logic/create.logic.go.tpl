package {{.TableUrl}}

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"
	"github.com/zeromicro/go-zero/core/logx"

{{- if .IsFlow }}
	"{{.Pkg}}/fsm/fsmx"
	"{{.Pkg}}/fsm/rpc/client/flow"
	"{{.Pkg}}/gencode/rpc/gencode"
{{- end}}

)

type Create{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Create{{.Table}}Logic {
	return &Create{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
//Create{{.Table}} 创建{{.TableComment}}
func (l *Create{{.Table}}Logic) Create{{.Table}}(req *types.Create{{.Table}}Request) (resp *types.CommonResponse, err error) {
	userId := jwtx.GetUserIdFromCtx(l.ctx)
{{- if .IsFlow }}
	busyName, _ := l.svcCtx.{{.Table}}Rpc.GetBusyName{{.Table}}(l.ctx, &{{.Db}}.NullRequest{})
	_, err = l.svcCtx.FlowRpc.GetFlowByBusy(l.ctx, &flow.GetFlowByBusyRequest{
		BusyName: busyName.Name,
		Type:     fsmx.FORM,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	db, err := l.svcCtx.GencodeRpc.GetGenSourceByName(l.ctx, &gencode.GetGenSourceRequest{
	Name: l.svcCtx.Config.Name,
	})
	if err != nil {
	return nil, errors.Wrapf(err, "req: %+v", req)
	}
	_, err = l.svcCtx.FlowRpc.GetFlowByBusy(l.ctx, &flow.GetFlowByBusyRequest{
	BusyName:  busyName.Name,
	Type:      fsmx.FORM,
	ModelName: db.Id,
	})
	if err != nil {
	return nil, errors.Wrapf(err, "req: %+v", req)
	}

{{- end }}

	res, err := l.svcCtx.{{.Table}}Rpc.Create{{.Table}}(l.ctx, &{{.Db}}.Create{{.Table}}Request{
		{{- range  .Columns }}
			{{- if .IsPk }}
			{{- else if .IsModelTime -}}
			{{- else if .IsControl -}}
			{{- else}}
				{{- if .IsPage}}
				{{- else}}
				{{.FieldName}}: req.{{.FieldName}},
				{{- end}}
			{{- end}}
		{{- end }}
		CreateBy: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}


{{- if .IsFlow }}
flowRes, err := l.svcCtx.FlowInstanceRpc.CreateFlowInstance(l.ctx, &flow.CreateFlowInstanceRequest{
		BusyName: res.BusyName,
		BusyId:   res.Id,
		Type:     string(fsmx.FORM),
		Enabled:  true,
		CreateBy: userId,
	})
	data, err := l.svcCtx.{{.Table}}Rpc.Get{{.Table}}(l.ctx, &{{.Db}}.Get{{.Table}}Request{
		Id: res.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	var upReq {{.Db}}.Update{{.Table}}Request
	_ = copier.Copy(&upReq, data)
	upReq.Status = flowRes.FlowStatus
	upReq.Remark = flowRes.Remark
	_, err = l.svcCtx.{{.Table}}Rpc.Update{{.Table}}(l.ctx, &upReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
{{- else}}
	_=res
{{- end}}


	return &types.CommonResponse{
		Code : errx.Success,
		Msg: "添加成功",
	}, nil
}
