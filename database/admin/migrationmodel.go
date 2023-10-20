// package model 自动生成模板Migration(版本)
package admin

import (
	"context"

	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ MigrationModel = (*customMigrationModel)(nil)

type (
	// MigrationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMigrationModel.
	// 更多的自定义方法在这里添加，通过接口方法
	MigrationModel interface {
		migrationModel
		FindAll(ctx context.Context, query *ListMigrationReq) ([]*Migration, int64, error)
	}

	customMigrationModel struct {
		*defaultMigrationModel
	}

	SearchMigrationModel struct {
		Migration
		modelx.Pagination
	}
)

// NewMigrationModel returns a model for the database table.
func NewMigrationModel(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) MigrationModel {
	return &customMigrationModel{
		defaultMigrationModel: newMigrationModel(conn, c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customMigrationModel) FindAll(ctx context.Context, query *ListMigrationReq) ([]*Migration, int64, error) {
	var resp []*Migration
	var count int64
	sql := gormx.SearchKey(m.gormDB, m.tableName(), query.SearchKey)
	res := m.gormDB.Scopes(
		gormx.MakeCondition(query.GetNeedSearch(), m.gormDB.Name()),
		gormx.Paginate(query.GetPageSize(), query.GetPageIndex()),
		gormx.SortBy(query.SortBY, query.Descending),
	).Preload("DictDatas").Where(sql).Find(&resp).Limit(-1).Offset(-1)
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
