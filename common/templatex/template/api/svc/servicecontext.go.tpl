package svc

import (
	"{{.ParentPkg}}/api/internal/config"

{{- range .Tables }}
	"{{.ParentPkg}}/rpc/client/{{.TableUrl}}"
{{- end}}
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config    config.Config
	Cache    cache.Cache
{{- range .Tables }}
	{{.Table}}Rpc {{.TableUrl}}.{{.Table}}
{{- end}}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Cache:    cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("sub-{{.Package}}-api"), nil),
{{- range .Tables }}
	{{.Table}}Rpc: {{.TableUrl}}.New{{.Table}}(zrpc.MustNewClient(c.{{.Service}}Rpc)),
{{- end}}
	}
}