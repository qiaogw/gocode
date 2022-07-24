package svc

import (
	"{{.ParentPkg}}/api/internal/config"
	"{{.ParentPkg}}/rpc/{{.Database}}"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	{{.Service}}Rpc {{.Database}}.{{.Service}}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		{{.Service}}Rpc: {{.Database}}.New{{.Service}}(zrpc.MustNewClient(c.{{.Service}}Rpc)),
	}
}