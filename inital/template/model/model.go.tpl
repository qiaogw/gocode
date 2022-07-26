// package model 自动生成模板{{.Table}}({{.TableComment}})
package model

import (
	"context"

	"{{.ParentPkg}}/common/global"
	"github.com/qiaogw/gocode/util/gormq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ {{.Table}}Model = (*custom{{.Table}}Model)(nil)

type (
	// {{.Table}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.Table}}Model.
	// 更多的自定义方法在这里添加，通过接口方法
	{{.Table}}Model interface {
		{{.PackageName}}Model
		FindAll(ctx context.Context, query *List{{.Table}}Req) ([]*{{.Table}}, error)
	}

	custom{{.Table}}Model struct {
		*default{{.Table}}Model
	}

	Search{{.Table}}Model struct {
		{{.Table}}
		global.Pagination
	}
)

// New{{.Table}}Model returns a model for the database table.
func New{{.Table}}Model(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) {{.Table}}Model {
	return &custom{{.Table}}Model{
		default{{.Table}}Model: new{{.Table}}Model(conn, c, gormX),
	}
}

// FindAll 条件查询列表
func (m *custom{{.Table}}Model) FindAll(ctx context.Context, query *List{{.Table}}Req) ([]*{{.Table}}, error) {
	var resp []*{{.Table}}
	querys, values := m.GeneralSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Scopes(
			gormq.MakeCondition(query.GetNeedSearch(), m.gormDB.Name()),
			gormq.Paginate(query.GetPageSize(), query.GetPageIndex()),
		).Find(&resp).Limit(-1).Offset(-1)
	})
	err := m.QueryRowsNoCacheCtx(ctx, &resp, querys, values...)
	switch err {
	case nil:
		return resp, nil
	case global.ErrNotFound:
		return nil, global.ErrNotFound
	default:
		return nil, err
	}
}

func (m *{{.Table}}) GetNeedSearch() interface{} {
	return *m
}
