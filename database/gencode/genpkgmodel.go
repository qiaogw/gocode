// Package model 自动生成模板GenPkg(包)
package gencode

import (
	"context"

	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ GenPkgModel = (*customGenPkgModel)(nil)

type (
	// GenPkgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGenPkgModel.
	// 更多的自定义方法在这里添加，通过接口方法
	GenPkgModel interface {
		genPkgModel
		FindAll(ctx context.Context, query *SearchGenPkgModel) ([]*GenPkg, int64, error)
	}

	customGenPkgModel struct {
		*defaultGenPkgModel
	}

	SearchGenPkgModel struct {
		GenPkg
		modelx.Pagination
		GenPkgOrder
	}
)

// NewGenPkgModel returns a model for the database table.
func NewGenPkgModel(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) GenPkgModel {
	return &customGenPkgModel{
		defaultGenPkgModel: newGenPkgModel(conn, c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customGenPkgModel) FindAll(ctx context.Context, query *SearchGenPkgModel) ([]*GenPkg, int64, error) {
	var resp []*GenPkg
	var count int64
	findKeySql := gormx.SearchKey(m.gormDB, m.tableName(), query.SearchKey)
	res := m.gormDB.Scopes(
		gormx.MakeCondition(query.GetNeedSearch(), m.gormDB.Name()),
		gormx.Paginate(query.GetPageSize(), query.GetPageIndex()),
		gormx.SortBy(query.SortBY, query.Descending),
	).Where(findKeySql).Find(&resp).Limit(-1).Offset(-1)
	res.Count(&count)
	err := res.Error
	switch err {
	case nil:
		return resp, count, nil
	case modelx.ErrNotFound:
		return nil, 0, modelx.ErrNotFound
	default:
		return nil, 0, err
	}
}

func (m *SearchGenPkgModel) GetNeedSearch() interface{} {
	return m.GenPkgOrder
}
