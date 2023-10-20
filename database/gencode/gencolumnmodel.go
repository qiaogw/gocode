// Package model 自动生成模板GenColumn(字段)
package gencode

import (
	"context"

	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ GenColumnModel = (*customGenColumnModel)(nil)

type (
	// GenColumnModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGenColumnModel.
	// 更多的自定义方法在这里添加，通过接口方法
	GenColumnModel interface {
		genColumnModel
		FindAll(ctx context.Context, query *ListGenColumnReq) ([]*Column, int64, error)
	}

	customGenColumnModel struct {
		*defaultGenColumnModel
	}

	SearchGenColumnModel struct {
		Column
		modelx.Pagination
	}
)

// NewGenColumnModel returns a model for the database table.
func NewGenColumnModel(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) GenColumnModel {
	return &customGenColumnModel{
		defaultGenColumnModel: newGenColumnModel(conn, c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customGenColumnModel) FindAll(ctx context.Context, query *ListGenColumnReq) ([]*Column, int64, error) {
	var resp []*Column
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
