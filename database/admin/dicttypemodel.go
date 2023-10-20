// Package model 自动生成模板DictType(字典)
package admin

import (
	"context"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ DictTypeModel = (*customDictTypeModel)(nil)

type (
	// DictTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDictTypeModel.
	// 更多的自定义方法在这里添加，通过接口方法
	DictTypeModel interface {
		dictTypeModel
		FindAll(ctx context.Context, query *ListDictTypeReq) ([]*DictType, int64, error)
	}

	customDictTypeModel struct {
		*defaultDictTypeModel
	}

	SearchDictTypeModel struct {
		DictType
		modelx.Pagination
	}
)

// NewDictTypeModel returns a model for the database table.
func NewDictTypeModel(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) DictTypeModel {
	return &customDictTypeModel{
		defaultDictTypeModel: newDictTypeModel(conn, c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customDictTypeModel) FindAll(ctx context.Context, query *ListDictTypeReq) ([]*DictType, int64, error) {
	var resp []*DictType
	var count int64
	sql := gormx.SearchKey(m.gormDB, m.tableName(), query.SearchKey)
	res := m.gormDB.Scopes(
		gormx.MakeCondition(query.GetNeedSearch(), m.gormDB.Name()),
		gormx.Paginate(query.GetPageSize(), query.GetPageIndex()),
		gormx.SortBy(query.SortBY, query.Descending),
	).Preload("DictDatas", func(db *gorm.DB) *gorm.DB {
		return db.Order("admin_dict_data.sort")
	}).Where(sql).Find(&resp).Limit(-1).Offset(-1)
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
