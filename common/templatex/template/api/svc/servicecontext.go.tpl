package svc

import (
	"{{.ParentPkg}}/api/internal/config"

{{- range .Tables }}
	"{{.ParentPkg}}/rpc/client/{{.TableUrl}}"
{{- end}}
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
{{- if .IsFlow }}
	"{{.Pkg}}/admin/rpc/client/user"
	"{{.Pkg}}/fsm/rpc/client/flow"
	"{{.Pkg}}/fsm/rpc/client/flowinstance"
{{- end}}
)

type ServiceContext struct {
	Config    config.Config
	Cache    cache.Cache
{{- range .Tables }}
	{{.Table}}Rpc {{.TableUrl}}.{{.Table}}
{{- end}}
{{- if .IsFlow }}
	UserRpc         user.User
	FlowInstanceRpc flowinstance.FlowInstance
	FlowRpc         flow.Flow
{{- end}}

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Cache:    cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("sub-{{.Package}}-api"), nil),
{{- range .Tables }}
	{{.Table}}Rpc: {{.TableUrl}}.New{{.Table}}(zrpc.MustNewClient(c.{{.Service}}Rpc)),
{{- end}}
{{- if .IsFlow }}
		FlowInstanceRpc: flowinstance.NewFlowInstance(zrpc.MustNewClient(c.FsmRpc)),
		UserRpc:         user.NewUser(zrpc.MustNewClient(c.AdminRpc)),
		FlowRpc:         flow.NewFlow(zrpc.MustNewClient(c.FsmRpc)),
{{- end}}
	}
}