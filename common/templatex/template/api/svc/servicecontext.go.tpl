package svc

import (
	"{{.ParentPkg}}/api/internal/config"
	"{{.ParentPkg}}/rpc/{{.FileName}}client"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config    config.Config
	{{.Service}}Rpc {{.FileName}}client.{{.Service}}
	Cache    cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		{{.Service}}Rpc: {{.FileName}}client.New{{.Service}}(zrpc.MustNewClient(c.{{.Service}}Rpc)),
		Cache:    cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("dc"), nil),
	}
}