// Package model 自动生成模板{{.Table}}({{.TableComment}})
package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ {{.Table}}Model = (*custom{{.Table}}Model)(nil)

type (
	// {{.Table}}Model 是一个需要自定义的接口，这里添加更多方法
	{{.Table}}Model interface {
		{{.PackageName}}Model
		FindAll(ctx context.Context, query *List{{.Table}}Req) ([]*{{.Table}},int64, error)
		Import(reader io.Reader) error
	}

	custom{{.Table}}Model struct {
		*default{{.Table}}Model
	}

	Search{{.Table}}Model struct {
		{{.Table}}
		modelx.Pagination
	}
)

// New{{.Table}}Model returns a model for the database table.
func New{{.Table}}Model(conn sqlx.SqlConn,{{if .withCache}}, c cache.CacheConf, opts ...cache.Option{{end}})) {{.Table}}Model {
	return &custom{{.Table}}Model{
		default{{.Table}}Model: new{{.Table}}Model(conn, c, opts...),
	}
}
